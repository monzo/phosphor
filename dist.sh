#!/bin/bash

set -e

DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
rm -rf   $DIR/dist/docker
mkdir -p $DIR/dist/docker
rm -rf   $DIR/.godeps
mkdir -p $DIR/.godeps
export GOPATH=$DIR/vendor:$GOPATH

arch=$(go env GOARCH)
version=$(awk '/const Version/ {print $NF}' < $DIR/internal/version/version.go | sed 's/"//g')
goversion=$(go version | awk '{print $3}')

for os in linux darwin freebsd; do
    echo "... building v$version for $os/$arch"
    BUILD=$(mktemp -d -t phosphor)
    TARGET="phosphor-$version.$os-$arch.$goversion"
    for app in phosphor phosphord; do
        GOOS=$os GOARCH=$arch CGO_ENABLED=0 go build -o $BUILD/$TARGET/bin/$app ./apps/$app
    done
    pushd $BUILD
    if [ "$os" == "linux" ]; then
        cp -r $TARGET/bin $DIR/dist/docker/
    fi
    tar czvf $TARGET.tar.gz $TARGET
    mv $TARGET.tar.gz $DIR/dist
    popd
    rm -r $BUILD
done

docker build -t mondough/phosphor:v$version .
if [[ ! $version == *"-"* ]]; then
    echo "Tagging mondough/phosphor:v$version as the latest release."
    docker tag -f mondough/phosphor:v$version mondough/phosphor:latest
fi
