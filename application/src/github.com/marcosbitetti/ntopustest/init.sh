#!/bin/bash

set -e

go get github.com/gin-gonic/gin
go get github.com/joho/godotenv
go get github.com/rabbitmq/amqp091-go
go get github.com/stretchr/testify
go get go.mongodb.org/mongo-driver

until go run main.go; do
    >&2 echo "whaiting required services"
    sleep 1
done

