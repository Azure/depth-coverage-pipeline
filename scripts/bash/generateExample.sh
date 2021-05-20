#!/bin/bash

set -e
set -x

if [ -z $1 ]; then
    echo "Please input resource map file path"
fi
MAP_FILE_PATH=$1

if [ -z $2 ]; then
    echo "Please input work dir path"
fi
Workspace_PATH=$2

for i in $(jq -c .jsonFileList[] ${MAP_FILE_PATH} | sed -e 's/^"//' -e 's/"$//')
do
    echo $i
    filename=$(echo $i | jq .jsonfile | sed -e 's/^"//' -e 's/"$//')
    ops=$(echo $i | jq .ops | sed -e 's/^"//' -e 's/"$//')

    if [ "$ops" === ""]; then
        oav generate-examples ${Workspace_PATH}/s/azure-rest-api-specs$filename -o=$ops
    else
        oav generate-examples ${Workspace_PATH}/s/azure-rest-api-specs$filename
    fi
done