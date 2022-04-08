#!/bin/bash

# this script execution will delete all images
# in the minikube environment

cd ..

REPLICA=3

# minikube environment
eval "$(minikube docker-env)"

# demo service
DEMO_LIST=(\
	nginx-web \
	unique-id \
	user-tag \
	media \
	text \
	compose-post \
	write-timeline \
)

for name in ${DEMO_LIST[*]}; do
  for (( num = 1; num <= $REPLICA; num++ )); do

    instance="$name-$num"
    CID=$(docker container ps | grep "$instance" | awk '{print $1}')
    IID=$(docker images | grep "$instance" | awk '{print $3}')

    # stop container
    # remove container
    if [ -n "$CID" ]; then
      echo "stop container CID=$CID"
      docker container stop "$CID"
      docker container rm "$CID"
    fi

    # remove image
    if [ -n "$IID" ]; then
        echo "remove image IID=$IID"
        docker image rm "$IID"
    fi

  done
done


# lb service
LB_LIST=(\
	lb-nginx-web \
	lb-unique-id \
	lb-user-tag \
	lb-media \
	lb-text \
	lb-compose-post \
	lb-write-timeline \
)

for name in ${LB_LIST[*]}; do

  CID=$(docker container ps | grep "$name" | awk '{print $1}')
  IID=$(docker images | grep "$name" | awk '{print $3}')

  # stop container
  # remove container
  if [ -n "$CID" ]; then
    echo "stop and remove container CID=$CID"
    docker container stop "$CID"
    docker container rm "$CID"
  fi

  # remove image
  if [ -n "$IID" ]; then
      echo "remove image IID=$IID"
      docker image rm "$IID"
  fi

done