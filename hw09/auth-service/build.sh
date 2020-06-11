#!/bin/bash

userservicetag=dachdev/otus-architect:hw09-auth-service
docker build -t $userservicetag app/.
docker push $userservicetag
