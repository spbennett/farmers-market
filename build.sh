#!/bin/bash
#
# Title: build.sh
#
# Description:
# Basic commands to build and run the docker container.
#

docker build -t farmers-market .
docker run -it --rm --name farmers-market-webserver -p 8080:8080 farmers-market