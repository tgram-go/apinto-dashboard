#!/bin/sh
# ===========================================================================
# File: init.sh
# Description: common variables & functions.
# ===========================================================================

set -e

# Global variables
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

# Invoke from project root
VERSION=`git describe --abbrev=0 --tags`

# Version function used for version string comparison
version() { echo "$@" | awk -F. '{ printf("%d%03d%03d%03d\n", $1,$2,$3,$4); }'; }

# Ensure output directory existed
mkdir_output() {
    if [ -z "$1" ]; then
        mkdir -p apinto-build
        OUTPUT_DIR=$(cd apinto-build > /dev/null && pwd)
    else
        OUTPUT_DIR="$1"
    fi
    echo "$OUTPUT_DIR"
}