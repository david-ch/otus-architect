#!/bin/bash

authservicetag=dachdev/otus-architect:hw21-auth-service
docker build -t $authservicetag app/.
docker push $authservicetag
