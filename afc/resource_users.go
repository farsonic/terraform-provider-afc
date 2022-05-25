package afc

import (
	"context"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	afcclient "github.com/farsonic/afc-client"
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
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceUsersCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	log.Printf("[DEBUG] %s: Beginning resourceUsersCreate", d.Id())
	var diags diag.Diagnostics
	c := m.(*ApiClient)

	name := d.Get("username").(string)
	role := d.Get("role").(string)
	token_lifetime := d.Get("token_lifetime").(string)

	a := afcclient.User{
		UserName:  name,
		Role:      role,
		TokenLife: token_lifetime,
	}

	res, err := c.afc_client.CreateUser(a)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("_id", res.ID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", res.UserName); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("role", res.Role); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("tokenlifetime", res.TokenLife); err != nil {
		return diag.FromErr(err)
	}

	d.SetId(res.ID)
	log.Printf("[DEBUG] %s: resourceUsersCreate finished successfully", d.Id())
	return diags
}

func resourceUsersRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	log.Printf("[DEBUG] %s: Beginning resourceUsersRead", d.Id())
	var diags diag.Diagnostics
	c := m.(*ApiClient)
	//res, err := c.afcclient.GetAllUsers()()()
	res, err := c.afc_client.GetAllUsers()
	if err != nil {
		return diag.FromErr(err)
	}
	if res != nil {
		resItems := flattenUsers(&res)
		if err := d.Set("avengers", resItems); err != nil {
			return diag.FromErr(err)
		}
	} else {
		return diag.Errorf("no data found in db, insert one")
	}
	log.Printf("[DEBUG] %s: resourceAvengersRead finished successfully", d.Id())
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

func flattenUsers(usersList *[]afcclient.User) []interface{} {
	if usersList != nil {
		users := make([]interface{}, len(*usersList))
		for i, user := range *usersList {
			al := make(map[string]interface{})

			al["username"] = user.UserName
			al["role"] = user.Role
			al["uuid"] = user.UUID

			users[i] = al
		}
		return users
	}
	return make([]interface{}, 0)
}
