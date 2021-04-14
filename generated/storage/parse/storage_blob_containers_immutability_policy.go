package parse
import (
    azSchema "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/schema"
    uuid "github.com/gofrs/uuid"
    "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
)
type StorageBlobContainersImmutabilityPolicyId struct {
    SubscriptionId string
    ResourceGroup string
    AccountName string
    ContainerName string
    Name string
}

func NewStorageBlobContainersImmutabilityPolicyID(subscriptionId string, resourcegroup string, accountname string, containername string, name string) StorageBlobContainersImmutabilityPolicyId {
    return StorageBlobContainersImmutabilityPolicyId {
        SubscriptionId: subscriptionId,
        ResourceGroup: resourcegroup,
        AccountName: accountname,
        ContainerName: containername,
        Name: name,
    }
}

func (id StorageBlobContainersImmutabilityPolicyId) ID() string {
    fmtString := "/subscriptions/%s/resourceGroups/%s/providers/Microsoft.Storage/storageAccounts/%s/blobServices/default/containers/%s/immutabilityPolicies/%s"
    return fmt.Sprintf(fmtString, id.SubscriptionId, id.ResourceGroup, id.AccountName, id.ContainerName, id.Name)
}

func StorageBlobContainersImmutabilityPolicyID(input string) (*StorageBlobContainersImmutabilityPolicyId, error) {
    id, err := azure.ParseAzureResourceID(input)
    if err != nil {
        return nil, fmt.Errorf("parsing storageBlobContainersImmutabilityPolicy ID %q: %+v", input, err)
    }

    storageBlobContainersImmutabilityPolicy := StorageBlobContainersImmutabilityPolicyId{
        SubscriptionId: id.SubscriptionID,
        ResourceGroup: id.ResourceGroup,
    }
    if storageBlobContainersImmutabilityPolicy.AccountName, err = id.PopSegment("storageAccounts"); err != nil {
        return nil, err
    }
    if storageBlobContainersImmutabilityPolicy.ContainerName, err = id.PopSegment("containers"); err != nil {
        return nil, err
    }
    if storageBlobContainersImmutabilityPolicy.Name, err = id.PopSegment("immutabilityPolicies"); err != nil {
        return nil, err
    }
    if err := id.ValidateNoEmptySegments(input); err != nil {
        return nil, err
    }

    return &storageBlobContainersImmutabilityPolicy, nil
}
