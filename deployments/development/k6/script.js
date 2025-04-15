import http from 'k6/http';
import {check, group, sleep} from 'k6';
// see more: https://medium.com/@ravipatel.it/step-by-step-guide-to-load-testing-with-k6-5afb625e231a
// K6 Cloud

// Understanding Virtual Users (VUs) and Duration
// Virtual Users (VUs) simulate real users accessing your application. You can specify the number of VUs and the duration of the test
export let options = {
    vus: 100, // Number of virtual users
    duration: '15s', // Duration of the test
};
// Let’s create a more complex scenario where the number of users ramps up, holds steady, and then ramps down.
// export let options = {
//     stages: [
//         { duration: '30s', target: 20 },
//         { duration: '1m', target: 20 },
//         { duration: '10s', target: 0 },
//     ],
// };

const BASE_URL = 'http://nginx:80';
// const BASE_URL = 'http://app:8000';
// const BASE_URL = 'http://host.docker.internal:80';

// export default function () {
//     http.get('https://test.k6.io');
//     sleep(1);
// }
export default function () {
    group("healthz", () => {
        let res = http.get(`${BASE_URL}/healthz/liveness`);
        check(res, {
            'liveness is status 200': (r) => r.status === 200,
        });

        res = http.get(`${BASE_URL}/healthz/readiness`);
        check(res, {
            'readiness is status 200': (r) => r.status === 200,
        });
    });

    // res = http.get(`${BASE_URL}/`);
    // check(res, {
    //     'main page loads': (r) => r.status === 200,
    // });

    // let payload = { item_id: '1', quantity: '2' };
    // res = http.post(`${BASE_URL}/add-item`, payload);
    // check(res, {
    //     'add item redirects': (r) => r.status === 302,
    // });

    // res = http.get(`${BASE_URL}/remove-cart-item?cart_item_id=1`);
    // check(res, {
    //     'remove item redirects': (r) => r.status === 302,
    // });

    sleep(1);
}