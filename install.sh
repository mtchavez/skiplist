#!/usr/bin/env bash
#
# Install dependencies

go version | grep 1.3 > /dev/null
go get -d -v ./...
if [ $? == 0 ]; then
    go get code.google.com/p/go.tools/cmd/cover
else
    go get golang.org/x/tools/cmd/cover
fi
