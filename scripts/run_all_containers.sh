#!/bin/bash

# this script execution will run all containers by using the images
# in the minikube environment

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

PORTS=(\
	8101 \
	8201 \
	8301 \
	\
	8102 \
	8202 \
	8302 \
	\
	8103 \
	8203 \
	8303 \
	\
	8104 \
	8204 \
	8304 \
	\
	8105 \
	8205 \
	8305 \
	\
	8106 \
	8206 \
	8306 \
	\
	8107 \
	8207 \
	8307 \
  \
  8011 \
  8012 \
  8013 \
  8014 \
  8015 \
  8056 \
  8067 \
)

for (( i = 0; i < $LEN; i++ )); do
  echo "run container ${SERVICES[i]}"
  docker run --rm -d --name "${SERVICES[i]}" -p "${PORTS[i]}:${PORTS[i]}" "${SERVICES[i]}"
done