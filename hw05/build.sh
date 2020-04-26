#!/bin/bash

apptag=dachdev/otus-architect:hw05-app
docker build -t $apptag app/.
docker push $apptag

dbtag=dachdev/otus-architect:hw05-db-init
docker build -t $dbtag db/.
docker push $dbtag

