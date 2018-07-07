#!/bin/bash
#
# Title: build.sh
#
# Description:
# Basic commands to build and run the docker container.
#

echo "Starting build..."

docker build -t farmers-market .

echo "Launching container on localhost:8080..."

docker run -it --rm --name farmers-market-webserver -p 8080:8080 farmers-market
