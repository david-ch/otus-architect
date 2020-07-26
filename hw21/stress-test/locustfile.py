import random
from faker import Faker
from locust import HttpUser, task, between, TaskSet
from urllib.parse import urlencode


common_headers = {'Host': 'arch.homework'}


class SearchProductTaskSet(TaskSet):
    wait_time = between(2, 6)
    search_letters = "цукенгшзхфывапролджэячсмитбюЦУКЕНГШЗФЫВАПРОЛДЖЭЯЧСМИТБЮ1234567890"
    ids = []

    @task(5)
    def get_product(self):
        if not self.ids:
            return

        self.client.get(
            url="/product/" + str(random.choice(self.ids)),
            headers=common_headers
        )

    @task(1)
    def end(self):
        self.interrupt()

    def on_start(self):
        self.ids = self.client.get(
            url="/product?" + urlencode({'search': random.choice(self.search_letters)}),
            headers=common_headers
        ).json()

        if not self.ids:
            self.interrupt()


class BrowsePromotedTaskSet(TaskSet):
    wait_time = between(2, 6)
    ids = []

    @task(5)
    def get_product(self):
        self.client.get(
            url="/product/" + str(random.choice(self.ids)),
            headers=common_headers
        )

    @task(1)
    def end(self):
        self.interrupt()

    def on_start(self):
        self.ids = self.client.get(
            url="/product/promoted",
            headers=common_headers
        ).json()

        if not self.ids:
            self.interrupt()


class QuickstartUser(HttpUser):
    wait_time = between(5, 9)
    tasks = {
        BrowsePromotedTaskSet: 2,
        SearchProductTaskSet: 1
    }

    fake = Faker('en_US')
    user_name = fake.user_name()
    password = user_name
    first_name = fake.first_name()
    last_name = fake.last_name()

    def on_start(self):
        self.client.post(
            url="/auth/register",
            json={
                "username": self.user_name,
                "firstName": self.first_name,
                "lastName": self.last_name,
                "password": self.password
            },
            headers=common_headers
        )

        self.client.post(
            url="/auth/login",
            json={
                "username": self.user_name,
                "password": self.password
            },
            headers=common_headers
        )
