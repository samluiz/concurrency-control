#!/usr/bin/bash

# args: nginx or db

docker-compose -f ./$1/docker-compose.yaml down
docker-compose -f ./$1/docker-compose.yaml up