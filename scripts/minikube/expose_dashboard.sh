#!/bin/bash

# this script will expose dc-ubuntu01(Floating IP is 10.2.0.6, Intranet IP is 10.2.0.55)
# minikube dashboard, run in background and ignore output

INTRA_NET_IP='10.2.0.55'
EXPOSE_PORT=8001

echo "run and expose dashboard..."

# run dashboard
nohup minikube dashboard > /dev/null 2>&1 &

# expose dashboard
nohup kubectl proxy --port=${EXPOSE_PORT} --address=${INTRA_NET_IP} --accept-hosts='^.*' > /dev/null 2>&1 &