#!/bin/sh

export GOOS=linux
export GOARCH=arm64
echo $GOOS $GOARCH

sh ./scripts/build.sh
