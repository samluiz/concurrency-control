#!/usr/bin/bash

(
  docker-compose -f ./db/docker-compose.yaml rm -f
  docker-compose -f ./nginx/docker-compose.yaml rm -f
  docker system prune -f --volumes
  docker rmi -f $(docker images -a -q)
)