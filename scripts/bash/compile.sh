
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
make build >> $GOPATH/../build.log
if [ "$?" != "0" ]; then
    echo -e "\e[31m[$(date -u)] ERROR: [$RP]: make build failed"
    echo $GOPATH
    cat $GOPATH/../build.log
    exit 1
fi