# !/usr/bin/env bash

ENDPOINT=http://localhost:4566
PROFILE=poc

# aws dynamodb list-tables --endpoint-url http://localhost:4566 --profile poc
aws dynamodb list-tables --endpoint-url ${ENDPOINT} --profile ${PROFILE}
