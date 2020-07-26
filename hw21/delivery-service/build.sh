#!/bin/bash

servicetag=dachdev/otus-architect:hw21-delivery-service
docker build -t $servicetag app/.
docker push $servicetag

dbtag=dachdev/otus-architect:hw21-delivery-db-init
docker build -t $dbtag db/.
docker push $dbtag