#!/bin/sh

export GOOS=linux
export GOARCH=amd64
echo $GOOS $GOARCH

sh ./scripts/build.sh
