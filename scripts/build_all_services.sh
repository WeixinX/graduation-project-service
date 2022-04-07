#!/bin/bash

# this script execution will build all service (demo and lb) by "go build"

cd ..

PRO_PATH=$(pwd)
DEMO_HOME="$PRO_PATH/service_demo"
LB_HOME="$PRO_PATH/service_load_balancer"

DEMO_LIST=(\
	nginx-web \
	unique-id \
	user-tag \
	media \
	text \
	compose-post \
	write-timeline \
)

# demo service
for name in ${DEMO_LIST[*]}; do
  echo "build demo service $name"
  cd "$DEMO_HOME/$name" \
    && go mod tidy \
    && go build -tags netgo -o "$name"
done

# lb service
echo "build lb service"
cd "$LB_HOME" \
  && go mod tidy \
  && go build -tags netgo -o lb-service