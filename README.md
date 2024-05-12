### Node Vs Go Benchmark

This project tries to benchmark NodeJS (Fastify) & Go (Fiber). In the API it checks JWT token and do a listing of users from PostgreSQL, which represents a typical real life scenario. For bench-marking it uses k6 with 10,000 virtual users for 30 seconds. 3 workers are forked in NodeJS as NodeJS is single threaded & Go utilises multiple threads.

To run just generate a new JWT token by calling [`http://127.0.0.1:3000/login`](http://127.0.0.1:3000/login) after starting the server. Now, update the JWT token in `benchmark.js` and run the benchmark with command: `k6 run benchmark.js`

Result: Go is about 20-25% faster than NodeJS, based on my observation after analysing several parameters and testing multiples times in different environments. Also, NodeJS is more memory intensive than Go.