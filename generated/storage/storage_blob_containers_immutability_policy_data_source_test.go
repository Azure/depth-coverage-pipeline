package storage_test

import (
    "testing"
)
type StorageBlobContainersImmutabilityPolicyDataSource struct{}
func TestAccStorageBlobContainersImmutabilityPolicyDataSource_basic(t *testing.T) {
    data := acceptance.BuildTestData(t, "data.azurerm_storage_blob_containers_immutability_policy", "test")
    r := StorageBlobContainersImmutabilityPolicyDataSource{}
    data.DataSourceTest(t, []resource.TestStep{
        {
            Config: r.basic(data),
            Check: resource.ComposeTestCheckFunc(

            ),
        },
    })
}
func(StorageBlobContainersImmutabilityPolicyDataSource) basic(data acceptance.TestData) string {
    return fmt.Sprintf(`
%s

data "azurerm_storage_blob_containers_immutability_policy" "test" {
  name = azurerm_storage_blob_containers_immutability_policy.test.name
  resource_group_name = azurerm_storage_blob_containers_immutability_policy.test.resource_group_name
  account_name = azurerm_storage_blob_containers_immutability_policy.test.account_name
  container_name = azurerm_storage_blob_containers_immutability_policy.test.container_name
}
`, StorageBlobContainersImmutabilityPolicyResource{}.basic(data))
}
