#!/bin/bash

docker-compose down;
docker rmi wkey-stock-app;
docker-compose up -d;