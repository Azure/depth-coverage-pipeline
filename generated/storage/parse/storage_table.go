package parse
import (
    azSchema "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/schema"
    uuid "github.com/gofrs/uuid"
    "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
)
type StorageTableId struct {
    SubscriptionId string
    ResourceGroup string
    AccountName string
    Name string
}

func NewStorageTableID(subscriptionId string, resourcegroup string, accountname string, name string) StorageTableId {
    return StorageTableId {
        SubscriptionId: subscriptionId,
        ResourceGroup: resourcegroup,
        AccountName: accountname,
        Name: name,
    }
}

func (id StorageTableId) ID() string {
    fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s/tableServices/default/tables/%s"
    return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroup, id.AccountName, id.Name)
}

func StorageTableID(input string) (*StorageTableId, error) {
    id, err := azure.ParseAzureResourceID(input)
    if err != nil {
        return nil, fmt.Errorf("parsing storageTable ID %q: %+v", input, err)
    }

    storageTable := StorageTableId{
        SubscriptionId: id.SubscriptionID,
        ResourceGroup: id.ResourceGroup,
    }
    if storageTable.AccountName, err = id.PopSegment("storageAccounts"); err != nil {
        return nil, err
    }
    if storageTable.Name, err = id.PopSegment("tables"); err != nil {
        return nil, err
    }
    if err := id.ValidateNoEmptySegments(input); err != nil {
        return nil, err
    }

    return &storageTable, nil
}
