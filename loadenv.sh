#! /bin/bash

if [ -f .env ]; then
    export $(cat .env | xargs)
else
    echo "ERROR: Environments file not exists"
    exit 1
fi
