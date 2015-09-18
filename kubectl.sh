#!/bin/bash
set -o xtrace
set -e
echo "kubectl"

docker build -t asia.gcr.io/winter-wonder-647/demo-onion:v$variable .
gcloud docker push asia.gcr.io/winter-wonder-647/demo-onion:v$variable
kubectl rolling-update demo-onion --poll-interval="1us" --update-period="1us" --image=asia.gcr.io/winter-wonder-647/demo-onion:v$variable

