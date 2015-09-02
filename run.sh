#!/bin/bash
set -o xtrace
sudo docker build -t gcr.io/mitac-cust-gcp-1/onion-node:v1 .
sudo docker run -d -p 8080:8080 gcr.io/mitac-cust-gcp-1/onion-node:v1
