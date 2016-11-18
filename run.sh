#!/usr/bin/env bash

version=$(cat VERSION)
args=""

if [ "$1" == 'build' ]; then
    command='build'
    args="-o ansible-gen"
    echo "Compiling version $version"
else
    command='run'
fi 

go $command $args -ldflags "-X main.cliVersion=$version" main.go