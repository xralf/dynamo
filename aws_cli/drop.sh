# !/usr/bin/env bash

ENDPOINT=http://localhost:4566
PROFILE=poc

aws dynamodb delete-table \
--endpoint-url ${ENDPOINT} \
--profile ${PROFILE} \
--table-name Music
