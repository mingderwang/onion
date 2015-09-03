#!/bin/bash
cd $1

go generate
#don't build docker locally as following script.
#./dockerize.sh

typeset -i variable=$(cat version)
((variable=variable+1))
echo $variable > version

docker build -t gcr.io/mitac-cust-gcp-1/demo-onion:v$variable .
gcloud docker push gcr.io/mitac-cust-gcp-1/demo-onion:v$variable
kubectl rolling-update demo-node --image=gcr.io/mitac-cust-gcp-1/demo-onion:v$variable

