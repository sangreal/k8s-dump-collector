#!/bin/bash
docker build -t 165463520094.dkr.ecr.ap-northeast-1.amazonaws.com/sn-dump-collector:$1 -f docker/Dockerfile.python .
docker push 165463520094.dkr.ecr.ap-northeast-1.amazonaws.com/sn-dump-collector:$1
