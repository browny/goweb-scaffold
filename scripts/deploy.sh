#!/bin/bash
# Usage:
# bash scripts/deploy.sh <target_host_name> <zone> <env>
# bash scripts/deploy.sh godemo asia-east1-b staging

TARGET=$1
ZONE=$2
ENV=$3

tar zcf ./app.tar.gz --exclude .git --exclude *.gz --exclude tags .

# Upload artifact
gcloud compute copy-files --zone=$ZONE app.tar.gz $TARGET:~/
  
# Deploy
gcloud compute ssh --zone "$ZONE" $TARGET --command "rm -rf ./app; mkdir app; tar zxf app.tar.gz -C ./app; cd app; bash scripts/docker-run.sh $ENV"


