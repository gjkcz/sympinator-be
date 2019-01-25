#!/bin/bash
docker build ./
# get current latest docker image name
NAME=$(docker ps -a --format "{{.Image}}" | sed -n 1p)

echo "SHSHS:$NAME SCHAME"
set -x
docker run -it -p 8080:8080 $NAME
set +x

