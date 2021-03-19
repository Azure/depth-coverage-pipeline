
#!/bin/bash
cd $GOPATH/src/github.com/terraform-providers/terraform-provider-azurerm

echo -e "\e[32m[$(date -u)] LOG: [$RP]: make goimports"
make goimports
if [ "$?" != "0" ]; then
    echo -e "\e[31m[$(date -u)] ERROR: [$RP]: make goimports failed"
    exit 1
fi

echo -e "\e[32m[$(date -u)] LOG: [$RP]: make fmt"
make fmt
if [ "$?" != "0" ]; then
    echo -e "\e[31m[$(date -u)] ERROR: [$RP]: make fmt failed"
    exit 1
fi

echo -e "\e[32m[$(date -u)] LOG: [$RP]: make build"
make build
if [ "$?" != "0" ]; then
    exit 1
fi