#!/usr/bin/env bash
curl -L -X PUT http://127.0.0.1:2379/v2/keys/seqs/userid -d value="0"
curl -L -X PUT http://127.0.0.1:2379/v2/keys/seqs/snowflake-uuid -d value="0"
docker run --name snowflake1 -e SERVICE_ID=snowflake1 -e MACHINE_ID=1 -p 40001:40001 -d -P snowflake
