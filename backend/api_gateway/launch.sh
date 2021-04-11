#!/usr/bin/env bash

set -ex

BASEDIR=$(dirname "${0}")
cd ${BASEDIR}

function usage() {
    echo "Add description of the script functions here."
    echo
    echo "Syntax: scriptTemplate [-h]"
    echo "options:"
    echo "-h     Print this Help."
    echo
}

function main() {
    args=""
    while getopts "?h" option; do
        case $option in
            h|?)
                usage
                exit;;
        esac
    done

    go run cmd/main.go ${args}
}

main $@