#!/bin/sh

export GOOS=linux
export GOARCH=arm64
echo $GOOS $GOARCH

echo "Building binary..."
sh ./scripts/build.sh

### Building web assets dist
echo "Building web/assets/dist..."
yarn --cwd ./web/source install && yarn --cwd ./web/source ts-patch install
yarn --cwd ./web/source build

