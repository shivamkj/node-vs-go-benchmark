const jwt = require("jsonwebtoken");

const SECRET_KEY =
  "c7fefec7e3c0d46dd5da07ff172e7c61bd982a63071fcff9ba24165bc6753a66";

const generateToken = (payload) => {
  const token = jwt.sign(payload, SECRET_KEY, { expiresIn: "10h" });
  return token;
};

module.exports = {
  generateToken,
};
