package storage

// leaving the Storage prefix here to avoid stuttering the property name for now
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=StorageBlobInventoryPolicy -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.Storage/storageAccounts/account1/inventoryPolicies/blobInventoryPolicy1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=StorageBlobContainersImmutabilityPolicy -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.Storage/storageAccounts/account1/blobServices/default/containers/container1/immutabilityPolicies/immutabilityPolicy1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=StorageQueue -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.Storage/storageAccounts/account1/queueServices/default/queues/queue1
//go:generate go run ../../tools/generator-resource-id/main.go -path=./ -name=StorageTable -id=/subscriptions/12345678-1234-9876-4563-123456789012/resourceGroups/resourceGroup1/providers/Microsoft.Storage/storageAccounts/account1/tableServices/default/tables/table1
