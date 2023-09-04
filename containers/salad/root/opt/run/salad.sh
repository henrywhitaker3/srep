#!/bin/bash

# Wait a bit before trying...
sleep 30

ready=0

while [ $ready -ne 1 ]; do
    kubectl get nodes | grep NotReady
    ready=$?
done

kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.12.4/cert-manager.yaml

sleep 15

kubectl apply -f /opt/manifests
