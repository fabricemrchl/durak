#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd)"

cd "${DIR}/.."
echo $(( $(cat version) + 1 )) > version
git pull && docker-compose -p durak -f deployments/docker-compose.prod.yml up -d --build
cd -
