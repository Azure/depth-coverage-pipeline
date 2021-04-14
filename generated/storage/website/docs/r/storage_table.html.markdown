---
subcategory: "Storage"
layout: "azurerm"
page_title: "Azure Resource Manager: azurerm_storage_table"
description: |-
  Manages a storage Table.
---

# azurerm_storage_table

Manages a storage Table.

## Example Usage

```hcl
resource "azurerm_resource_group" "example" {
  name     = "example-storage"
  location = "West Europe"
}

resource "azurerm_storage_account" "example" {
  name                     = "examplesads"
  resource_group_name      = azurerm_resource_group.example.name
  location                 = azurerm_resource_group.example.location
  account_tier             = "Standard"
  account_replication_type = "LRS"
}

resource "azurerm_storage_table" "example" {
  name = "example-table"
  resource_group_name = azurerm_resource_group.example.name
  account_name = azurerm_storage_account.example.name
}
```

## Arguments Reference

The following arguments are supported:

* `name` - (Required) The name which should be used for this storage Table. Changing this forces a new storage Table to be created.

* `resource_group_name` - (Required) The name of the Resource Group where the storage Table should exist. Changing this forces a new storage Table to be created.

* `account_name` - (Required) The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only. Changing this forces a new storage Table to be created.

## Attributes Reference

In addition to the Arguments listed above - the following Attributes are exported:

* `id` - The ID of the storage Table.

* `table_name` - Table name under the specified account.

* `type` - The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts".

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 30 minutes) Used when creating the storage Table.
* `read` - (Defaults to 5 minutes) Used when retrieving the storage Table.
* `update` - (Defaults to 30 minutes) Used when updating the storage Table.
* `delete` - (Defaults to 30 minutes) Used when deleting the storage Table.

## Import

storage Tables can be imported using the `resource id`, e.g.

```shell
terraform import azurerm_storage_table.example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Storage/storageAccounts/account1/tableServices/default/tables/table1
```