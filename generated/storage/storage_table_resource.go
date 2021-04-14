package storage

import (
    azSchema "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/schema"
    uuid "github.com/gofrs/uuid"
    "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
)
func resourceArmStorageTable() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmStorageTableCreate,
        Read:   resourceArmStorageTableRead,
        Delete: resourceArmStorageTableDelete,

        Timeouts: &schema.ResourceTimeout{
            Create: schema.DefaultTimeout(30 * time.Minute),
            Read: schema.DefaultTimeout(5 * time.Minute),
            Delete: schema.DefaultTimeout(30 * time.Minute),
        },

        Importer: azSchema.ValidateResourceIDPriorToImport(func(id string) error {
            _, err := parse.StorageTableID(id)
            return err
        }),


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.StorageTableName,
            },

            "resource_group_name": azure.SchemaResourceGroupName(),

            "account_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validation.StringLenBetween(3, 24),
            },

            "table_name": {
                Type: schema.TypeString,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}
func resourceArmStorageTableCreate(d *schema.ResourceData, meta interface{}) error {
    subscriptionId := meta.(*clients.Client).Account.SubscriptionId
    client := meta.(*clients.Client).Storage.TableClient
    ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
    defer cancel()

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group_name").(string)
    accountName := d.Get("account_name").(string)

    id := parse.NewStorageTableID(subscriptionId, resourceGroup, accountName, name).ID()

    existing, err :=client.Get(ctx, resourceGroup, accountName, name)
    if err != nil {
        if !utils.ResponseWasNotFound(existing.Response) {
            return fmt.Errorf("checking for existing Storage Table %q (Resource Group %q / accountName %q): %+v", name, resourceGroup, accountName, err)
        }
    }
    if !utils.ResponseWasNotFound(existing.Response) {
        return tf.ImportAsExistsError("azurerm_storage_table", id)
    }

    if _, err :=client.Create(ctx, resourceGroup, accountName, name); err != nil {
        return fmt.Errorf("creating Storage Table %q (Resource Group %q / accountName %q): %+v", name, resourceGroup, accountName, err)
    }

    d.SetId(id)
    return resourceArmStorageTableRead(d, meta)
}

func resourceArmStorageTableRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*clients.Client).Storage.TableClient
    ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
    defer cancel()

    id, err := parse.StorageTableID(d.Id())
    if err != nil {
        return err
    }

    resp, err :=client.Get(ctx, id.ResourceGroup, id.AccountName, id.Name)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] storage %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("retrieving Storage Table %q (Resource Group %q / accountName %q): %+v", id.Name, id.ResourceGroup, id.AccountName, err)
    }
    d.Set("name", id.Name)
    d.Set("resource_group_name", id.ResourceGroup)
    d.Set("account_name", id.AccountName)
    if tableProperties := resp.TableProperties; tableProperties != nil {
        d.Set("table_name", tableProperties.TableName)
    }
    d.Set("type", resp.Type)
    return nil
}


func resourceArmStorageTableDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*clients.Client).Storage.TableClient
    ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
    defer cancel()

    id, err := parse.StorageTableID(d.Id())
    if err != nil {
        return err
    }

    if _, err :=client.Delete(ctx, id.ResourceGroup, id.AccountName, id.Name); err != nil {
        return fmt.Errorf("deleting Storage Table %q (Resource Group %q / accountName %q): %+v", id.Name, id.ResourceGroup, id.AccountName, err)
    }
    return nil
}
