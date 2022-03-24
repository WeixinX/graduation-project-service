#!/bin/bash

# this script execution will start all test demo service.

cd ..

PRO_PATH=$(pwd)
LOG_PATH="$PRO_PATH/log"
SERVICE_DEMO_HOME="$PRO_PATH/service_demo"
SERVICE_LIST=(\
	nginx_web \
	unique_id \
	user_tag \
	media \
	text \
	compose_post \
	write_timeline \
	)

if [ ! -e $LOG_PATH ]; then
	mkdir $LOG_PATH && echo "dir log created!"
fi

for name in ${SERVICE_LIST[*]}; do
	if [ ! -e "$LOG_PATH/$name.log" ]; then
		touch "$LOG_PATH/$name.log" && echo "file $name.log created!"
	fi

	echo "$name started..."
	cd $SERVICE_DEMO_HOME/$name \
		&& go build -o $name main.go \
		&& nohup ./$name -config_file ./config/config_test.json > "$LOG_PATH/$name.log" 2>&1 &
done
