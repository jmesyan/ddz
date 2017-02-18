#!/usr/bin/env bash
docker rm -f agent1
docker run --rm=true --name agent1 -h agent_dev -it -p 8888:8888 -p 6060:6060 -e SERVICE_ID=agent1 agent
