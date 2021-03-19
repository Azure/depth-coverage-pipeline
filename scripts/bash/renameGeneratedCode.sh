#!/usr/bin/env bash

set -e

if [ -z $1 ]; then
    echo "Please input rp path"
fi
RP_PATH=$1

if [ -z $2 ]; then
    echo "Please input rp path"
fi
SUFFIX=$2
# SUFFIX="forpipeline"
RP_FOLDER=`basename ${RP_PATH}`
ORIGINAL_RP_FOLDER=`echo ${RP_FOLDER} | sed "s/${SUFFIX}//g"`

CLIENT_IMPORT=`cat ${RP_PATH}/terraform-configuration/configuration_client | sed -n 1p`
CLIENT_STRUCT=`cat ${RP_PATH}/terraform-configuration/configuration_client | sed -n 2p`
CLIENT_NEW=`cat ${RP_PATH}/terraform-configuration/configuration_client | sed -n 3p`

SERVICE_IMPORT=`cat ${RP_PATH}/terraform-configuration/configuration_services | sed -n 1p`
SERVICE_REGISTRATION=`cat ${RP_PATH}/terraform-configuration/configuration_services | sed -n 2p`

# modify import in client.go
CLIENT_IMPORT=`echo ${CLIENT_IMPORT} | sed "s|${ORIGINAL_RP_FOLDER}|${RP_FOLDER}|g"`

# modify client struct in client.go
ORIGINAL_CLIENT_NAME=`echo ${CLIENT_STRUCT} | awk -F ' ' '{print $1}'`
CLIENT_NAME="${ORIGINAL_CLIENT_NAME}${SUFFIX}"
CLIENT_STRUCT=`echo ${CLIENT_STRUCT} | sed "s/${ORIGINAL_CLIENT_NAME}/${CLIENT_NAME}/g" | sed "s/${ORIGINAL_RP_FOLDER}/${RP_FOLDER}/g"`

# modify client new in client.go
CLIENT_NEW=`echo ${CLIENT_NEW} | sed "s/${ORIGINAL_CLIENT_NAME}/${CLIENT_NAME}/g" | sed "s/${ORIGINAL_RP_FOLDER}/${RP_FOLDER}/g"`

# modify import in services.go
SERVICE_IMPORT=`echo ${SERVICE_IMPORT} | sed "s|${ORIGINAL_RP_FOLDER}|${RP_FOLDER}|g"`
SERVICE_IMPORT="${RP_FOLDER} ${SERVICE_IMPORT}"

# modify registration in services.go
SERVICE_REGISTRATION=`echo ${SERVICE_REGISTRATION} | sed "s/${ORIGINAL_RP_FOLDER}/${RP_FOLDER}/g"`

# change client.go and services.go
CLIENT_IMPORT_ORI="web \"github.com\/terraform-providers\/terraform-provider-azurerm\/azurerm\/internal\/services\/web\/client\""
CLIENT_STRUCT_ORI="*web.Client"
CLIENT_NEW_ORI="client.Web = web.NewClient(o)"

SERVICE_IMPORT_ORI="\"github.com\/terraform-providers\/terraform-provider-azurerm\/azurerm\/internal\/services\/web\""
SERVICE_REGISTRATION_ORI="web.Registration{},"

sed -i "/$CLIENT_IMPORT_ORI/a $CLIENT_IMPORT" ./azurerm/internal/clients/client.go
sed -i "/$CLIENT_STRUCT_ORI/a $CLIENT_STRUCT" ./azurerm/internal/clients/client.go
sed -i "/$CLIENT_NEW_ORI/a $CLIENT_NEW" ./azurerm/internal/clients/client.go
sed -i "/$SERVICE_IMPORT_ORI/a $SERVICE_IMPORT" ./azurerm/internal/provider/services.go
sed -i "/$SERVICE_REGISTRATION_ORI/a $SERVICE_REGISTRATION" ./azurerm/internal/provider/services.go

# change the client name in resource.go, resource_test.go and datasource.go
find ${RP_PATH}/*_resource.go | sort | while read f; do sed -i "s/Client)\.${ORIGINAL_CLIENT_NAME}/Client)\.${CLIENT_NAME}/g" $f; done
find ${RP_PATH}/*_resource_test.go | sort | while read f; do sed -i "s/client\.${ORIGINAL_CLIENT_NAME}/client\.${CLIENT_NAME}/g" $f; done
find ${RP_PATH}/*_data_source.go | sort | while read f; do sed -i "s/Client)\.${ORIGINAL_CLIENT_NAME}/Client)\.${CLIENT_NAME}/g" $f; done
