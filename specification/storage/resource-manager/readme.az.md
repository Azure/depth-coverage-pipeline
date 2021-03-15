

## Azure CLI

These settings apply only when `--az` is specified on the command line.

``` yaml $(az)
az:
  extensions: storage
  namespace: azure.mgmt.storage
  package-name: azure-mgmt-storage
  disable-checks: true
  randomize-names: true
  test-unique-resource: true
az-output-folder: $(azure-cli-folder)/src/azure-cli/azure/cli/command_modules/storage
python-sdk-output-folder: "$(azure-cli-folder)/src/azure-cli/azure/cli/command_modules/storage/vendored_sdks/storage"
compatible-level: track2
```

### -----start of auto generated cli-directive----- ###
``` yaml $(az)
cli:
  cli-directive:
    - where:
        group: '*'
        op: '*'
      hidden: true
    - where:
        group: 'BlobInventoryPolicies'
        op: CreateOrUpdate
        apiVersion: '2019-06-01'
      hidden: false
    - where:
        group: 'Table'
        op: Delete
        apiVersion: '2019-06-01'
      hidden: false
```
### -----end of auto generated cli-directive----- ###