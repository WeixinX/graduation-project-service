#!/bin/bash

# this script will expose dc-ubuntu01(Floating IP is 10.2.0.6, Intranet IP is 10.2.0.55)
# minikube service of es and kibana, run in background and ignore output

INTRA_NET_IP='10.2.0.55'
ES_PORT=9200
KIBANA_PORT=5601

echo "expose es and kibana..."

# Elasticsearch
nohup kubectl port-forward --address ${INTRA_NET_IP} service/es -n elastic ${ES_PORT}:${ES_PORT} > /dev/null 2>&1 &

# Kibana
nohup kubectl port-forward --address ${INTRA_NET_IP} service/kibana -n elastic ${KIBANA_PORT}:${KIBANA_PORT} > /dev/null 2>&1 &