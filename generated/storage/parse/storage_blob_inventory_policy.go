package parse
import (
    azSchema "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/schema"
    uuid "github.com/gofrs/uuid"
    "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
)
type StorageBlobInventoryPolicyId struct {
    SubscriptionId string
    ResourceGroup string
    AccountName string
    Name string
}

func NewStorageBlobInventoryPolicyID(subscriptionId string, resourcegroup string, accountname string, name string) StorageBlobInventoryPolicyId {
    return StorageBlobInventoryPolicyId {
        SubscriptionId: subscriptionId,
        ResourceGroup: resourcegroup,
        AccountName: accountname,
        Name: name,
    }
}

func (id StorageBlobInventoryPolicyId) ID() string {
    fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s/inventoryPolicies/%s"
    return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroup, id.AccountName, id.Name)
}

func StorageBlobInventoryPolicyID(input string) (*StorageBlobInventoryPolicyId, error) {
    id, err := azure.ParseAzureResourceID(input)
    if err != nil {
        return nil, fmt.Errorf("parsing storageBlobInventoryPolicy ID %q: %+v", input, err)
    }

    storageBlobInventoryPolicy := StorageBlobInventoryPolicyId{
        SubscriptionId: id.SubscriptionID,
        ResourceGroup: id.ResourceGroup,
    }
    if storageBlobInventoryPolicy.AccountName, err = id.PopSegment("storageAccounts"); err != nil {
        return nil, err
    }
    if storageBlobInventoryPolicy.Name, err = id.PopSegment("inventoryPolicies"); err != nil {
        return nil, err
    }
    if err := id.ValidateNoEmptySegments(input); err != nil {
        return nil, err
    }

    return &storageBlobInventoryPolicy, nil
}
