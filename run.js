import http from 'k6/http';
import { check, sleep } from 'k6';


export let options = {
  stages: [
    { duration: '5s', target: 10 },
    { duration: '5s', target: 15 },
    { duration: '5s', target: 30 },
    { duration: '5s', target: 15 },
    { duration: '5s', target: 10 },
  ],
};

export default function () {
  let endpoint = "http://" + __ENV.ENDPOINT + "/ping"
  let res = http.get(endpoint);
  check(res, { 'status was 200': r => r.status == 200 });
  sleep(1);
}