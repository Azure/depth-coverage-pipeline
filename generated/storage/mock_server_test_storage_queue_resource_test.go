package storage_test

import (
    "testing"
)

func TestAccStorageQueue_mockserver0(t *testing.T) {
    data := acceptance.BuildTestData(t, "azurerm_storage_queue", "test")
    r := StorageQueueResource{}
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

func (r StorageQueueResource) mockserver0() string {
    return `
provider "azurerm" {
  features {}
}

resource "azurerm_storage_queue" "test" {
  name = "queue6185"
  resource_group_name = "res3376"
  account_name = "sto328"
}
`
}

func TestAccStorageQueue_mockserver1(t *testing.T) {
    data := acceptance.BuildTestData(t, "azurerm_storage_queue", "test")
    r := StorageQueueResource{}
    data.ResourceTest(t, r, []resource.TestStep{
        {
            Config: r.mockserver1(),
            Check: resource.ComposeTestCheckFunc(
                check.That(data.ResourceName).ExistsInAzure(r),
            ),
        },
        data.ImportStep(),
    })
}

func (r StorageQueueResource) mockserver1() string {
    return `
provider "azurerm" {
  features {}
}

resource "azurerm_storage_queue" "test" {
  name = "queue6185"
  resource_group_name = "res3376"
  account_name = "sto328"
}
`
}
