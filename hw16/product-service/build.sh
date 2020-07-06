#!/bin/bash

productservicetag=dachdev/otus-architect:hw16-product-service
docker build -t $productservicetag app/.
docker push $productservicetag

productdbtag=dachdev/otus-architect:hw16-product-db-init
docker build -t $productdbtag db/.
docker push $productdbtag
