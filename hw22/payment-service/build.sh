#!/bin/bash

paymentservicetag=dachdev/otus-architect:hw22-payment-service
docker build -t $paymentservicetag app/.
docker push $paymentservicetag
