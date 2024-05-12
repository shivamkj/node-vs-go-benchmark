const fastify = require("fastify")();
const jwt = require("jsonwebtoken");
const { generateToken } = require("./jwt.util");
const { Client } = require("pg");
const cluster = require("cluster");

const workersCount = 3;

const SECRET_KEY =
  "c7fefec7e3c0d46dd5da07ff172e7c61bd982a63071fcff9ba24165bc6753a66";

// PostgreSQL connection
const client = new Client({
  connectionString: "postgresql://shivam:pass@localhost:5432/qnify",
});

// Connect to PostgreSQL
client.connect();

fastify.register(async function plugin(fastify, opts) {
  // Middleware to verify JWT token
  fastify.addHook("onRequest", (request, reply, done) => {
    const token = request.headers.authorization;

    if (!token) {
      reply.code(401).send({ error: "Unauthorized" });
      return;
    }

    try {
      const decoded = jwt.verify(token, SECRET_KEY);
      request.user = decoded;
    } catch (err) {
      reply.code(401).send({ error: "Invalid token" });
      return;
    }

    done();
  });

  // Route to get all users
  fastify.get("/users", async (request, reply) => {
    try {
      const query =
        "SELECT id, email, mobilenum, firstname, lastname, profilepic FROM users LIMIT 20";
      const result = await client.query(query);
      reply.send(result.rows);
    } catch (err) {
      reply.code(500).send({ error: "Internal Server Error" });
    }
  });
});

// Route to generate a JWT token
fastify.get("/login", (request, reply) => {
  const payload = {
    userId: 123, // Replace with actual user ID or any other payload you want to include in the token
  };
  const token = generateToken(payload);
  reply.send({ token });
});

if (cluster.isMaster) {
  console.log("Workers", workersCount);
  for (let index = 0; index < workersCount; index++) {
    cluster.fork();
  }
} else {
  // Start the server
  fastify.listen({ port: 3000 }, (err, address) => {
    if (err) {
      console.error(err);
      process.exit(1);
    }
    console.log(`Server listening at ${address} pid: ${process.pid}`);
  });
}
