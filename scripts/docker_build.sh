#!/bin/sh
# ===========================================================================
# File: build.sh
# Description: usage: ./build.sh [outdir]
# ===========================================================================

# exit when any command fails
set -e

cd "$(dirname "$0")/../"
. ./scripts/init.sh

#VERSION=`git describe --abbrev=0 --tags`
MODE=$1
OUTPUT_DIR=$(mkdir_output "$2")
PLATFORM=$3
OPTIONS=""
if [[ "${PLATFORM}" == "mac" ]];then
  OPTIONS="--platform=linux/amd64"
fi

# 编译
./scripts/build.sh ${OUTPUT_DIR} ${MODE} ${VERSION}

docker build ${OPTIONS} -t docker.eolinker.com/docker/apinto/apserver:${VERSION} --build-arg VERSION=${VERSION} --build-arg APP="apserver" --build-arg DIR=${OUTPUT_DIR} -f scripts/Dockerfile ./

