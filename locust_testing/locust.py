import time
from locust import HttpUser, task, between

# To run locust 
# locust -f locust.py --host http://localhost:3000 --users 5000 --spawn-rate 20


# HttpUser is the user agent that can visit different cdend points
class WebsiteUser(HttpUser):
    '''Basic load testing '''
    # This is how long (s) it takes after it visited an endpoint  
    wait_time = between(1, 5)

    # Task is the rows which shows the performance of that request in locust
    @task
    def index_page(self):
        self.client.get(url="/hello")

    # Slow get request example
    @task
    def slow_page(self):
        self.client.get(url="/slow")

    '''
    post request example
    @task
    def post_req(self):
        self.client.post("/login", json={"username":"foo", "password":"bar"})
    '''

    '''Authentication testing '''
    # Boilerplate for storing token in a variable 
    def __init__(self, parent):
        super(WebsiteUser, self).__init__(parent)
        self.token = ""

    wait_time = between(1, 2)
    
    # Login to get token value 
    def on_start(self):
        with self.client.get(url="/login") as response:
            self.token = response.json()["token"]

    # Access /secret with token obtain in on_start method 
    @task
    def secret_page(self):
        self.client.get(url="/secret", headers={"authorization": self.token})
    
    # test for failed login
    @task
    def failed_login(self):
        self.client.get(url="/secret", headers={"authorization": "wrong-token"})