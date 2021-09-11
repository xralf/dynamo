# !/usr/bin/env bash

ENDPOINT=http://localhost:4566
PROFILE=poc

aws dynamodb create-table \
--endpoint-url ${ENDPOINT} \
--profile ${PROFILE} \
--table-name Music \
--attribute-definitions \
  AttributeName=Artist,AttributeType=S \
  AttributeName=SongTitle,AttributeType=S \
--key-schema \
  AttributeName=Artist,KeyType=HASH \
  AttributeName=SongTitle,KeyType=RANGE \
--provisioned-throughput \
  ReadCapacityUnits=10,WriteCapacityUnits=5
