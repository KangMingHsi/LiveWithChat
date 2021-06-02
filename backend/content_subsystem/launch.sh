#!/usr/bin/env bash
set -e

BASEDIR=$(dirname "${0}")
cd ${BASEDIR}

function usage() {
    echo "Add description of the script functions here."
    echo
    echo "Syntax: scriptTemplate [-h]"
    echo "options:"
    echo "-h     Print this Help."
    echo "-i     Launch db in memory"
    echo "-m     Migrate database first"
    echo "-l     Process in local"
    echo
}

function main() {
    args=""
    while getopts "?himl" option; do
        case $option in
            m)
                go run cmd/migrate/main.go
                exit;;
            i)
                args="${args} -inmem"
                ;;
            l)
                args="${args} -local"
                ;;
            h|?)
                usage
                exit;;
        esac
    done

    go run cmd/activate/main.go ${args}
}

main $@