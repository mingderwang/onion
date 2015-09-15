#!/bin/bash
set -o xtrace

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


time docker build -t asia.gcr.io/winter-wonder-647/demo-onion:v$variable .
time gcloud docker push asia.gcr.io/winter-wonder-647/demo-onion:v$variable
kubectl rolling-update demo-node --poll-interval="1us" --update-period="1us" --image=asia.gcr.io/winter-wonder-647/demo-onion:v$variable

