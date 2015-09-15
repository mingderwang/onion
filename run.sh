#!/bin/bash
set -o xtrace
docker build -t asia.gcr.io/winter-wonder-647/onion-node:v1 .
docker run -d -p 8080:8080 asia.gcr.io/winter-wonder-647/onion-node:v1
