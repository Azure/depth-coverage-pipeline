package parse
import (
    azSchema "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/schema"
    uuid "github.com/gofrs/uuid"
    "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
)
type StorageQueueId struct {
    SubscriptionId string
    ResourceGroup string
    AccountName string
    Name string
}

func NewStorageQueueID(subscriptionId string, resourcegroup string, accountname string, name string) StorageQueueId {
    return StorageQueueId {
        SubscriptionId: subscriptionId,
        ResourceGroup: resourcegroup,
        AccountName: accountname,
        Name: name,
    }
}

func (id StorageQueueId) ID() string {
    fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s/queueServices/default/queues/%s"
    return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroup, id.AccountName, id.Name)
}

func StorageQueueID(input string) (*StorageQueueId, error) {
    id, err := azure.ParseAzureResourceID(input)
    if err != nil {
        return nil, fmt.Errorf("parsing storageQueue ID %q: %+v", input, err)
    }

    storageQueue := StorageQueueId{
        SubscriptionId: id.SubscriptionID,
        ResourceGroup: id.ResourceGroup,
    }
    if storageQueue.AccountName, err = id.PopSegment("storageAccounts"); err != nil {
        return nil, err
    }
    if storageQueue.Name, err = id.PopSegment("queues"); err != nil {
        return nil, err
    }
    if err := id.ValidateNoEmptySegments(input); err != nil {
        return nil, err
    }

    return &storageQueue, nil
}
