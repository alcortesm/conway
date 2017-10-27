#!/usr/bin/env bash
set -u -e -x -o pipefail
echo "" > coverage.txt
for d in $(go list ./...); do
    go test -race -coverprofile=profile.out -covermode=atomic $d
    if [ -f profile.out ]; then
        cat profile.out >> coverage.txt
        rm profile.out
    fi
done
