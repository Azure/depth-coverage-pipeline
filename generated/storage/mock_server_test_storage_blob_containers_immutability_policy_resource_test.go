package storage_test

import (
    "testing"
)

func TestAccStorageBlobContainersImmutabilityPolicy_mockserver0(t *testing.T) {
    data := acceptance.BuildTestData(t, "azurerm_storage_blob_containers_immutability_policy", "test")
    r := StorageBlobContainersImmutabilityPolicyResource{}
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

func (r StorageBlobContainersImmutabilityPolicyResource) mockserver0() string {
    return `
provider "azurerm" {
  features {}
}

resource "azurerm_storage_blob_containers_immutability_policy" "test" {
  name = "default"
  resource_group_name = "res1782"
  account_name = "sto7069"
  container_name = "container6397"
  allow_protected_append_writes = true
  immutability_period_since_creation_in_days = 3
}
`
}
