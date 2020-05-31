#!/bin/bash

helm install nginx stable/nginx-ingress -f nginx-ingress.yaml --atomic

