#!/bin/bash


# create user
curl -X POST -H "Content-Type: application/json" http://127.0.0.1:8080/user -d '{"id": 1, "name": "lanxi", "age": 20, "email": "lan@qq.com", "phone": "13598457057", "address": "xxx"}'
curl -X POST -H "Content-Type: application/json" http://127.0.0.1:8080/user -d '{"id": 12, "name": "yousifu", "age": 25, "email": "yousifu@qq.com", "phone": "13598457057", "address": "xxx"}'

# get user
curl -X GET http://127.0.0.1:8080/user/1
curl -X GET http://127.0.0.1:8080/user/12
# update user
curl -X PUT -H "Content-Type: application/json" http://127.0.0.1:8080/user -d '{"id": 1, "name": "lanxi", "age": 90, "email": "lanxi@qq.com", "phone": "18800313697", "address": "yyy"}'

# get user
curl -X GET http://127.0.0.1:8080/user/1
# delete user
curl -X DELETE http://127.0.0.1:8080/user/1
# get user
curl -X GET http://127.0.0.1:8080/user/1