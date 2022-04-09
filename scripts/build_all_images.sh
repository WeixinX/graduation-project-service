#!/bin/bash

# this script execution will build all images by using the service binaries
# in the minikube environment

cd ..

DOCKER_ENV=$1
#TEST_CONFIG=$2

PRO_PATH=$(pwd)
DF_HOME="$PRO_PATH/manifests/docker"
REPLICA=3

# Docker environment
if [ "$DOCKER_ENV" == "minikube" ]; then
    echo "script execution in the minikube docker-env environment..."
    eval "$(minikube docker-env)"

elif [ "$DOCKER_ENV" == "host" ]; then
    echo "script execution in the host docker environment..."

else
    echo "[ERROR] first arg can only be 'minikube' or 'host'"
    exit 1
fi

# Test configure
# the difference between 'test' and 'no test' is that
# the 'host' of 'call_url' in the configuration file is 'localhost'
#if [ "$TEST_CONFIG" == "test" ]; then
#    echo "script execution using test configure..."
#    DF_HOME="$PRO_PATH/manifests/test/docker"
#
#elif [ "$TEST_CONFIG" == "no_test" ]; then
#    echo "script execution no using test configure..."
#    DF_HOME="$PRO_PATH/manifests/docker"
#
#else
#    echo "[ERROR] second arg can only be 'test' or 'no_test'"
#    exit 1
#fi


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