package client

import (
    azSchema "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/schema"
    uuid "github.com/gofrs/uuid"
    "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
)
type Client struct {
    BlobInventoryPolicyClient    *storage.BlobInventoryPoliciesClient
    BlobContainerClient    *storage.BlobContainersClient
    QueueClient    *storage.QueueClient
    TableClient    *storage.TableClient
}

func NewClient(o *common.ClientOptions) *Client {
    blobInventoryPolicyClient := storage.NewBlobInventoryPoliciesClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
    o.ConfigureClient(&blobInventoryPolicyClient.Client, o.ResourceManagerAuthorizer)

    blobContainerClient := storage.NewBlobContainersClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
    o.ConfigureClient(&blobContainerClient.Client, o.ResourceManagerAuthorizer)

    queueClient := storage.NewQueueClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
    o.ConfigureClient(&queueClient.Client, o.ResourceManagerAuthorizer)

    tableClient := storage.NewTableClientWithBaseURI(o.ResourceManagerEndpoint, o.SubscriptionId)
    o.ConfigureClient(&tableClient.Client, o.ResourceManagerAuthorizer)

    return &Client{
        BlobInventoryPolicyClient:    &blobInventoryPolicyClient,
        BlobContainerClient:    &blobContainerClient,
        QueueClient:    &queueClient,
        TableClient:    &tableClient,
    }
}
