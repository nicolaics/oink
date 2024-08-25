import sys
import requests as rq
import random

def create_users(BACKEND_ROOT):
    r = rq.post(f"http://{BACKEND_ROOT}/api/v1/user/register", json={
                        "name" : "Adam adam",
                        "email" : "Adam@gmail.com",
                        "password" : "adampassword"
                      })
    r = rq.post(f"http://{BACKEND_ROOT}/api/v1/user/register", json={
                        "name" : "Bob bob",
                        "email" : "Bob@gmail.com",
                        "password" : "bobpassword"
                      })
    r = rq.post(f"http://{BACKEND_ROOT}/api/v1/user/register", json={
                        "name" : "Charlie charlie",
                        "email" : "Charlie@gmail.com",
                        "password" : "charliepassword"
                      })
    r = rq.post(f"http://{BACKEND_ROOT}/api/v1/user/register", json={
                        "name" : "Delta delta",
                        "email" : "Delta@gmail.com",
                        "password" : "deltapassword"
                      })
    r = rq.post(f"http://{BACKEND_ROOT}/api/v1/user/register", json={
                        "name" : "Adam adam Two",
                        "email" : "AdamTwo@gmail.com",
                        "password" : "adamtwopassword"
                      })
    r = rq.post(f"http://{BACKEND_ROOT}/api/v1/user/register", json={
                        "name" : "Bob bob Two",
                        "email" : "BobTwo@gmail.com",
                        "password" : "bobtwopassword"
                      })
    r = rq.post(f"http://{BACKEND_ROOT}/api/v1/user/register", json={
                        "name" : "Charlie charlie Two",
                        "email" : "CharlieTwo@gmail.com",
                        "password" : "charlietwopassword"
                      })
    r = rq.post(f"http://{BACKEND_ROOT}/api/v1/user/register", json={
                        "name" : "Delta delta Two",
                        "email" : "DeltaTwo@gmail.com",
                        "password" : "deltatwopassword"
                      })
    r = rq.post(f"http://{BACKEND_ROOT}/api/v1/user/register", json={
                        "name" : "Zeta Zombo",
                        "email" : "ZetaZombo@gmail.com",
                        "password" : "zetazombopassword"
                      })
    r = rq.post(f"http://{BACKEND_ROOT}/api/v1/user/register", json={
                        "name" : "Zeta Zombo Two",
                        "email" : "ZetaZomboTwo@gmail.com",
                        "password" : "zetazombotwopassword"
                      })

def create_transaction(BACKEND_ROOT):
    for j in range(0, 100):
        for i in range(0, 10+1):
            r = rq.post(f"http://{BACKEND_ROOT}/api/v1/transaction/create", json={
                                "userid" : i,
                                "amount" : random.randint(1000, 10000)
                              })

def print_users(BACKEND_ROOT):
    r = rq.get(f"http://{BACKEND_ROOT}/api/v1/user/leaderboard")
    print(r.json())

def main():
    BACKEND_ROOT = input("enter backend host:port: ")
    create_users(BACKEND_ROOT)
    create_transaction(BACKEND_ROOT)
    print_users(BACKEND_ROOT)

if __name__ == "__main__":
    main()

