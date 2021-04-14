---
subcategory: "Storage"
layout: "azurerm"
page_title: "Azure Resource Manager: Data Source: azurerm_storage_blob_inventory_policy"
description: |-
  Gets information about an existing storage BlobInventoryPolicy.
---

# Data Source: azurerm_storage_blob_inventory_policy

Use this data source to access information about an existing storage BlobInventoryPolicy.

## Example Usage

```hcl
data "azurerm_storage_blob_inventory_policy" "example" {
  name = "example-blobinventorypolicy"
  resource_group_name = "existing"
  account_name = "existing"
}

output "id" {
  value = data.azurerm_storage_blob_inventory_policy.example.id
}
```

## Arguments Reference

The following arguments are supported:

* `name` - (Required) The name of this storage BlobInventoryPolicy. Possible value is &#34;default&#34;is allowed.

* `resource_group_name` - (Required) The name of the Resource Group where the storage BlobInventoryPolicy exists.

* `account_name` - (Required) The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.

## Attributes Reference

In addition to the Arguments listed above - the following Attributes are exported:

* `id` - The ID of the storage BlobInventoryPolicy.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `read` - (Defaults to 5 minutes) Used when retrieving the storage BlobInventoryPolicy.