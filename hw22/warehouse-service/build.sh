#!/bin/bash

servicetag=dachdev/otus-architect:hw22-warehouse-service
docker build -t $servicetag app/.
docker push $servicetag

dbtag=dachdev/otus-architect:hw22-warehouse-db-init
docker build -t $dbtag db/.
docker push $dbtag