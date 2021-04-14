package storage_test

import (
    "testing"
)

func TestAccStorageTable_mockserver0(t *testing.T) {
    data := acceptance.BuildTestData(t, "azurerm_storage_table", "test")
    r := StorageTableResource{}
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

func (r StorageTableResource) mockserver0() string {
    return `
provider "azurerm" {
  features {}
}

resource "azurerm_storage_table" "test" {
  name = "table6185"
  resource_group_name = "res3376"
  account_name = "sto328"
}
`
}
