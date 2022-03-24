#!/bin/bash

# this script execution will start all test load balancer between demo services

cd ..

PRO_PATH=$(pwd)
LOG_PATH="$PRO_PATH/log"
CONFIG_PATH="$PRO_PATH/config/service_load_balancer/test"
SERVICE_LB_HOME="$PRO_PATH/service_load_balancer"
LB_LIST=(\
	lb_client2nginx_web \
	lb_nginx_web2unique_id \
	lb_nginx_web2user_tag \
	lb_nginx_web2media \
	lb_nginx_web2text \
	lb_text2compose_post \
	lb_compose_post2write_timeline \
	)

if [ ! -e $LOG_PATH ]; then
	mkdir $LOG_PATH && echo "dir log created!"
fi

for name in ${LB_LIST[*]}; do
	if [ ! -e "$LOG_PATH/$name.log" ]; then
		touch "$LOG_PATH/$name.log" && echo "file $name.log created!"
	fi

	echo "$name started..."
	cd $SERVICE_LB_HOME \
		&& go build -o service_load_balancer main.go\
		&& nohup ./service_load_balancer -config_file "$CONFIG_PATH/$name.json" > "$LOG_PATH/$name.log" 2>&1 &
done
