#!/bin/bash

curl --location --request POST 'localhost:3000/webapi/v1/user' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "nome":"Dr Rans Chucrutis",
    "sexo":"M",
    "altura":1.90,
    "peso":58,
    "imc":16.06
}'

sleep 1s

curl --location --request POST 'localhost:3000/webapi/v1/user' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "nome":"Srta Clotilde",
    "sexo":"F",
    "altura":1.70,
    "peso":43,
    "imc":14.87
}'

sleep 1s

curl --location --request POST 'localhost:3000/webapi/v1/user' \
--header 'Content-Type: text/plain' \
--data-raw '{
    "nome":"Roberval Ladrão de Chocolate",
    "sexo":"M",
    "altura":0.90,
    "peso":13,
    "imc":16.04
}'



