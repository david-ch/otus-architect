#!/bin/bash

tag=dachdev/otus-architect:hw03

docker build -t $tag .
docker push $tag