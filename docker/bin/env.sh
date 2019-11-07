#!/usr/bin/env bash

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null && pwd )"
cd $(dirname ${SCRIPT_DIR}/../..)

PROJECT_NAME=lemmas
DOCKER_FILE="-f docker-compose.yml"
COMPOSE="docker-compose -p ${PROJECT_NAME} ${DOCKER_FILE}"
