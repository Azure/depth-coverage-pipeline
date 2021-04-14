package storage_test

import (
    "testing"
)

func TestAccStorageBlobInventoryPolicy_ci0(t *testing.T) {
    data := acceptance.BuildTestData(t, "azurerm_storage_blob_inventory_policy", "test")
    r := StorageBlobInventoryPolicyResource{}
    data.ResourceTest(t, r, []resource.TestStep{
        {
            Config: r.ci0(data),
            Check: resource.ComposeTestCheckFunc(
                check.That(data.ResourceName).ExistsInAzure(r),
            ),
        },
        data.ImportStep(),
    })
}

func (r StorageBlobInventoryPolicyResource) ci0(data acceptance.TestData) string {
    template := r.citemplate(data)
    return fmt.Sprintf(`
%s

resource "azurerm_storage_blob_inventory_policy" "test" {
  name = "acctest-sbip-%d"
  resource_group_name = azurerm_resource_group.test.name
  account_name = azurerm_storage_account.test.name
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
`, template, data.RandomInteger)
}

func (r StorageBlobInventoryPolicyResource) citemplate(data acceptance.TestData) string {
    return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_resource_group" "test" {
  name     = "acctest-storage-%d"
  location = "%s"
}

resource "azurerm_storage_account" "test" {
  name                     = "acctestsads%s"
  resource_group_name      = azurerm_resource_group.test.name
  location                 = azurerm_resource_group.test.location
  account_tier             = "Standard"
  account_replication_type = "LRS"
}
`, data.RandomInteger, data.Locations.Primary, data.RandomString)
}