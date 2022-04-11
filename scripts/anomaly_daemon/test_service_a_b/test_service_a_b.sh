#!/bin/bash

# this script execution will be used in the Dockerfile to let the container start multiple services

nohup ./service-b > /dev/null 2>&1 &

./service-a