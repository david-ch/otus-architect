#!/bin/bash

authservicetag=dachdev/otus-architect:hw16-auth-service
docker build -t $authservicetag app/.
docker push $authservicetag
