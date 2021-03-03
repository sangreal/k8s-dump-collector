#!/bin/bash
docker build -t martynwin/sn-dump-collector:$1 -f docker/Dockerfile.python .
docker push martynwin/sn-dump-collector:$1
