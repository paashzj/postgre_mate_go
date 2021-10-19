#!/bin/bash

docker run -p 5432:5432 -d --rm ttbb/postgre:mate
sleep 30
container_id=`docker ps|tail -n 1|awk '{print $1}'`
docker exec $container_id psql -U sh -d ttbb -f mate/sql/init_healthcheck.sql
