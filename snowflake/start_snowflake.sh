#!/usr/bin/env bash
docker rm -f snowflake1
docker build --no-cache --rm=true -t snowflake .
docker run --rm=true --name snowflake1 -e SERVICE_ID=snowflake1 -e MACHINE_ID=1 --net=host -p 40001:40001 -d -P snowflake
