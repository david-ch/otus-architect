#!/bin/bash

apptag=dachdev/otus-architect:hw07-app
docker build -t $apptag app/.
docker push $apptag

dbtag=dachdev/otus-architect:hw07-db-init
docker build -t $dbtag db/.
docker push $dbtag

stresstesttag=dachdev/otus-architect:hw07-stress-test
docker build -t $stresstesttag stress-test/.
docker push $stresstesttag
