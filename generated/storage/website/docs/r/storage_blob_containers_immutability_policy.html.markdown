---
subcategory: "Storage"
layout: "azurerm"
page_title: "Azure Resource Manager: azurerm_storage_blob_containers_immutability_policy"
description: |-
  Manages a storage BlobContainersImmutabilityPolicy.
---

# azurerm_storage_blob_containers_immutability_policy

Manages a storage BlobContainersImmutabilityPolicy.

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

resource "azurerm_storage_blob_containers_immutability_policy" "example" {
  name = "example-blobcontainersimmutabilitypolicy"
  resource_group_name = azurerm_resource_group.example.name
  account_name = azurerm_storage_account.example.name
  container_name = "container6397"
}
```

## Arguments Reference

The following arguments are supported:

* `name` - (Required) The name which should be used for this storage BlobContainersImmutabilityPolicy. Possible value is &#34;default&#34;is allowed. Changing this forces a new storage BlobContainersImmutabilityPolicy to be created.

* `resource_group_name` - (Required) The name of the Resource Group where the storage BlobContainersImmutabilityPolicy should exist. Changing this forces a new storage BlobContainersImmutabilityPolicy to be created.

* `account_name` - (Required) The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only. Changing this forces a new storage BlobContainersImmutabilityPolicy to be created.

* `container_name` - (Required) The name of the blob container within the specified storage account. Blob container names must be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-) character must be immediately preceded and followed by a letter or number. Changing this forces a new storage BlobContainersImmutabilityPolicy to be created.

---

* `allow_protected_append_writes` - (Optional) This property can only be changed for unlocked time-based retention policies. When enabled, new blocks can be written to an append blob while maintaining immutability protection and compliance. Only new blocks can be added and any existing blocks cannot be modified or deleted. This property cannot be changed with ExtendImmutabilityPolicy API.

* `immutability_period_since_creation_in_days` - (Optional) The immutability period for the blobs in the container since the policy creation, in days.

## Attributes Reference

In addition to the Arguments listed above - the following Attributes are exported:

* `id` - The ID of the storage BlobContainersImmutabilityPolicy.

* `state` - The ImmutabilityPolicy state of a blob container, possible values include: Locked and Unlocked.

* `type` - The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts".

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 30 minutes) Used when creating the storage BlobContainersImmutabilityPolicy.
* `read` - (Defaults to 5 minutes) Used when retrieving the storage BlobContainersImmutabilityPolicy.
* `delete` - (Defaults to 30 minutes) Used when deleting the storage BlobContainersImmutabilityPolicy.

## Import

storage BlobContainersImmutabilityPolicies can be imported using the `resource id`, e.g.

```shell
terraform import azurerm_storage_blob_containers_immutability_policy.example /subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.Storage/storageAccounts/account1/blobServices/default/containers/container1/immutabilityPolicies/immutabilityPolicy1
```