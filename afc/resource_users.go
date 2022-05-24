package afc

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceUsers() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceUsersCreate,
		ReadContext:   resourceUsersRead,
		UpdateContext: resourceUsersUpdate,
		DeleteContext: resourceUsersDelete,
		Schema: map[string]*schema.Schema{
			"users": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_source_uuid": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "UUID value",
						},
						"auth_source_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Authentication Source",
						},
						"distinguished_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Distinguised name (DN) of client certificate",
						},
						"role": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "User role",
						},
						"token_lifetime": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Lifetime of generated authentication token",
						},
						"username": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Username",
						},
						"uuid": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "UUID",
						},
					},
				},
			},
			"auth_source_uuid": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "UUID value",
			},
			"auth_source_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Authentication Source",
			},
			"distinguished_name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Distinguised name (DN) of client certificate",
			},
			"role": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "User role",
			},
			"token_lifetime": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Lifetime of generated authentication token",
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Username",
			},
			"uuid": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "UUID",
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceUsersCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	log.Printf("[DEBUG] %s: Beginning resourceUsersCreate", d.Id())
	var diags diag.Diagnostics

	return diags
}

func resourceUsersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	log.Printf("[DEBUG] %s: Beginning resourceUsersRead", d.Id())
	var diags diag.Diagnostics
	return diags
}

func resourceUsersUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: Beginning resourceUsersUpdate", d.Id())
	var diags diag.Diagnostics
	//return resourceOrderRead(ctx, d, m)
	return diags
}

func resourceUsersDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	log.Printf("[DEBUG] %s: Beginning resourceUsersDelete", d.Id())
	var diags diag.Diagnostics

	return diags
}
