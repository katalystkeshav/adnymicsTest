#!/bin/bash

set -x

#build image for master container
mv DockerfileMaster Dockerfile

docker build -t keepalived_master:1.0 .

mv Dockerfile DockerfileMaster

#building image for backup container
mv DockerfileBackup Dockerfile

docker build -t keepalived_backup:1.0 .

mv Dockerfile DockerfileBackup