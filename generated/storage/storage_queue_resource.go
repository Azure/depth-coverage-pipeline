package storage

import (
    azSchema "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/schema"
    uuid "github.com/gofrs/uuid"
    "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
)
func resourceArmStorageQueue() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmStorageQueueCreate,
        Read:   resourceArmStorageQueueRead,
        Update: resourceArmStorageQueueUpdate,
        Delete: resourceArmStorageQueueDelete,

        Timeouts: &schema.ResourceTimeout{
            Create: schema.DefaultTimeout(30 * time.Minute),
            Read: schema.DefaultTimeout(5 * time.Minute),
            Update: schema.DefaultTimeout(30 * time.Minute),
            Delete: schema.DefaultTimeout(30 * time.Minute),
        },

        Importer: azSchema.ValidateResourceIDPriorToImport(func(id string) error {
            _, err := parse.StorageQueueID(id)
            return err
        }),


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validate.StorageQueueName,
            },

            "resource_group_name": azure.SchemaResourceGroupName(),

            "account_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validation.StringLenBetween(3, 24),
            },

            "metadata": {
                Type: schema.TypeMap,
                Optional: true,
                Elem: &schema.Schema{
                    Type: schema.TypeString,
                },
            },

            "approximate_message_count": {
                Type: schema.TypeInt,
                Computed: true,
            },

            "type": {
                Type: schema.TypeString,
                Computed: true,
            },
        },
    }
}
func resourceArmStorageQueueCreate(d *schema.ResourceData, meta interface{}) error {
    subscriptionId := meta.(*clients.Client).Account.SubscriptionId
    client := meta.(*clients.Client).Storage.QueueClient
    ctx, cancel := timeouts.ForCreate(meta.(*clients.Client).StopContext, d)
    defer cancel()

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group_name").(string)
    accountName := d.Get("account_name").(string)

    id := parse.NewStorageQueueID(subscriptionId, resourceGroup, accountName, name).ID()

    existing, err :=client.Get(ctx, resourceGroup, accountName, name)
    if err != nil {
        if !utils.ResponseWasNotFound(existing.Response) {
            return fmt.Errorf("checking for existing Storage Queue %q (Resource Group %q / accountName %q): %+v", name, resourceGroup, accountName, err)
        }
    }
    if !utils.ResponseWasNotFound(existing.Response) {
        return tf.ImportAsExistsError("azurerm_storage_queue", id)
    }


    queue := storage.Queue{
        QueueProperties: &storage.QueueProperties{
            Metadata: utils.ExpandMapStringPtrString(d.Get("metadata").(map[string]interface{})),
        },
    }
    if _, err :=client.Create(ctx, resourceGroup, accountName, name, queue); err != nil {
        return fmt.Errorf("creating Storage Queue %q (Resource Group %q / accountName %q): %+v", name, resourceGroup, accountName, err)
    }

    d.SetId(id)
    return resourceArmStorageQueueRead(d, meta)
}

func resourceArmStorageQueueRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*clients.Client).Storage.QueueClient
    ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
    defer cancel()

    id, err := parse.StorageQueueID(d.Id())
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
        return fmt.Errorf("retrieving Storage Queue %q (Resource Group %q / accountName %q): %+v", id.Name, id.ResourceGroup, id.AccountName, err)
    }
    d.Set("name", id.Name)
    d.Set("resource_group_name", id.ResourceGroup)
    d.Set("account_name", id.AccountName)
    if queueProperties := resp.QueueProperties; queueProperties != nil {
        d.Set("metadata", utils.FlattenMapStringPtrString(queueProperties.Metadata))
        d.Set("approximate_message_count", queueProperties.ApproximateMessageCount)
    }
    d.Set("type", resp.Type)
    return nil
}

func resourceArmStorageQueueUpdate(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*clients.Client).Storage.QueueClient
    ctx, cancel := timeouts.ForUpdate(meta.(*clients.Client).StopContext, d)
    defer cancel()

    id, err := parse.StorageQueueID(d.Id())
    if err != nil {
        return err
    }

    queue := storage.Queue{
        QueueProperties: &storage.QueueProperties{
        },
    }
    if d.HasChange("name") {
        queue.Name = utils.String(d.Get("name").(string))
    }
    if d.HasChange("metadata") {
        queue.QueueProperties.Metadata = utils.ExpandMapStringPtrString(d.Get("metadata").(map[string]interface{}))
    }
    if d.HasChange("id") {
        queue.ID = utils.String(d.Get("id").(string))
    }
    if d.HasChange("type") {
        queue.Type = utils.String(d.Get("type").(string))
    }

    if _, err :=client.Update(ctx, id.ResourceGroup, id.AccountName, id.Name, queue); err != nil {
        return fmt.Errorf("updating Storage Queue %q (Resource Group %q / accountName %q): %+v", id.Name, id.ResourceGroup, id.AccountName, err)
    }
    return resourceArmStorageQueueRead(d, meta)
}

func resourceArmStorageQueueDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*clients.Client).Storage.QueueClient
    ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
    defer cancel()

    id, err := parse.StorageQueueID(d.Id())
    if err != nil {
        return err
    }

    if _, err :=client.Delete(ctx, id.ResourceGroup, id.AccountName, id.Name); err != nil {
        return fmt.Errorf("deleting Storage Queue %q (Resource Group %q / accountName %q): %+v", id.Name, id.ResourceGroup, id.AccountName, err)
    }
    return nil
}
