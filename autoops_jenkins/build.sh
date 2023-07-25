#!/bin/bash

source /etc/profile

go mod tidy
swag init
go build .

docker build -t autoops-jenkions:latest .
docker run -itd --net host --name autoops-jenkins -p 51000:51000 autoops-jenkions:latest 
