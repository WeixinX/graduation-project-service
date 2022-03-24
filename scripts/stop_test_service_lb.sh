#!/bin/bash

# this script execution will stop all test load balancer between demo services

pid_list=()
pids=$(ps -ef | grep service_load_balancer | grep -v grep | awk '{print $2}')
pid_list+=($pids)

for i in ${!pid_list[*]}; do
	kill -9 ${pid_list[$i]} && echo "service_lb#$i has been stopped..."
done
