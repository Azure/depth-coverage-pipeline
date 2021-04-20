package clients

import (
	"context"

	eventhub "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/eventhub/client"

	loadbalancers "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/loadbalancer/client"

	resource "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/resource/client"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/common"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/features"
)

type Client struct {
	// StopContext is used for propagating control from Terraform Core (e.g. Ctrl/Cmd+C)
	StopContext context.Context

	Account  *ResourceManagerAccount
	Features features.UserFeatures


	Eventhub              *eventhub.Client

	LoadBalancers         *loadbalancers.Client

	Resource              *resource.Client

}

// NOTE: it should be possible for this method to become Private once the top level Client's removed

func (client *Client) Build(ctx context.Context, o *common.ClientOptions) error {
	autorest.Count429AsRetry = false
	// Disable the Azure SDK for Go's validation since it's unhelpful for our use-case
	validation.Disabled = true

	client.Features = o.Features
	client.StopContext = ctx

	client.Eventhub = eventhub.NewClient(o)

	client.LoadBalancers = loadbalancers.NewClient(o)

	client.Resource = resource.NewClient(o)


	return nil
}
