#!/bin/bash
BASE_DIR="$(cd "$(dirname "$0" )" && pwd )"
SRC_DIR=$BASE_DIR/src
LOCAL_BUILD_DIR=$BASE_DIR/bin
BUILD_DIR=/usr/local/bin

cd $SRC_DIR

if [[ $# -eq 0 ]]; then
    echo "Running tests..."
    go test ./...
    if [[ $rc -ne 0 ]]; then
        echo "Tests failed with exit code $rc!"
        exit $rc
    fi
    echo "Building spekkio to $LOCAL_BUILD_DIR..."
    go build -o $LOCAL_BUILD_DIR/spekkio 2> /dev/null
    rc=$?
    if [[ $rc -ne 0 ]]; then
        echo "Build failed with exit code $rc!"
        exit $rc
    fi
    echo "Done!"
elif [[ $# -eq 1 ]]; then
    case $1 in
    p | prod)
        echo "Running tests..."
        go test ./...
        if [[ $rc -ne 0 ]]; then
            echo "Tests failed with exit code $rc!"
            exit $rc
        fi
        echo "Building spekkio to $BUILD_DIR..."
        go build -o $BUILD_DIR/spekkio 2> /dev/null
        rc=$?
        if [[ $rc -ne 0 ]]; then
            echo "Build failed with exit code $rc!"
            exit $rc
        fi
        echo "Done!"
        ;;
    d | dryrun | t | test)
        echo "Dry running tests..."
        stamp=$(date +"%s")
        output=/tmp/$stamp.out
        go test -v ./... -coverprofile=$output
        go tool cover -html=$output
        echo "Done!"
        ;;
    *)
        echo "Invalid options."
        exit 1
        ;;
    esac
fi
