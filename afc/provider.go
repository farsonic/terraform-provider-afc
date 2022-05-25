package afc

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	afcclient "github.com/farsonic/afc-client"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"users_resource": resourceUsers(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"users_datasource": dataSourceUsers(),
		},
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("AFC_URL", "http://localhost:8000"),
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("AFC_USERNAME", nil),
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("AFC_PASSWORD", nil),
			},
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("AFC_TOKEN", nil),
			},
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return NewApiClient(d)

}

type ApiClient struct {
	data       *schema.ResourceData
	afc_client *afcclient.Client
}

//NewApiClient will return a new instance of ApiClient using which we can communicate with AFC
func NewApiClient(d *schema.ResourceData) (*ApiClient, diag.Diagnostics) {
	c := &ApiClient{data: d}
	client, err := c.NewAFCClient()
	if err != nil {
		return c, diag.FromErr(err)
	}
	c.afc_client = client
	return c, nil

}

func (a *ApiClient) NewAFCClient() (*afcclient.Client, error) {
	host := a.data.Get("host").(string)
	c, err := afcclient.NewClient(&host)
	if err != nil {
		return c, err
	}
	return c, nil
}
