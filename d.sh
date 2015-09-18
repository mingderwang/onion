#!/bin/bash
set -o xtrace
set -e

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


