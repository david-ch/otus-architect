#!/bin/bash

orderservicetag=dachdev/otus-architect:hw22-order-service
docker build -t $orderservicetag app/.
docker push $orderservicetag

orderdbtag=dachdev/otus-architect:hw22-order-db-init
docker build -t $orderdbtag db/.
docker push $orderdbtag
