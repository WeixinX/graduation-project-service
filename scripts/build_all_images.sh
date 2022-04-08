#!/bin/bash

# this script execution will build all images by using the service binaries
# in the minikube environment

cd ..

DOCKER_ENV=$1
PRO_PATH=$(pwd)
DF_HOME="$PRO_PATH/manifests/docker"
REPLICA=3

if [ "$DOCKER_ENV" == "minikube" ]; then
    # minikube environment
    echo "script execution in the minikube docker-env environment..."
    eval "$(minikube docker-env)"
else
    echo "script execution in the host docker environment..."
fi


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

    # build image
    echo "build image $instance"
    docker image build -t "$instance" -f "$DF_HOME/$name/$num.Dockerfile" .

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

  # build image
  echo "build image $name"
  docker image build -t "$name" -f "$DF_HOME/lb-service/$name.Dockerfile" .

done