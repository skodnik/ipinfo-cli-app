#!/usr/bin/env bash

package="ipinfo"
platform=$1

platform_split=(${platform//-/ })
GOOS=${platform_split[0]}
GOARCH=${platform_split[1]}
output_name=$package'-'$GOOS'-'$GOARCH
if [ $GOOS = "windows" ]; then
    output_name+='.exe'
fi

env GOOS=$GOOS GOARCH=$GOARCH go build -o ./build/bin/$output_name ./cmd/ipinfo/ipinfo.go
if [ $? -ne 0 ]; then
    echo 'An error has occurred! Aborting the script execution...'
    exit 1
fi

echo $output_name
