#!/bin/bash
docker build -t 165463520094.dkr.ecr.ap-northeast-1.amazonaws.com/sn-oom-jar:$1 -f docker/Dockerfile .
docker push 165463520094.dkr.ecr.ap-northeast-1.amazonaws.com/sn-oom-jar:$1
