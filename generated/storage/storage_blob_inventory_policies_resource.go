package storage

import (
    azSchema "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/schema"
    uuid "github.com/gofrs/uuid"
    "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2021-01-01/storage"
)
func resourceArmStorageBlobInventoryPolicy() *schema.Resource {
    return &schema.Resource{
        Create: resourceArmStorageBlobInventoryPolicyCreateUpdate,
        Read:   resourceArmStorageBlobInventoryPolicyRead,
        Update: resourceArmStorageBlobInventoryPolicyCreateUpdate,
        Delete: resourceArmStorageBlobInventoryPolicyDelete,

        Timeouts: &schema.ResourceTimeout{
            Create: schema.DefaultTimeout(30 * time.Minute),
            Read: schema.DefaultTimeout(5 * time.Minute),
            Update: schema.DefaultTimeout(30 * time.Minute),
            Delete: schema.DefaultTimeout(30 * time.Minute),
        },

        Importer: azSchema.ValidateResourceIDPriorToImport(func(id string) error {
            _, err := parse.StorageBlobInventoryPolicyID(id)
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

            "policy": {
                Type: schema.TypeList,
                Optional: true,
                MaxItems: 1,
                Elem: &schema.Resource{
                    Schema: map[string]*schema.Schema{
                        "destination": {
                            Type: schema.TypeString,
                            Required: true,
                        },

                        "enabled": {
                            Type: schema.TypeBool,
                            Required: true,
                        },

                        "rules": {
                            Type: schema.TypeSet,
                            Required: true,
                            Elem: &schema.Resource{
                                Schema: map[string]*schema.Schema{
                                    "name": {
                                        Type: schema.TypeString,
                                        Required: true,
                                    },

                                    "definition": {
                                        Type: schema.TypeList,
                                        Required: true,
                                        MaxItems: 1,
                                        Elem: &schema.Resource{
                                            Schema: map[string]*schema.Schema{
                                                "filters": {
                                                    Type: schema.TypeList,
                                                    Required: true,
                                                    MaxItems: 1,
                                                    Elem: &schema.Resource{
                                                        Schema: map[string]*schema.Schema{
                                                            "blob_types": {
                                                                Type: schema.TypeSet,
                                                                Required: true,
                                                                Elem: &schema.Schema{
                                                                    Type: schema.TypeString,
                                                                },
                                                            },

                                                            "include_blob_versions": {
                                                                Type: schema.TypeBool,
                                                                Optional: true,
                                                            },

                                                            "include_snapshots": {
                                                                Type: schema.TypeBool,
                                                                Optional: true,
                                                            },

                                                            "prefix_match": {
                                                                Type: schema.TypeSet,
                                                                Optional: true,
                                                                Elem: &schema.Schema{
                                                                    Type: schema.TypeString,
                                                                },
                                                            },
                                                        },
                                                    },
                                                },
                                            },
                                        },
                                    },

                                    "enabled": {
                                        Type: schema.TypeBool,
                                        Required: true,
                                    },
                                },
                            },
                        },

                        "type": {
                            Type: schema.TypeString,
                            Required: true,
                        },
                    },
                },
            },

            "last_modified_time": {
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
func resourceArmStorageBlobInventoryPolicyCreateUpdate(d *schema.ResourceData, meta interface{}) error {
    subscriptionId := meta.(*clients.Client).Account.SubscriptionId
    client := meta.(*clients.Client).Storage.BlobInventoryPolicyClient
    ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
    defer cancel()

    name := d.Get("name").(string)
    resourceGroup := d.Get("resource_group_name").(string)
    accountName := d.Get("account_name").(string)

    id := parse.NewStorageBlobInventoryPolicyID(subscriptionId, resourceGroup, accountName, name).ID()

    if d.IsNewResource() {
        existing, err :=client.Get(ctx, resourceGroup, accountName)
        if err != nil {
            if !utils.ResponseWasNotFound(existing.Response) {
                return fmt.Errorf("checking for existing Storage BlobInventoryPolicy %q (Resource Group %q / accountName %q): %+v", name, resourceGroup, accountName, err)
            }
        }
        if !utils.ResponseWasNotFound(existing.Response) {
            return tf.ImportAsExistsError("azurerm_storage_blob_inventory_policy", id)
        }
    }


    props := storage.BlobInventoryPolicy{
        BlobInventoryPolicyProperties: &storage.BlobInventoryPolicyProperties{
            Policy: expandArmBlobInventoryPolicyBlobInventoryPolicySchema(d.Get("policy").([]interface{})),
        },
    }
    if _, err :=client.CreateOrUpdate(ctx, resourceGroup, accountName, props); err != nil {
        return fmt.Errorf("creating/updating Storage BlobInventoryPolicy %q (Resource Group %q / accountName %q): %+v", name, resourceGroup, accountName, err)
    }

    d.SetId(id)
    return resourceArmStorageBlobInventoryPolicyRead(d, meta)
}

func resourceArmStorageBlobInventoryPolicyRead(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*clients.Client).Storage.BlobInventoryPolicyClient
    ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
    defer cancel()

    id, err := parse.StorageBlobInventoryPolicyID(d.Id())
    if err != nil {
        return err
    }

    resp, err :=client.Get(ctx, id.ResourceGroup, id.AccountName)
    if err != nil {
        if utils.ResponseWasNotFound(resp.Response) {
            log.Printf("[INFO] storage %q does not exist - removing from state", d.Id())
            d.SetId("")
            return nil
        }
        return fmt.Errorf("retrieving Storage BlobInventoryPolicy %q (Resource Group %q / accountName %q): %+v", id.Name, id.ResourceGroup, id.AccountName, err)
    }
    d.Set("name", id.Name)
    d.Set("resource_group_name", id.ResourceGroup)
    d.Set("account_name", id.AccountName)
    if props := resp.BlobInventoryPolicyProperties; props != nil {
        if err := d.Set("policy", flattenArmBlobInventoryPolicyBlobInventoryPolicySchema(props.Policy)); err != nil {
            return fmt.Errorf("setting `policy`: %+v", err)
        }
        d.Set("last_modified_time", props.LastModifiedTime.Format(time.RFC3339))
    }
    d.Set("type", resp.Type)
    return nil
}

func resourceArmStorageBlobInventoryPolicyDelete(d *schema.ResourceData, meta interface{}) error {
    client := meta.(*clients.Client).Storage.BlobInventoryPolicyClient
    ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
    defer cancel()

    id, err := parse.StorageBlobInventoryPolicyID(d.Id())
    if err != nil {
        return err
    }

    if _, err :=client.Delete(ctx, id.ResourceGroup, id.AccountName); err != nil {
        return fmt.Errorf("deleting Storage BlobInventoryPolicy %q (Resource Group %q / accountName %q): %+v", id.Name, id.ResourceGroup, id.AccountName, err)
    }
    return nil
}

func expandArmBlobInventoryPolicyBlobInventoryPolicySchema(input []interface{}) *storage.BlobInventoryPolicySchema {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})
    return &storage.BlobInventoryPolicySchema{
        Enabled: utils.Bool(v["enabled"].(bool)),
        Destination: utils.String(v["destination"].(string)),
        Type: utils.String(v["type"].(string)),
        Rules: expandArmBlobInventoryPolicyBlobInventoryPolicyRuleArray(v["rules"].(*schema.Set).List()),
    }
}

func expandArmBlobInventoryPolicyBlobInventoryPolicyRuleArray(input []interface{}) *[]storage.BlobInventoryPolicyRule {
    results := make([]storage.BlobInventoryPolicyRule, 0)
    for _, item := range input {
        v := item.(map[string]interface{})
        results = append(results, storage.BlobInventoryPolicyRule{
            Enabled: utils.Bool(v["enabled"].(bool)),
            Name: utils.String(v["name"].(string)),
            Definition: expandArmBlobInventoryPolicyBlobInventoryPolicyDefinition(v["definition"].([]interface{})),
        })
    }
    return &results
}

func expandArmBlobInventoryPolicyBlobInventoryPolicyDefinition(input []interface{}) *storage.BlobInventoryPolicyDefinition {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})
    return &storage.BlobInventoryPolicyDefinition{
        Filters: expandArmBlobInventoryPolicyBlobInventoryPolicyFilter(v["filters"].([]interface{})),
    }
}

func expandArmBlobInventoryPolicyBlobInventoryPolicyFilter(input []interface{}) *storage.BlobInventoryPolicyFilter {
    if len(input) == 0 {
        return nil
    }
    v := input[0].(map[string]interface{})
    return &storage.BlobInventoryPolicyFilter{
        PrefixMatch: utils.ExpandStringSlice(v["prefix_match"].(*schema.Set).List()),
        BlobTypes: utils.ExpandStringSlice(v["blob_types"].(*schema.Set).List()),
        IncludeBlobVersions: utils.Bool(v["include_blob_versions"].(bool)),
        IncludeSnapshots: utils.Bool(v["include_snapshots"].(bool)),
    }
}

func flattenArmBlobInventoryPolicyBlobInventoryPolicySchema(input *storage.BlobInventoryPolicySchema) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    var destination string
    if input.Destination != nil {
        destination = *input.Destination
    }
    var enabled bool
    if input.Enabled != nil {
        enabled = *input.Enabled
    }
    var t string
    if input.Type != nil {
        t = *input.Type
    }
    return []interface{}{
        map[string]interface{}{
            "destination": destination,
            "enabled": enabled,
            "rules": flattenArmBlobInventoryPolicyBlobInventoryPolicyRuleArray(input.Rules),
            "type": t,
        },
    }
}

func flattenArmBlobInventoryPolicyBlobInventoryPolicyRuleArray(input *[]storage.BlobInventoryPolicyRule) []interface{} {
    results := make([]interface{}, 0)
    if input == nil {
        return results
    }

    for _, item := range *input {
        var name string
        if item.Name != nil {
            name = *item.Name
        }
        var enabled bool
        if item.Enabled != nil {
            enabled = *item.Enabled
        }
        results = append(results, map[string]interface{}{
            "name": name,
            "definition": flattenArmBlobInventoryPolicyBlobInventoryPolicyDefinition(item.Definition),
            "enabled": enabled,
        })
    }
    return results
}

func flattenArmBlobInventoryPolicyBlobInventoryPolicyDefinition(input *storage.BlobInventoryPolicyDefinition) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    return []interface{}{
        map[string]interface{}{
            "filters": flattenArmBlobInventoryPolicyBlobInventoryPolicyFilter(input.Filters),
        },
    }
}

func flattenArmBlobInventoryPolicyBlobInventoryPolicyFilter(input *storage.BlobInventoryPolicyFilter) []interface{} {
    if input == nil {
        return make([]interface{}, 0)
    }

    var includeBlobVersions bool
    if input.IncludeBlobVersions != nil {
        includeBlobVersions = *input.IncludeBlobVersions
    }
    var includeSnapshots bool
    if input.IncludeSnapshots != nil {
        includeSnapshots = *input.IncludeSnapshots
    }
    return []interface{}{
        map[string]interface{}{
            "blob_types": utils.FlattenStringSlice(input.BlobTypes),
            "include_blob_versions": includeBlobVersions,
            "include_snapshots": includeSnapshots,
            "prefix_match": utils.FlattenStringSlice(input.PrefixMatch),
        },
    }
}
