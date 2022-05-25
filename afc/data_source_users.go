package afc

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var HostURL string = "http://localhost:8000"

//dataSourceUsers is the Users data source which will pull information on all Users
func dataSourceUsers() *schema.Resource {
	return &schema.Resource{
		ReadContext: resourceUsersRead,
		Schema: map[string]*schema.Schema{
			"users": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_source_uuid": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "UUID value",
						},
						"auth_source_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Authentication Source",
						},
						"distinguished_name": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Distinguised name (DN) of client certificate",
						},
						"role": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "User role",
						},
						"token_lifetime": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "Lifetime of generated authentication token",
						},
						"username": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Username",
						},
						"uuid": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "UUID",
						},
					},
				},
			},
		},
	}
}
