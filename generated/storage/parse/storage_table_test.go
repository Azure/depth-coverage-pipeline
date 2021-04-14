package parse

import (
    "testing"
    "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/resourceid"
)

var _ resourceid.Formatter = StorageTableId {}
func TestStorageTableIDFormatter(t *testing.T) {
    actual := NewStorageTableID("12345678-1234-9876-4563-123456789012", "resourceGroup1", "account1", "table1").ID()
    expected := "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.Storage/storageAccounts/account1/tableServices/default/tables/table1"
    if actual != expected {
        t.Fatalf("Expected %q but got %q", expected, actual)
    }
}

func TestStorageTableID(t *testing.T) {
    testData := []struct {
        Input    string
        Error    bool
        Expected *StorageTableId
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
            // missing tableServices
            Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.Storage/storageAccounts/account1/",
            Error: true,
        },
        {
            // missing default
            Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.Storage/storageAccounts/account1/tableServices/",
            Error: true,
        },
        {
            // missing tables
            Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.Storage/storageAccounts/account1/tableServices/default/",
            Error: true,
        },
        {
            // missing value for tables
            Input: "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.Storage/storageAccounts/account1/tableServices/default/tables/",
            Error: true,
        },
        {
            // valid
            Input:    "/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.Storage/storageAccounts/account1/tableServices/default/tables/table1",
            Expected: &StorageTableId{
                SubscriptionId:"12345678-1234-9876-4563-123456789012",
                ResourceGroup:"resourceGroup1",
                AccountName:"account1",
                Name:"table1",
            },
        },
        {
            // upper-cased
            Input:    "/SUBSCRIPTIONS/12345678-1234-9876-4563-123456789012/RESOURCEGROUPS/RESOURCEGROUP1/PROVIDERS/MICROSOFT.STORAGE/STORAGEACCOUNTS/ACCOUNT1/TABLESERVICES/DEFAULT/TABLES/TABLE1",
            Error: true,
        },
    }

    for _, v := range testData {
        t.Logf("[DEBUG] Testing %q", v.Input)

        actual, err := StorageTableID(v.Input)
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

        if actual.Name != v.Expected.Name {
            t.Fatalf("Expected %q but got %q for Name", v.Expected.Name, actual.Name)
        }
    }
}
