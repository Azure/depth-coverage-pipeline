---
subcategory: "Storage"
layout: "azurerm"
page_title: "Azure Resource Manager: Data Source: azurerm_storage_queue"
description: |-
  Gets information about an existing storage Queue.
---

# Data Source: azurerm_storage_queue

Use this data source to access information about an existing storage Queue.

## Example Usage

```hcl
data "azurerm_storage_queue" "example" {
  name = "example-queue"
  resource_group_name = "existing"
  account_name = "existing"
}

output "id" {
  value = data.azurerm_storage_queue.example.id
}
```

## Arguments Reference

The following arguments are supported:

* `name` - (Required) The name of this storage Queue.

* `resource_group_name` - (Required) The name of the Resource Group where the storage Queue exists.

* `account_name` - (Required) The name of the storage account within the specified resource group. Storage account names must be between 3 and 24 characters in length and use numbers and lower-case letters only.

## Attributes Reference

In addition to the Arguments listed above - the following Attributes are exported:

* `id` - The ID of the storage Queue.

## Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `read` - (Defaults to 5 minutes) Used when retrieving the storage Queue.