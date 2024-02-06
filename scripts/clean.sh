#!/usr/bin/bash

(
  docker-compose -f ./db/docker-compose.yaml rm -f
  docker-compose -f ./db/docker-compose.yaml --rmi all
  docker-compose -f ./nginx/docker-compose.yaml rm -f
  docker-compose -f ./nginx/docker-compose.yaml --rmi all
  docker system prune -f
  docker rmi $(docker images -a -q)
)