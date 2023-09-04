#!/bin/bash

# Wait a bit before trying...
sleep 30

ready=0

while [ $ready -ne 1 ]; do
    kubectl get nodes | grep NotReady
    ready=$?
done

kubectl apply -f /opt/manifests
