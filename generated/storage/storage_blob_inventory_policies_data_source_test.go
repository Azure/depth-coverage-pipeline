package storage_test

import (
    "testing"
)
type StorageBlobInventoryPolicyDataSource struct{}
func TestAccStorageBlobInventoryPolicyDataSource_basic(t *testing.T) {
    data := acceptance.BuildTestData(t, "data.azurerm_storage_blob_inventory_policy", "test")
    r := StorageBlobInventoryPolicyDataSource{}
    data.DataSourceTest(t, []resource.TestStep{
        {
            Config: r.basic(data),
            Check: resource.ComposeTestCheckFunc(

            ),
        },
    })
}
func(StorageBlobInventoryPolicyDataSource) basic(data acceptance.TestData) string {
    return fmt.Sprintf(`
%s

data "azurerm_storage_blob_inventory_policy" "test" {
  name = azurerm_storage_blob_inventory_policy.test.name
  resource_group_name = azurerm_storage_blob_inventory_policy.test.resource_group_name
  account_name = azurerm_storage_blob_inventory_policy.test.account_name
}
`, StorageBlobInventoryPolicyResource{}.basic(data))
}
