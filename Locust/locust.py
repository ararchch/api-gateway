import time
from locust import HttpUser, task, between

# To run locust 
# locust -f locust.py --host http://localhost:3000 --users 5000 --spawn-rate 20


# HttpUser is the user agent that can visit different cdend points
class WebsiteUser(HttpUser):
    '''Basic load testing '''
    # This is how long (s) it takes after it visited an endpoint  
    wait_time = between(0, 1)

    #post request example
    @task
    def multiply_req(self):
        self.client.post("/multiply", json={"FirstNum":"3", "SecondNum":"2"})

    # @task
    # def addition_req(self):
    #     self.client.post("/add", json={"FirstNum":"3", "SecondNum":"2"})

    # @task
    # def division_req(self):
    #     self.client.post("/divide", json={"FirstNum":"3", "SecondNum":"2"})
    