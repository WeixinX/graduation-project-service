#!/bin/bash

# this script will expose dc-ubuntu01(Floating IP is 10.2.0.6, Intranet IP is 10.2.0.55)
# jaeger-query and collector, run in background and ignore output

INTRA_NET_IP='10.2.0.55'
QUERY_PORT=16686
COLLECTOR_PORT=14250

echo "expose jaeger-query and collector..."

# query
nohup kubectl port-forward --address ${INTRA_NET_IP} service/jaeger-query -n tracing ${QUERY_PORT}:${QUERY_PORT} > /dev/null 2>&1 &

# collector
nohup kubectl port-forward --address ${INTRA_NET_IP} service/jaeger-collector -n tracing ${COLLECTOR_PORT}:${COLLECTOR_PORT} > /dev/null 2>&1 &