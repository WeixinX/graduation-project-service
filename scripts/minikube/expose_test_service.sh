#!/bin/bash

# this script will expose dc-ubuntu01(Floating IP is 10.2.0.6, Intranet IP is 10.2.0.55)
# service-a, service-b, service-c, run in background and ignore output

INTRA_NET_IP='10.2.0.55'
SERVICE_A_PORT=8003
SERVICE_B_PORT=8002
SERVICE_C_PORT=8004

echo "expose test service..."

# service-a
nohup kubectl port-forward --address ${INTRA_NET_IP} service/service-a -n service ${SERVICE_A_PORT}:${SERVICE_A_PORT} > /dev/null 2>&1 &

# service-b
nohup kubectl port-forward --address ${INTRA_NET_IP} service/service-b -n service ${SERVICE_B_PORT}:${SERVICE_B_PORT} > /dev/null 2>&1 &

# service-c
nohup kubectl port-forward --address ${INTRA_NET_IP} service/service-c -n service ${SERVICE_C_PORT}:${SERVICE_C_PORT} > /dev/null 2>&1 &