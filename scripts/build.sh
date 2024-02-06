#!/usr/bin/bash

docker build -t concurrency-control -t samluiz/concurrency-control -f Dockerfile ../
docker push samluiz/concurrency-control