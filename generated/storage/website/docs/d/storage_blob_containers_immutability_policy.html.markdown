---
subcategory: "Storage"
layout: "azurerm"
page_title: "Azure Resource Manager: Data Source: azurerm_storage_blob_containers_immutability_policy"
description: |-
  Gets information about an existing storage BlobContainersImmutabilityPolicy.
---

# Data Source: azurerm_storage_blob_containers_immutability_policy

Use this data source to access information about an existing storage BlobContainersImmutabilityPolicy.

## Example Usage

```hcl
data "azurerm_storage_blob_containers_immutability_policy" "example" {
  name = "example-blobcontainersimmutabilitypolicy"
  resource_group_name = "existing"
  account_name = "existing"
  container_name = "existing"
}

output "id" {
  value = data.azurerm_storage_blob_containers_immutability_policy.example.id
}
```

## Arguments Reference

The following arguments are supported:

* `name` - (Required) The name of this storage BlobContainersImmutabilityPolicy. Possible value is &#34;default&#34;is allowed.

* `resource_group_name` - (Required) The name of the Resource Group where the storage BlobContainersImmutabilityPolicy exists.

* `account_name` - (Required) The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.

* `container_name` - (Required) The name of the blob container within the specified storage account. Blob container names must be between 3 and 63 characters in length and use numbers, lower-case letters and dash (-) only. Every dash (-) character must be immediately preceded and followed by a letter or number.

## Attributes Reference

In addition to the Arguments listed above - the following Attributes are exported:

* `id` - The ID of the storage BlobContainersImmutabilityPolicy.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `read` - (Defaults to 5 minutes) Used when retrieving the storage BlobContainersImmutabilityPolicy.