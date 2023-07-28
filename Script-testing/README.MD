# Unit testing

Unit testing is done via creating multiple scripts in go and postman. 

## Accepting HTTP request 

This test was to check if the Hertz HTTP Server was up and running and capable of responding to HTTP requests. Test script was written in javascript and tested via postman. 

Test was get requests sent to `http://127.0.0.1:8080/ping` endpoint after running the http server.

```javascript
// Test for response "message": "pong" 100 times
for (let i = 0; i < 100; i++) {
    pm.test(`Test ${i + 1}: Response message is 'pong'`, function () {
        pm.expect(pm.response.json().message).to.equal('pong');
    });
}
```

All 100 tests passed indicating that HTTP server was running and able to respond to requests. 

See image below for test results

![Postman results](images/postman-test.png)

## Load Balancing

The nature of this feature means the best way to test it would be to run the gateway with load balancing, and then without, and compare statistics. 

1. We used a testing tool called `locust` to help us with this. Locust creates a `swarm` of users that spam the selected endpoint(s) with queries and then presents statistics of the tests. 
2. Details of test:
    1. We created 15 multiplication servers Rate Limited to 2000 connections and 200 QPS
    2. Test was ran with 5000 users with 1000 spawn rate
    3. We tested for the average response rate with/without the load balancer in place 
3. Results:
    1. Test with load balancer
