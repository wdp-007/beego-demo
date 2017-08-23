#!/bin/bash

show_usage="args: [-c , -r , -s, -p, ] [--config=, --registry=, --service=, --port= ]"

while [ -n "$1" ]
do
    case "$1" in
        -c|--config) CONFIG=$2; shift 2;;
        -r|--registry) REGISTRY=$2; shift 2;;
        -s|--service) SERVICE=$2; shift 2;;
        -p|--port) PORT=$2; shift 2;;
        --) break ;;
        *) echo $1,$2,$show_usage; break ;;
    esac
done

rm -f ./kunlun
cd ./src
go version
go build -ldflags "-s -w" -o kunlun main.go
mv ./kunlun ../
cd ..

if [ ! $REGISTRY ]; then
  ./kunlun -c $CONFIG -s $SERVICE
else
  if [ ! $PORT ]; then
    ./kunlun -c $CONFIG -s $SERVICE -r $REGISTRY
  else
    ./kunlun -c $CONFIG -s $SERVICE -r $REGISTRY -p $PORT
  fi
fi
