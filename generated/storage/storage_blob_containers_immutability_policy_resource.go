package storage

import (
    azSchema "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/schema"
    uuid "github.com/gofrs/uuid"
    "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
)
func resourceArmStorageBlobContainersImmutabilityPolicy() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmStorageBlobContainersImmutabilityPolicyCreateUpdate,
        Read:   resourceArmStorageBlobContainersImmutabilityPolicyRead,
        Update: resourceArmStorageBlobContainersImmutabilityPolicyCreateUpdate,
        Delete: resourceArmStorageBlobContainersImmutabilityPolicyDelete,

        Timeouts: &schema.ResourceTimeout{
            Create: schema.DefaultTimeout(30 * time.Minute),
            Read: schema.DefaultTimeout(5 * time.Minute),
            Update: schema.DefaultTimeout(30 * time.Minute),
            Delete: schema.DefaultTimeout(30 * time.Minute),
        },

        Importer: azSchema.ValidateResourceIDPriorToImport(func(id string) error {
            _, err := parse.StorageBlobContainersImmutabilityPolicyID(id)
            return err
        }),


        Schema: map[string]*schema.Schema{
            "name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validation.StringInSlice([]string{
                    "default",
                }, false),
            },

            "resource_group_name": azure.SchemaResourceGroupName(),

            "account_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validation.StringLenBetween(3, 24),
            },

            "container_name": {
                Type: schema.TypeString,
                Required: true,
                ForceNew: true,
                ValidateFunc: validation.StringLenBetween(3, 63),
            },

            "allow_protected_append_writes": {
                Type: schema.TypeBool,
                Optional: true,
            },

            "immutability_period_since_creation_in_days": {
                Type: schema.TypeInt,
                Optional: true,
            },

            "state": {
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
func resourceArmStorageBlobContainersImmutabilityPolicyCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    subscriptionId := meta.(*clients.Client).Account.SubscriptionId
    client := meta.(*clients.Client).Storage.BlobContainerClient
    ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
    defer cancel()

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group_name").(string)
    accountName := d.Get("account_name").(string)
    containerName := d.Get("container_name").(string)

    id := parse.NewStorageBlobContainersImmutabilityPolicyID(subscriptionId, resourceGroup, accountName, containerName, name).ID()

    if d.IsNewResource() {
        existing, err :=client.GetImmutabilityPolicy(ctx, resourceGroup, accountName, containerName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("checking for existing Storage BlobContainer %q (Resource Group %q / accountName %q / containerName %q): %+v", name, resourceGroup, accountName, containerName, err)
            }
        }
        if !utils.ResponseWasNotFound(existing.Response) {
            return tf.ImportAsExistsError("azurerm_storage_blob_container", id)
        }
    }


    parameters := storage.ImmutabilityPolicy{
        ImmutabilityPolicyProperty: &storage.ImmutabilityPolicyProperty{
            AllowProtectedAppendWrites: utils.Bool(d.Get("allow_protected_append_writes").(bool)),
            ImmutabilityPeriodSinceCreationInDays: utils.Int32(int32(d.Get("immutability_period_since_creation_in_days").(int))),
        },
    }
    if _, err :=client.CreateOrUpdateImmutabilityPolicy(ctx, resourceGroup, accountName, containerName, &parameters); err != nil {
        return fmt.Errorf("creating/updating Storage BlobContainersImmutabilityPolicy %q (Resource Group %q / accountName %q / containerName %q): %+v", name, resourceGroup, accountName, containerName, err)
    }

    d.SetId(id)
    return resourceArmStorageBlobContainersImmutabilityPolicyRead(d, meta)
}

func resourceArmStorageBlobContainersImmutabilityPolicyRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*clients.Client).Storage.BlobContainerClient
    ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
    defer cancel()

    id, err := parse.StorageBlobContainersImmutabilityPolicyID(d.Id())
    if err != nil {
        return err
    }

    resp, err :=client.GetImmutabilityPolicy(ctx, id.ResourceGroup, id.AccountName, id.ContainerName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] storage %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("retrieving Storage BlobContainersImmutabilityPolicy %q (Resource Group %q / accountName %q / containerName %q): %+v", id.Name, id.ResourceGroup, id.AccountName, id.ContainerName, err)
    }
    d.Set("name", id.Name)
    d.Set("resource_group_name", id.ResourceGroup)
    d.Set("account_name", id.AccountName)
    d.Set("container_name", id.ContainerName)
    if props := resp.ImmutabilityPolicyProperty; props != nil {
        d.Set("allow_protected_append_writes", props.AllowProtectedAppendWrites)
        d.Set("immutability_period_since_creation_in_days", props.ImmutabilityPeriodSinceCreationInDays)
        d.Set("state", props.State)
    }
    d.Set("type", resp.Type)
    return nil
}

func resourceArmStorageBlobContainersImmutabilityPolicyDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*clients.Client).Storage.BlobContainerClient
    ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
    defer cancel()

    id, err := parse.StorageBlobContainersImmutabilityPolicyID(d.Id())
    if err != nil {
        return err
    }

    if _, err :=client.DeleteImmutabilityPolicy(ctx, id.ResourceGroup, id.AccountName, id.ContainerName); err != nil {
        return fmt.Errorf("deleting Storage BlobContainersImmutabilityPolicy %q (Resource Group %q / accountName %q / containerName %q): %+v", id.Name, id.ResourceGroup, id.AccountName, id.ContainerName, err)
    }
    return nil
}
