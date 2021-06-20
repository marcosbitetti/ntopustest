#!/bin/bash

go get github.com/rabbitmq/amqp091-go
go get github.com/joho/godotenv
go get github.com/gorilla/websocket

go run *.go
