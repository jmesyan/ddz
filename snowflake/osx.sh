#!/usr/bin/env bash

# --net=host, bind the container to the host interface
# so other container can access this container via localhost:port
# ps. --net=host is not working on macOS, see https://github.com/docker/for-mac/issues/68

docker run -d -p 2379:2379 -p 2380:2380 --name etcd quay.io/coreos/etcd \
    /usr/local/bin/etcd \
    --data-dir=data.etcd --name etcd0 \
    --initial-advertise-peer-urls http://127.0.0.1:2380 --listen-peer-urls http://127.0.0.1:2380 \
    --advertise-client-urls http://0.0.0.0:2379 --listen-client-urls http://0.0.0.0:2379 \
    --initial-cluster etcd0=http://127.0.0.1:2380 \
    --initial-cluster-state new --initial-cluster-token my-etcd-token


docker build --no-cache --rm=true -t snowflake .
curl -L -X PUT http://127.0.0.1:2379/v2/keys/seqs/userid -d value="0"
curl -L -X PUT http://127.0.0.1:2379/v2/keys/seqs/snowflake-uuid -d value="0"
docker run --name snowflake1 -e SERVICE_ID=snowflake1 -e MACHINE_ID=1 -p 40001:40001 -d -P snowflake
