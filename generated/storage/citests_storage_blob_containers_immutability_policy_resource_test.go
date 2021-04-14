package storage_test

import (
    "testing"
)

func TestAccStorageBlobContainersImmutabilityPolicy_ci0(t *testing.T) {
    data := acceptance.BuildTestData(t, "azurerm_storage_blob_containers_immutability_policy", "test")
    r := StorageBlobContainersImmutabilityPolicyResource{}
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

func (r StorageBlobContainersImmutabilityPolicyResource) ci0(data acceptance.TestData) string {
    template := r.citemplate(data)
    return fmt.Sprintf(`
%s

resource "azurerm_storage_blob_containers_immutability_policy" "test" {
  name = "acctest-sbcip-%d"
  resource_group_name = azurerm_resource_group.test.name
  account_name = azurerm_storage_account.test.name
  container_name = "container6397"
  allow_protected_append_writes = true
  immutability_period_since_creation_in_days = 3
}
`, template, data.RandomInteger)
}

func (r StorageBlobContainersImmutabilityPolicyResource) citemplate(data acceptance.TestData) string {
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