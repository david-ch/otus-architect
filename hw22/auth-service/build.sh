#!/bin/bash

authservicetag=dachdev/otus-architect:hw22-auth-service
docker build -t $authservicetag app/.
docker push $authservicetag
