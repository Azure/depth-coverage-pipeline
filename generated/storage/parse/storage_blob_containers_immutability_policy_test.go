package parse

import (
    "testing"
    "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/resourceid"
)

var _ resourceid.Formatter = StorageBlobContainersImmutabilityPolicyId {}
func TestStorageBlobContainersImmutabilityPolicyIDFormatter(t *testing.T) {
    actual := NewStorageBlobContainersImmutabilityPolicyID("12345678-1234-9876-4563-123456789012", "resourceGroup1", "account1", "container1", "immutabilityPolicy1").ID()
    expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.Storage/storageAccounts/account1/blobServices/default/containers/container1/immutabilityPolicies/immutabilityPolicy1"
    if actual != expected {
        t.Fatalf("Expected %q but got %q", expected, actual)
    }
}

func TestStorageBlobContainersImmutabilityPolicyID(t *testing.T) {
    testData := []struct {
        Input    string
        Error    bool
        Expected *StorageBlobContainersImmutabilityPolicyId
    }{
        {
            // empty
            Input: "",
            Error: true,
        },
        {
            // missing subscriptions
            Input: "/",
            Error: true,
        },
        {
            // missing value for subscriptions
            Input: "/subscriptions/",
            Error: true,
        },
        {
            // missing resourceGroups
            Input: "/subscriptions/12345678-1234-9876-4563-123456789012/",
            Error: true,
        },
        {
            // missing value for resourceGroups
            Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/",
            Error: true,
        },
        {
            // missing storageAccounts
            Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.Storage/",
            Error: true,
        },
        {
            // missing value for storageAccounts
            Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.Storage/storageAccounts/",
            Error: true,
        },
        {
            // missing blobServices
            Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.Storage/storageAccounts/account1/",
            Error: true,
        },
        {
            // missing default
            Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.Storage/storageAccounts/account1/blobServices/",
            Error: true,
        },
        {
            // missing containers
            Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.Storage/storageAccounts/account1/blobServices/default/",
            Error: true,
        },
        {
            // missing value for containers
            Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.Storage/storageAccounts/account1/blobServices/default/containers/",
            Error: true,
        },
        {
            // missing immutabilityPolicies
            Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.Storage/storageAccounts/account1/blobServices/default/containers/container1/",
            Error: true,
        },
        {
            // missing value for immutabilityPolicies
            Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.Storage/storageAccounts/account1/blobServices/default/containers/container1/immutabilityPolicies/",
            Error: true,
        },
        {
            // valid
            Input:    "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.Storage/storageAccounts/account1/blobServices/default/containers/container1/immutabilityPolicies/immutabilityPolicy1",
            Expected: &StorageBlobContainersImmutabilityPolicyId{
                SubscriptionId:"12345678-1234-9876-4563-123456789012",
                ResourceGroup:"resourceGroup1",
                AccountName:"account1",
                ContainerName:"container1",
                Name:"immutabilityPolicy1",
            },
        },
        {
            // upper-cased
            Input:    "/SUBSCRIPTIONS/12345678-1234-9876-4563-123456789012/RESOURCEGROUPS/RESOURCEGROUP1/PROVIDERS/MICROSOFT.STORAGE/STORAGEACCOUNTS/ACCOUNT1/BLOBSERVICES/DEFAULT/CONTAINERS/CONTAINER1/IMMUTABILITYPOLICIES/IMMUTABILITYPOLICY1",
            Error: true,
        },
    }

    for _, v := range testData {
        t.Logf("[DEBUG] Testing %q", v.Input)

        actual, err := StorageBlobContainersImmutabilityPolicyID(v.Input)
        if err != nil {
            if v.Error {
                continue
            }
            t.Fatalf("Expected a value but got an error: %s", err)
        }

        if actual.SubscriptionId != v.Expected.SubscriptionId {
            t.Fatalf("Expected %q but got %q for SubscriptionId", v.Expected.SubscriptionId, actual.SubscriptionId)
        }

        if actual.ResourceGroup != v.Expected.ResourceGroup {
            t.Fatalf("Expected %q but got %q for ResourceGroup", v.Expected.ResourceGroup, actual.ResourceGroup)
        }

        if actual.AccountName != v.Expected.AccountName {
            t.Fatalf("Expected %q but got %q for AccountName", v.Expected.AccountName, actual.AccountName)
        }

        if actual.ContainerName != v.Expected.ContainerName {
            t.Fatalf("Expected %q but got %q for ContainerName", v.Expected.ContainerName, actual.ContainerName)
        }

        if actual.Name != v.Expected.Name {
            t.Fatalf("Expected %q but got %q for Name", v.Expected.Name, actual.Name)
        }
    }
}
