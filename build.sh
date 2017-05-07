#!/bin/bash
set -e

VERSION=dev
REPO=${REPO:-husseingalal}
TAG=${TAG:-${VERSION}}
IMAGE=${REPO}/rancher-log-dashboard:${TAG}

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o rancher-log-dashboard
docker build -t ${IMAGE} .

echo Built ${IMAGE}

docker push ${IMAGE}
