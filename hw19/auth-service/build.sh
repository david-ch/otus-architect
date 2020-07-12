#!/bin/bash

authservicetag=dachdev/otus-architect:hw19-auth-service
docker build -t $authservicetag app/.
docker push $authservicetag
