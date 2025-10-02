#!/bin/bash

if [ -f .env ]; then
    source .env
fi

cd sql/schema
export GOOSE_DRIVER=turso
export GOOSE_DBSTRING=$DATABASE_URL
goose up
