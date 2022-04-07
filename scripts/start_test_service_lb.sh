#!/bin/bash

# this script execution will start all test load balancer between demo services

cd ..

PRO_PATH=$(pwd)
LOG_PATH="$PRO_PATH/log"
CONFIG_PATH="$PRO_PATH/config/service_load_balancer/test"
SERVICE_LB_HOME="$PRO_PATH/service_load_balancer"

LB_LIST=(\
	lb-nginx-web \
	lb-unique-id \
	lb-user-tag \
	lb-media \
	lb-text \
	lb-compose-post \
	lb-write-timeline \
	)

if [ ! -e "$LOG_PATH" ]; then
	mkdir "$LOG_PATH" && echo "dir log created!"
fi

for name in ${LB_LIST[*]}; do
	if [ ! -e "$LOG_PATH/$name.log" ]; then
		touch "$LOG_PATH/$name.log" && echo "file $name.log created!"
	fi

	echo "$name started..."
	cd "$SERVICE_LB_HOME" \
	  && go mod tidy \
		&& go build -tags netgo -o lb-service \
		&& nohup ./lb-service -config_file "$CONFIG_PATH/$name.json" > "$LOG_PATH/$name.log" 2>&1 &
done
