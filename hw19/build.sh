#!/bin/bash

modules="auth-service order-service user-service"

for module in $modules; do
(
  cd "$module" || exit
  ./build.sh
)
done
