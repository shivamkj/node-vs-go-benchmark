import http from "k6/http";
import { sleep, check } from "k6";

export const options = {
  vus: 10000, // Number of virtual users
  duration: "30s",
};

const baseURL = "http://localhost:3000";

// NOTE: update token before starting test, token is valid for 10 hours
const token =
  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTU1NjM4MjIsInVzZXJfaWQiOiIxMjMifQ.ZPhH7d2PbT87_VB57c-cpakq9xDrK8ayPZW7e-47OAw";

export default function () {
  const headers = {
    Authorization: token,
  };

  const res = http.get(`${baseURL}/users`, {
    headers,
    tags: {
      name: "GetUsers",
    },
  });

  check(res, {
    "status is 200": (r) => r.status === 200,
  });

  sleep(1); // Add a small delay between requests to simulate real-world usage
}
