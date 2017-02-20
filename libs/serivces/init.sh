#!/usr/bin/env bash
curl -L http://localhost:2379/v2/keys/backends/names -XPUT --data-urlencode value@names.txt
curl http://localhost:2379/v2/keys/seqs/snowflake-uuid -XPUT -d value="0"
curl http://localhost:2379/v2/keys/seqs/userid -XPUT -d value="0"

