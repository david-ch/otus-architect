#!/bin/bash

modules="auth-service delivery-service order-service payment-service user-service product-service warehouse-service"

for module in $modules; do
(
  cd "$module" || exit
  ./build.sh
)
done
