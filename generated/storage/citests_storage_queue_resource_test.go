package storage_test

import (
    "testing"
)

func TestAccStorageQueue_ci0(t *testing.T) {
    data := acceptance.BuildTestData(t, "azurerm_storage_queue", "test")
    r := StorageQueueResource{}
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

func (r StorageQueueResource) ci0(data acceptance.TestData) string {
    template := r.citemplate(data)
    return fmt.Sprintf(`
%s

resource "azurerm_storage_queue" "test" {
  name = "acctest-sq-%d"
  resource_group_name = azurerm_resource_group.test.name
  account_name = azurerm_storage_account.test.name
}
`, template, data.RandomInteger)
}

func TestAccStorageQueue_ci1(t *testing.T) {
    data := acceptance.BuildTestData(t, "azurerm_storage_queue", "test")
    r := StorageQueueResource{}
    data.ResourceTest(t, r, []resource.TestStep{
        {
            Config: r.ci1(data),
            Check: resource.ComposeTestCheckFunc(
                check.That(data.ResourceName).ExistsInAzure(r),
            ),
        },
        data.ImportStep(),
    })
}

func (r StorageQueueResource) ci1(data acceptance.TestData) string {
    template := r.citemplate(data)
    return fmt.Sprintf(`
%s

resource "azurerm_storage_queue" "test" {
  name = "acctest-sq-%d"
  resource_group_name = azurerm_resource_group.test.name
  account_name = azurerm_storage_account.test.name
}
`, template, data.RandomInteger)
}

func (r StorageQueueResource) citemplate(data acceptance.TestData) string {
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