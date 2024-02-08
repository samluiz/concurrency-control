#!/usr/bin/bash

(
  docker-compose -f ./docker/docker-compose.yaml rm -f
  docker system prune -f --volumes
  docker rmi -f $(docker images -a -q)
)