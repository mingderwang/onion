#!/bin/bash

typeset -i variable=$(cat version)
((variable=variable+1))
echo $variable > version

cd $1

FILE=Makefile

if [ -f $FILE ];
then
   echo "File $FILE exists."
   make clean
fi


go generate
#don't build docker locally as following script.
#./dockerize.sh


docker build -t gcr.io/mitac-cust-gcp-1/demo-onion:v$variable .
gcloud docker push gcr.io/mitac-cust-gcp-1/demo-onion:v$variable
kubectl rolling-update demo-node --image=gcr.io/mitac-cust-gcp-1/demo-onion:v$variable

