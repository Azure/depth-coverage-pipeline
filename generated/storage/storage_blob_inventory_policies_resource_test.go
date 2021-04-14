package storage_test

import (
    "testing"
    "context"
)

type StorageBlobInventoryPolicyResource struct{}

func TestAccStorageBlobInventoryPolicy_basic(t *testing.T) {
    data := acceptance.BuildTestData(t, "azurerm_storage_blob_inventory_policy", "test")
    r := StorageBlobInventoryPolicyResource{}
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

func TestAccStorageBlobInventoryPolicy_requiresImport(t *testing.T) {
    data := acceptance.BuildTestData(t, "azurerm_storage_blob_inventory_policy", "test")
    r := StorageBlobInventoryPolicyResource{}
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

func TestAccStorageBlobInventoryPolicy_complete(t *testing.T) {
    data := acceptance.BuildTestData(t, "azurerm_storage_blob_inventory_policy", "test")
    r := StorageBlobInventoryPolicyResource{}
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

func TestAccStorageBlobInventoryPolicy_update(t *testing.T) {
    data := acceptance.BuildTestData(t, "azurerm_storage_blob_inventory_policy", "test")
    r := StorageBlobInventoryPolicyResource{}
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

func TestAccStorageBlobInventoryPolicy_updatePolicy(t *testing.T) {
    data := acceptance.BuildTestData(t, "azurerm_storage_blob_inventory_policy", "test")
    r := StorageBlobInventoryPolicyResource{}
    data.ResourceTest(t, r, []resource.TestStep{
        {
            Config: r.complete(data),
            Check: resource.ComposeTestCheckFunc(
                check.That(data.ResourceName).ExistsInAzure(r),
            ),
        },
        data.ImportStep(),
        {
            Config: r.updatePolicy(data),
            Check: resource.ComposeTestCheckFunc(
                check.That(data.ResourceName).ExistsInAzure(r),
            ),
        },
        data.ImportStep(),
    })
}

func (r StorageBlobInventoryPolicyResource) Exists(ctx context.Context, client *clients.Client, state *terraform.InstanceState) (*bool, error) {
    id, err := parse.StorageBlobInventoryPolicyID(state.ID)
    if err != nil {
        return nil, err
    }
    resp, err := client.Storage.BlobInventoryPolicyClient.Get(ctx, id.ResourceGroup, id.AccountName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            return utils.Bool(false), nil
        }
        return nil, fmt.Errorf("retrieving Storage BlobInventoryPolicy %q (Resource Group %q / accountName %q): %+v", id.Name, id.ResourceGroup, id.AccountName, err)
    }
    return utils.Bool(true), nil
}

func (r StorageBlobInventoryPolicyResource) template(data acceptance.TestData) string {
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

func (r StorageBlobInventoryPolicyResource) basic(data acceptance.TestData) string {
    template := r.template(data)
    return fmt.Sprintf(`
%s

resource "azurerm_storage_blob_inventory_policy" "test" {
  name = "acctest-sbip-%d"
  resource_group_name = azurerm_resource_group.test.name
  account_name = azurerm_storage_account.test.name
}
`, template, data.RandomInteger)
}

func (r StorageBlobInventoryPolicyResource) requiresImport(data acceptance.TestData) string {
    config := r.basic(data)
    return fmt.Sprintf(`
%s

resource "azurerm_storage_blob_inventory_policy" "import" {
  name = azurerm_storage_blob_inventory_policy.test.name
  resource_group_name = azurerm_storage_blob_inventory_policy.test.resource_group_name
  account_name = azurerm_storage_blob_inventory_policy.test.account_name
}
`, config)
}

func (r StorageBlobInventoryPolicyResource) complete(data acceptance.TestData) string {
    template := r.template(data)
    return fmt.Sprintf(`
%s

resource "azurerm_storage_blob_inventory_policy" "test" {
  name = "acctest-sbip-%d"
  resource_group_name = azurerm_resource_group.test.name
  account_name = azurerm_storage_account.test.name
  policy {
    destination = "containerName"
    enabled = true
    rules {
      name = "inventoryPolicyRule1"
      definition {
        filters {
          blob_types = ["blockBlob", "appendBlob", "pageBlob"]
          include_blob_versions = true
          include_snapshots = true
          prefix_match = ["inventoryprefix1", "inventoryprefix2"]
        }
      }
      enabled = true
    }
    type = "Inventory"
  }
}
`, template, data.RandomInteger)
}

func (r StorageBlobInventoryPolicyResource) updatePolicy(data acceptance.TestData) string {
    template := r.template(data)
    return fmt.Sprintf(`
%s

resource "azurerm_storage_blob_inventory_policy" "test" {
  name = "acctest-sbip-%d"
  resource_group_name = azurerm_resource_group.test.name
  account_name = azurerm_storage_account.test.name
  policy {
    destination = "containerName"
    enabled = true
    rules {
      name = "inventoryPolicyRule1"
      definition {
        filters {
          blob_types = ["blockBlob", "appendBlob", "pageBlob"]
          include_blob_versions = true
          include_snapshots = true
          prefix_match = ["inventoryprefix1", "inventoryprefix2"]
        }
      }
      enabled = true
    }
    type = "Inventory"
  }
}
`, template, data.RandomInteger)
}
