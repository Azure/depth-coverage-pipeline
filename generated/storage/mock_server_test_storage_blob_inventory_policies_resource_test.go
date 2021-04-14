package storage_test

import (
    "testing"
)

func TestAccStorageBlobInventoryPolicy_mockserver0(t *testing.T) {
    data := acceptance.BuildTestData(t, "azurerm_storage_blob_inventory_policy", "test")
    r := StorageBlobInventoryPolicyResource{}
    data.ResourceTest(t, r, []resource.TestStep{
        {
            Config: r.mockserver0(),
            Check: resource.ComposeTestCheckFunc(
                check.That(data.ResourceName).ExistsInAzure(r),
            ),
        },
        data.ImportStep(),
    })
}

func (r StorageBlobInventoryPolicyResource) mockserver0() string {
    return `
provider "azurerm" {
  features {}
}

resource "azurerm_storage_blob_inventory_policy" "test" {
  name = "default"
  resource_group_name = "res7687"
  account_name = "sto9699"
  policy {
    destination = "containerName"
    enabled = true
    rules {
      name = "inventoryPolicyRule1"
      definition {
        filters {
          blob_types = ["blockBlob", "appendBlob", "pageBlob"]
          include_blob_versions = true
          include_snapshots = true
          prefix_match = ["inventoryprefix1", "inventoryprefix2"]
        }
      }
      enabled = true
    }
    type = "Inventory"
  }
}
`
}
