#!/bin/bash
cd $1
go generate
#don't build docker locally as following script.
#./dockerize.sh

docker build -t gcr.io/mitac-cust-gcp-1/demo-onion:v4 .
gcloud docker push gcr.io/mitac-cust-gcp-1/demo-onion:v4
