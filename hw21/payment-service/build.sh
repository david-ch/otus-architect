#!/bin/bash

paymentservicetag=dachdev/otus-architect:hw21-payment-service
docker build -t $paymentservicetag app/.
docker push $paymentservicetag
