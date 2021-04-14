package storage_test

import (
    "testing"
    "context"
)

type StorageBlobContainersImmutabilityPolicyResource struct{}

func TestAccStorageBlobContainersImmutabilityPolicy_basic(t *testing.T) {
    data := acceptance.BuildTestData(t, "azurerm_storage_blob_containers_immutability_policy", "test")
    r := StorageBlobContainersImmutabilityPolicyResource{}
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

func TestAccStorageBlobContainersImmutabilityPolicy_requiresImport(t *testing.T) {
    data := acceptance.BuildTestData(t, "azurerm_storage_blob_containers_immutability_policy", "test")
    r := StorageBlobContainersImmutabilityPolicyResource{}
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

func TestAccStorageBlobContainersImmutabilityPolicy_complete(t *testing.T) {
    data := acceptance.BuildTestData(t, "azurerm_storage_blob_containers_immutability_policy", "test")
    r := StorageBlobContainersImmutabilityPolicyResource{}
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

func TestAccStorageBlobContainersImmutabilityPolicy_update(t *testing.T) {
    data := acceptance.BuildTestData(t, "azurerm_storage_blob_containers_immutability_policy", "test")
    r := StorageBlobContainersImmutabilityPolicyResource{}
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

func (r StorageBlobContainersImmutabilityPolicyResource) Exists(ctx context.Context, client *clients.Client, state *terraform.InstanceState) (*bool, error) {
    id, err := parse.StorageBlobContainersImmutabilityPolicyID(state.ID)
    if err != nil {
        return nil, err
    }
    resp, err := client.Storage.BlobContainerClient.GetImmutabilityPolicy(ctx, id.ResourceGroup, id.AccountName, id.ContainerName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            return utils.Bool(false), nil
        }
        return nil, fmt.Errorf("retrieving Storage BlobContainersImmutabilityPolicy %q (Resource Group %q / accountName %q / containerName %q): %+v", id.Name, id.ResourceGroup, id.AccountName, id.ContainerName, err)
    }
    return utils.Bool(true), nil
}

func (r StorageBlobContainersImmutabilityPolicyResource) template(data acceptance.TestData) string {
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

func (r StorageBlobContainersImmutabilityPolicyResource) basic(data acceptance.TestData) string {
    template := r.template(data)
    return fmt.Sprintf(`
%s

resource "azurerm_storage_blob_containers_immutability_policy" "test" {
  name = "acctest-sbcip-%d"
  resource_group_name = azurerm_resource_group.test.name
  account_name = azurerm_storage_account.test.name
  container_name = "container6397"
}
`, template, data.RandomInteger)
}

func (r StorageBlobContainersImmutabilityPolicyResource) requiresImport(data acceptance.TestData) string {
    config := r.basic(data)
    return fmt.Sprintf(`
%s

resource "azurerm_storage_blob_containers_immutability_policy" "import" {
  name = azurerm_storage_blob_containers_immutability_policy.test.name
  resource_group_name = azurerm_storage_blob_containers_immutability_policy.test.resource_group_name
  account_name = azurerm_storage_blob_containers_immutability_policy.test.account_name
  container_name = azurerm_storage_blob_containers_immutability_policy.test.container_name
}
`, config)
}

func (r StorageBlobContainersImmutabilityPolicyResource) complete(data acceptance.TestData) string {
    template := r.template(data)
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
