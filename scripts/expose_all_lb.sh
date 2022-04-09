#!/bin/bash

# this script will expose dc-ubuntu01(Floating IP is 10.2.0.6, Intranet IP is 10.10.10.49)
# all load balancer service, run in background and ignore output

INTRA_NET_IP='10.10.10.49'
NAMESPACE="service"

LEN=7
SERVICES=(\
  lb-nginx-web \
  lb-unique-id \
  lb-user-tag \
  lb-media \
  lb-text \
  lb-compose-post \
  lb-write-timeline \
)

PORTS=(\
  8011 \
  8012 \
  8013 \
  8014 \
  8015 \
  8056 \
  8067 \
)

echo "expose all balancer service..."

for (( i = 0; i < $LEN; i++ )); do
    echo "expose ${SERVICES[i]}"
    nohup kubectl port-forward --address ${INTRA_NET_IP} service/"${SERVICES[i]}" -n "$NAMESPACE" "${PORTS[i]}":"${PORTS[i]}" > /dev/null 2>&1 &
done
