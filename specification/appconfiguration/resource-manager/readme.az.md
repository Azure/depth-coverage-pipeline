

## Azure CLI

These settings apply only when `--az` is specified on the command line.

``` yaml $(az) && $(target-mode) == "core"
az:
  extensions: appconfiguration
  namespace: azure.mgmt.appconfiguration
  package-name: azure-mgmt-appconfiguration
  disable-checks: true
  randomize-names: true
  test-unique-resource: true
az-output-folder: $(azure-cli-folder)/src/azure-cli/azure/cli/command_modules/appconfiguration
python-sdk-output-folder: "$(azure-cli-folder)/src/azure-cli/azure/cli/command_modules/appconfiguration/vendored_sdks/appconfiguration"
compatible-level: 'track2'
```