#!/bin/bash

authservicetag=dachdev/otus-architect:hw15-auth-service
docker build -t $authservicetag app/.
docker push $authservicetag
