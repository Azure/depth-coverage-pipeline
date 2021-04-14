package storage_test

import (
    "testing"
)
type StorageTableDataSource struct{}
func TestAccStorageTableDataSource_basic(t *testing.T) {
    data := acceptance.BuildTestData(t, "data.azurerm_storage_table", "test")
    r := StorageTableDataSource{}
    data.DataSourceTest(t, []resource.TestStep{
        {
            Config: r.basic(data),
            Check: resource.ComposeTestCheckFunc(

            ),
        },
    })
}
func(StorageTableDataSource) basic(data acceptance.TestData) string {
    return fmt.Sprintf(`
%s

data "azurerm_storage_table" "test" {
  name = azurerm_storage_table.test.name
  resource_group_name = azurerm_storage_table.test.resource_group_name
  account_name = azurerm_storage_table.test.account_name
}
`, StorageTableResource{}.basic(data))
}
