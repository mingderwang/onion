#!/bin/bash
set -o xtrace
set -e

echo "ming"

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

export GOPATH=/opt/go
echo $GOPATH
#/usr/local/go/bin/go generate
