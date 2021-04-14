package storage

type Registration struct{}

// Name is the name of this Service
func (r Registration) Name() string {
    return "Storage"
}

// WebsiteCategories returns a list of categories which can be used for the sidebar
func (r Registration) WebsiteCategories() []string {
    return []string{
        "Storage",
    }
}

// SupportedDataSources returns the supported Data Sources supported by this Service
func (r Registration) SupportedDataSources() map[string]*schema.Resource {
    return map[string]*schema.Resource{
        "azurerm_storage_blob_inventory_policy":    dataSourceStorageBlobInventoryPolicy(),
        "azurerm_storage_blob_containers_immutability_policy":    dataSourceStorageBlobContainersImmutabilityPolicy(),
        "azurerm_storage_queue":    dataSourceStorageQueue(),
        "azurerm_storage_table":    dataSourceStorageTable(),
    }
}

// SupportedResources returns the supported Resources supported by this Service
func (r Registration) SupportedResources() map[string]*schema.Resource {
    return map[string]*schema.Resource{
        "azurerm_storage_blob_inventory_policy":    resourceArmStorageBlobInventoryPolicy(),
        "azurerm_storage_blob_containers_immutability_policy":    resourceArmStorageBlobContainersImmutabilityPolicy(),
        "azurerm_storage_queue":    resourceArmStorageQueue(),
        "azurerm_storage_table":    resourceArmStorageTable(),
    }
}
