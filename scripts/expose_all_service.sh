#!/bin/bash

# this script will expose dc-ubuntu01(Floating IP is 10.2.0.6, Intranet IP is 10.10.10.49)
# all load balancer service, run in background and ignore output

INTRA_NET_IP='10.10.10.49'
NAMESPACE="service"

LEN=28
SERVICES=(\
	nginx-web-1 \
	nginx-web-2 \
	nginx-web-3 \
	\
	unique-id-1 \
	unique-id-2 \
	unique-id-3 \
	\
	user-tag-1 \
	user-tag-2 \
	user-tag-3 \
	\
	media-1 \
	media-2 \
	media-3 \
	\
	text-1 \
	text-2 \
	text-3 \
	\
	compose-post-1 \
	compose-post-2 \
	compose-post-3 \
	\
	write-timeline-1 \
	write-timeline-2 \
	write-timeline-3 \
  \
  lb-nginx-web \
  lb-unique-id \
  lb-user-tag \
  lb-media \
  lb-text \
  lb-compose-post \
  lb-write-timeline \
)

echo "expose all service..."

for (( i = 0; i < $LEN; i++ )); do
    echo "expose ${SERVICES[i]}"
    nohup kubectl port-forward --address ${INTRA_NET_IP} service/"${SERVICES[i]}" -n "$NAMESPACE" "${PORTS[i]}":"${PORTS[i]}" > /dev/null 2>&1 &
done
