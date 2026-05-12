#!/bin/bash

kubectl port-forward service/excalidraw-frontend 8080:80 &
kubectl port-forward service/excalidraw-storage 8081:8081 &
kubectl port-forward service/excalidraw-room 8082:8082 &

wait
