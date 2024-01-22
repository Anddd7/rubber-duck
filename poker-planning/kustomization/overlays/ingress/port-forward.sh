#!/bin/bash

kubectl port-forward service/poker-planning-frontend 8080:80 &
kubectl port-forward service/poker-planning-server 8000:8000 &

wait
