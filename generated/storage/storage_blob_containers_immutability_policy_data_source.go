package storage

import (
    azSchema "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/schema"
    uuid "github.com/gofrs/uuid"
    "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
)
func dataSourceStorageBlobContainersImmutabilityPolicy() *schema.Resource {
    return &schema.Resource{
        Read:   dataSourceArmStorageBlobContainersImmutabilityPolicyRead,

        Timeouts: &schema.ResourceTimeout{
            Read: schema.DefaultTimeout(5 * time.Minute),
        },

        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validation.StringInSlice([]string{
                    "default",
                }, false),
            },

            "resource_group_name": azure.SchemaResourceGroupNameForDataSource(),

            "account_name": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validation.StringLenBetween(3, 24),
            },

            "container_name": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validation.StringLenBetween(3, 63),
            },
        },
    }
}

func dataSourceArmStorageBlobContainersImmutabilityPolicyRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*clients.Client).Storage.BlobContainerClient
    ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
    defer cancel()

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group_name").(string)
    accountName := d.Get("account_name").(string)
    containerName := d.Get("container_name").(string)

    resp, err :=client.GetImmutabilityPolicy(ctx, resourceGroup, accountName, containerName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            return fmt.Errorf("Storage BlobContainersImmutabilityPolicy %q (Resource Group %q / accountName %q / containerName %q) does not exist", name, resourceGroup, accountName, containerName)
        }
        return fmt.Errorf("retrieving Storage BlobContainersImmutabilityPolicy %q (Resource Group %q / accountName %q / containerName %q): %+v", name, resourceGroup, accountName, containerName, err)
    }
    if resp.ID == nil || *resp.ID== "" {
        return fmt.Errorf("empty or nil ID returned for Storage BlobContainer %q (Resource Group %q / accountName %q / containerName %q) ID", name, resourceGroup, accountName, containerName)
    }

    d.SetId(*resp.ID)
    d.Set("name", name)
    d.Set("resource_group_name", resourceGroup)
    d.Set("account_name", accountName)
    d.Set("container_name", containerName)
    return nil
}
