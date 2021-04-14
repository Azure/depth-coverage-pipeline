---
subcategory: "Storage"
layout: "azurerm"
page_title: "Azure Resource Manager: azurerm_storage_queue"
description: |-
  Manages a storage Queue.
---

# azurerm_storage_queue

Manages a storage Queue.

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

resource "azurerm_storage_queue" "example" {
  name = "example-queue"
  resource_group_name = azurerm_resource_group.example.name
  account_name = azurerm_storage_account.example.name
}
```

## Arguments Reference

The following arguments are supported:

* `name` - (Required) The name which should be used for this storage Queue. Changing this forces a new storage Queue to be created.

* `resource_group_name` - (Required) The name of the Resource Group where the storage Queue should exist. Changing this forces a new storage Queue to be created.

* `account_name` - (Required) The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only. Changing this forces a new storage Queue to be created.

---

* `metadata` - (Optional) A name-value pair that represents queue metadata.

## Attributes Reference

In addition to the Arguments listed above - the following Attributes are exported:

* `id` - The ID of the storage Queue.

* `approximate_message_count` - Integer indicating an approximate number of messages in the queue. This number is not lower than the actual number of messages in the queue, but could be higher.

* `type` - The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts".

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 30 minutes) Used when creating the storage Queue.
* `read` - (Defaults to 5 minutes) Used when retrieving the storage Queue.
* `update` - (Defaults to 30 minutes) Used when updating the storage Queue.
* `delete` - (Defaults to 30 minutes) Used when deleting the storage Queue.

## Import

storage Queues can be imported using the `resource id`, e.g.

```shell
terraform import azurerm_storage_queue.example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Storage/storageAccounts/account1/queueServices/default/queues/queue1
```