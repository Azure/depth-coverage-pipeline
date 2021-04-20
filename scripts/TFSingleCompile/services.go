package provider

import (
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/sdk"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/eventhub"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/loadbalancer"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource"
)

//go:generate go run ../tools/generator-services/main.go -path=../../../

func SupportedTypedServices() []sdk.TypedServiceRegistration {
	return []sdk.TypedServiceRegistration{
		eventhub.Registration{},
		loadbalancer.Registration{},
		resource.Registration{},
	}
}

func SupportedUntypedServices() []sdk.UntypedServiceRegistration {
	return []sdk.UntypedServiceRegistration{
		resource.Registration{},
		placeholder
	}
}
