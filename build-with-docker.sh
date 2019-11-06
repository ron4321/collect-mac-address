#!/bin/sh
set -x

docker run \
--rm \
-v "$PWD":/app \
-w /app \
golang:1.13-alpine /bin/sh build.sh collect-mac-address