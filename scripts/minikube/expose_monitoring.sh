#!/bin/bash

# this script will expose dc-ubuntu01(Floating IP is 10.2.0.6, Intranet IP is 10.2.0.55)
# prometheus and grafana, run in background and ignore output

INTRA_NET_IP='10.2.0.55'
PROMETHEUS_PORT=9090
GRAFANA_PORT=3000

echo "expose prometheus and grafana..."

# prometheus
nohup kubectl port-forward --address ${INTRA_NET_IP} service/prometheus -n monitoring ${PROMETHEUS_PORT}:${PROMETHEUS_PORT} > /dev/null 2>&1 &

# grafana
nohup kubectl port-forward --address ${INTRA_NET_IP} service/grafana -n monitoring ${GRAFANA_PORT}:${GRAFANA_PORT} > /dev/null 2>&1 &