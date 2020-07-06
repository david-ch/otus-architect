#!/bin/bash

helm install prom stable/prometheus-operator -f prometheus.yaml --atomic
