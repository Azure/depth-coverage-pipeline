package storage

import (
    azSchema "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/schema"
    uuid "github.com/gofrs/uuid"
    "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
)
func dataSourceStorageTable() *schema.Resource {
    return &schema.Resource{
        Read:   dataSourceArmStorageTableRead,

        Timeouts: &schema.ResourceTimeout{
            Read: schema.DefaultTimeout(5 * time.Minute),
        },

        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validate.StorageTableName,
            },

            "resource_group_name": azure.SchemaResourceGroupNameForDataSource(),

            "account_name": {
                Type: schema.TypeString,
                Required: true,
                ValidateFunc: validation.StringLenBetween(3, 24),
            },
        },
    }
}

func dataSourceArmStorageTableRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*clients.Client).Storage.TableClient
    ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
    defer cancel()

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group_name").(string)
    accountName := d.Get("account_name").(string)

    resp, err :=client.Get(ctx, resourceGroup, accountName, name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            return fmt.Errorf("Storage Table %q (Resource Group %q / accountName %q) does not exist", name, resourceGroup, accountName)
        }
        return fmt.Errorf("retrieving Storage Table %q (Resource Group %q / accountName %q): %+v", name, resourceGroup, accountName, err)
    }
    if resp.ID == nil || *resp.ID== "" {
        return fmt.Errorf("empty or nil ID returned for Storage Table %q (Resource Group %q / accountName %q) ID", name, resourceGroup, accountName)
    }

    d.SetId(*resp.ID)
    d.Set("name", name)
    d.Set("resource_group_name", resourceGroup)
    d.Set("account_name", accountName)
    return nil
}