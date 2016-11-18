#!/usr/bin/env bash

version=$(cat VERSION)
compileArgs=""
runArgs=""

if [ "$1" == 'build' ]; then
    command='build'
    compileArgs="-o ansible-gen"
    echo "Compiling version $version"
else
    command='run'
    runArgs="${@:1}"
fi 

go $command $compileArgs -ldflags "-X main.cliVersion=$version" main.go $runArgs