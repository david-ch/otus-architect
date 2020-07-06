#!/bin/bash

modules="auth-service product-service user-service"

for module in $modules; do
(
  cd "$module" || exit
  ./build.sh
)
done
