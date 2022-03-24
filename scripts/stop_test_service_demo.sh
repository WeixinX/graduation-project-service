#!/bin/bash

# this script exec will stop all test demo service

SERVICE_LIST=(\
	nginx_web \
	unique_id \
	user_tag \
	media \
	text \
	compose_post \
	write_timeline \
	)

for name in ${SERVICE_LIST[*]}; do
	pid=$(ps -ef | grep $name | grep -v grep | awk '{print $2}')

	if [ "$pid" ]; then
		kill -9 $pid && echo "$name has been stopped..."
	else
		echo "$name is not started!"
	fi
done
