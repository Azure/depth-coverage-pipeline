package storage_test

import (
    "testing"
)
type StorageQueueDataSource struct{}
func TestAccStorageQueueDataSource_basic(t *testing.T) {
    data := acceptance.BuildTestData(t, "data.azurerm_storage_queue", "test")
    r := StorageQueueDataSource{}
    data.DataSourceTest(t, []resource.TestStep{
        {
            Config: r.basic(data),
            Check: resource.ComposeTestCheckFunc(

            ),
        },
    })
}
func(StorageQueueDataSource) basic(data acceptance.TestData) string {
    return fmt.Sprintf(`
%s

data "azurerm_storage_queue" "test" {
  name = azurerm_storage_queue.test.name
  resource_group_name = azurerm_storage_queue.test.resource_group_name
  account_name = azurerm_storage_queue.test.account_name
}
`, StorageQueueResource{}.basic(data))
}
