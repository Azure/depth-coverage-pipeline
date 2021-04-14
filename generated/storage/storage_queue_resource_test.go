package storage_test

import (
    "testing"
    "context"
)

type StorageQueueResource struct{}

func TestAccStorageQueue_basic(t *testing.T) {
    data := acceptance.BuildTestData(t, "azurerm_storage_queue", "test")
    r := StorageQueueResource{}
    data.ResourceTest(t, r, []resource.TestStep{
        {
            Config: r.basic(data),
            Check: resource.ComposeTestCheckFunc(
                check.That(data.ResourceName).ExistsInAzure(r),
            ),
        },
        data.ImportStep(),
    })
}

func TestAccStorageQueue_requiresImport(t *testing.T) {
    data := acceptance.BuildTestData(t, "azurerm_storage_queue", "test")
    r := StorageQueueResource{}
    data.ResourceTest(t, r, []resource.TestStep{
        {
            Config: r.basic(data),
            Check: resource.ComposeTestCheckFunc(
                check.That(data.ResourceName).ExistsInAzure(r),
            ),
        },
        data.RequiresImportErrorStep(r.requiresImport),
    })
}

func TestAccStorageQueue_complete(t *testing.T) {
    data := acceptance.BuildTestData(t, "azurerm_storage_queue", "test")
    r := StorageQueueResource{}
    data.ResourceTest(t, r, []resource.TestStep{
        {
            Config: r.complete(data),
            Check: resource.ComposeTestCheckFunc(
                check.That(data.ResourceName).ExistsInAzure(r),
            ),
        },
        data.ImportStep(),
    })
}

func TestAccStorageQueue_update(t *testing.T) {
    data := acceptance.BuildTestData(t, "azurerm_storage_queue", "test")
    r := StorageQueueResource{}
    data.ResourceTest(t, r, []resource.TestStep{
        {
            Config: r.basic(data),
            Check: resource.ComposeTestCheckFunc(
                check.That(data.ResourceName).ExistsInAzure(r),
            ),
        },
        data.ImportStep(),
        {
            Config: r.complete(data),
            Check: resource.ComposeTestCheckFunc(
                check.That(data.ResourceName).ExistsInAzure(r),
            ),
        },
        data.ImportStep(),
        {
            Config: r.basic(data),
            Check: resource.ComposeTestCheckFunc(
                check.That(data.ResourceName).ExistsInAzure(r),
            ),
        },
        data.ImportStep(),
    })
}

func (r StorageQueueResource) Exists(ctx context.Context, client *clients.Client, state *terraform.InstanceState) (*bool, error) {
    id, err := parse.StorageQueueID(state.ID)
    if err != nil {
        return nil, err
    }
    resp, err := client.Storage.QueueClient.Get(ctx, id.ResourceGroup, id.AccountName, id.Name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            return utils.Bool(false), nil
        }
        return nil, fmt.Errorf("retrieving Storage Queue %q (Resource Group %q / accountName %q): %+v", id.Name, id.ResourceGroup, id.AccountName, err)
    }
    return utils.Bool(true), nil
}

func (r StorageQueueResource) template(data acceptance.TestData) string {
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

func (r StorageQueueResource) basic(data acceptance.TestData) string {
    template := r.template(data)
    return fmt.Sprintf(`
%s

resource "azurerm_storage_queue" "test" {
  name = "acctest-sq-%d"
  resource_group_name = azurerm_resource_group.test.name
  account_name = azurerm_storage_account.test.name
}
`, template, data.RandomInteger)
}

func (r StorageQueueResource) requiresImport(data acceptance.TestData) string {
    config := r.basic(data)
    return fmt.Sprintf(`
%s

resource "azurerm_storage_queue" "import" {
  name = azurerm_storage_queue.test.name
  resource_group_name = azurerm_storage_queue.test.resource_group_name
  account_name = azurerm_storage_queue.test.account_name
}
`, config)
}

func (r StorageQueueResource) complete(data acceptance.TestData) string {
    template := r.template(data)
    return fmt.Sprintf(`
%s

resource "azurerm_storage_queue" "test" {
  name = "acctest-sq-%d"
  resource_group_name = azurerm_resource_group.test.name
  account_name = azurerm_storage_account.test.name
  metadata = {
  }
}
`, template, data.RandomInteger)
}
