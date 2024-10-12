#!/bin/bash

# 20241012 only stop running gapi, centos7 test use
# find gapi pid
pid=$(ps -ef | grep "gapi_web/gapi" | grep -v "grep" | awk '{print $2}')

# kill process by pid
if [ -n "$pid" ]; then
    echo "Killing gapi process with PID: $pid"
    kill -9 $pid
else
    echo "No gapi process found to kill."
fi

