#!/bin/bash
docker build -t martynwin/sn-oom-jar:$1 -f docker/Dockerfile .
docker push martynwin/sn-oom-jar:$1
