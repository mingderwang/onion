#!/bin/bash

#set -o xtrace
#set -e
export GOPATH=/opt/go
export GOROOT=/usr/local/go
export PATH=$GOPATH/bin:$GOROOT/bin:/home/mwang/google-cloud-sdk/bin:$PATH

typeset -i variable=$(cat version)
((variable=variable+1))
echo $variable > version

cd $1

FILE=Makefile

if [ -f $FILE ];
then
   make clean
   echo "File $FILE exists."
fi


sudo GOPATH=/opt/go /usr/local/go/bin/go get github.com/mingderwang/ginger
go generate

docker build -t asia.gcr.io/winter-wonder-647/demo-onion:v$variable .
gcloud docker push asia.gcr.io/winter-wonder-647/demo-onion:v$variable
/home/mwang/google-cloud-sdk/bin/kubectl rolling-update demo-onion --poll-interval="1us" --update-period="1us" --image=asia.gcr.io/winter-wonder-647/demo-onion:v$variable

