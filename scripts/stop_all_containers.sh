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

DOCKER_ENV=$1
# Docker environment
if [ "$DOCKER_ENV" == "minikube" ]; then
    echo "script execution in the minikube docker-env environment..."
    eval "$(minikube docker-env)"

elif [ "$DOCKER_ENV" == "host" ]; then
    echo "script execution in the host docker environment..."

    # create docker network
    NET_ID=$(docker network ls | grep "$NET_NAME" | awk '{print $1}')
    if [ -z "$NET_ID" ]; then
        docker network create "$NET_NAME"
    fi

else
    echo "[ERROR] first arg can only be 'minikube' or 'host'"
    exit 1
fi

for (( i = 0; i < $LEN; i++ )); do
  echo "stop container ${SERVICES[i]}"
  CID=$(docker container ps | grep "${SERVICES[i]}" | awk '{print $1}')
  if [ -n "$CID" ]; then
      docker container stop "$CID"
  fi
done