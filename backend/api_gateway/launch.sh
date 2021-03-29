#!/usr/bin/env bash

set -ex

BASEDIR=$(dirname "${0}")
echo "${BASEDIR}"

cd ${BASEDIR}
go run cmd/main.go ${1}