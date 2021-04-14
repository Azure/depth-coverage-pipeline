---
subcategory: "Storage"
layout: "azurerm"
page_title: "Azure Resource Manager: azurerm_storage_blob_inventory_policy"
description: |-
  Manages a storage BlobInventoryPolicy.
---

# azurerm_storage_blob_inventory_policy

Manages a storage BlobInventoryPolicy.

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

resource "azurerm_storage_blob_inventory_policy" "example" {
  name = "example-blobinventorypolicy"
  resource_group_name = azurerm_resource_group.example.name
  account_name = azurerm_storage_account.example.name
}
```

## Arguments Reference

The following arguments are supported:

* `name` - (Required) The name which should be used for this storage BlobInventoryPolicy. Possible value is &#34;default&#34;is allowed. Changing this forces a new storage BlobInventoryPolicy to be created.

* `resource_group_name` - (Required) The name of the Resource Group where the storage BlobInventoryPolicy should exist. Changing this forces a new storage BlobInventoryPolicy to be created.

* `account_name` - (Required) The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only. Changing this forces a new storage BlobInventoryPolicy to be created.

---

* `policy` - (Optional) A `policy` block as defined below.

---

An `policy` block exports the following:

* `destination` - (Required) Container name where blob inventory files are stored. Must be pre-created.

* `enabled` - (Required) Should the policy be enabled?

* `rules` - (Required) A `rules` block as defined below.

* `type` - (Required) The valid value is Inventory.

---

An `rules` block exports the following:

* `name` - (Required) The name which should be used for this rules.

* `definition` - (Required) A `definition` block as defined below.

* `enabled` - (Required) Should the rules be enabled?

---

An `definition` block exports the following:

* `filters` - (Required) A `filters` block as defined below.

---

An `filters` block exports the following:

* `blob_types` - (Required) An array of predefined enum values. Valid values include blockBlob, appendBlob, pageBlob. Hns accounts does not support pageBlobs.

---

* `include_blob_versions` - (Optional) Includes blob versions in blob inventory when value set to true.

* `include_snapshots` - (Optional) Includes blob snapshots in blob inventory when value set to true.

* `prefix_match` - (Optional) An array of strings for blob prefixes to be matched.

## Attributes Reference

In addition to the Arguments listed above - the following Attributes are exported:

* `id` - The ID of the storage BlobInventoryPolicy.

* `last_modified_time` - Returns the last modified date and time of the blob inventory policy.

* `type` - The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts".

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 30 minutes) Used when creating the storage BlobInventoryPolicy.
* `read` - (Defaults to 5 minutes) Used when retrieving the storage BlobInventoryPolicy.
* `delete` - (Defaults to 30 minutes) Used when deleting the storage BlobInventoryPolicy.

## Import

storage BlobInventoryPolicies can be imported using the `resource id`, e.g.

```shell
terraform import azurerm_storage_blob_inventory_policy.example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Storage/storageAccounts/account1/inventoryPolicies/blobInventoryPolicy1
```