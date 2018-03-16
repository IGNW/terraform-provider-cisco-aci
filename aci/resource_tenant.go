
package aci

import "github.com/hashicorp/terraform/helper/schema"

func resourceAciTenant() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciTenantCreate,
		Update: resourceAciTenantUpdate,
		Read:   resourceAciTenantRead,
		Delete: resourceAciTenantDelete,

		SchemaVersion: 1,

		Schema: MergeSchemaMaps(
			GetBaseSchema(),
			map[string]*schema.Schema{
				"l3_net_identifier": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"vrfs": &schema.Schema{
					Type:     schema.TypeList,
					Optional: true,
					Elem:     &schema.Schema{Type: schema.TypeString},
				},
				"endpoint_groups": &schema.Schema{
					Type:     schema.TypeList,
					Optional: true,
					Elem:     &schema.Schema{Type: schema.TypeString},
				},
				"app_profiles": &schema.Schema{
					Type:     schema.TypeList,
					Optional: true,
					Elem:     &schema.Schema{Type: schema.TypeString},
				},
				"contracts": &schema.Schema{
					Type:     schema.TypeList,
					Optional: true,
					Elem:     &schema.Schema{Type: schema.TypeString},
				},
				"filters": &schema.Schema{
					Type:     schema.TypeList,
					Optional: true,
					Elem:     &schema.Schema{Type: schema.TypeString},
				},
			},
		),
	}
}

func resourceAciTenantCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAciTenantRead(d *schema.ResourceData, m interface{}) error {

	//TODO: replace with client implementation
	d.Set("name", "http-only")
	d.Set("alias", "tf-example")
	d.Set("status", "created")
	d.Set("tags", "[]")
	return nil
}

func resourceAciTenantUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAciTenantDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}