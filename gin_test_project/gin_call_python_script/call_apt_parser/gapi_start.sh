#!/bin/bash

# 20241009 update
chmod +x /opt/tools/gapi_web/gapi
/opt/tools/gapi_web/gapi

# keep entrypoint always alive in docker
echo "Start gapi, keep running...\n"

while true
do
    sleep 60
done
