package aci

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	cage "github.com/ignw/cisco-aci-go-sdk/src/service"
)

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

func resourceAciTenantCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cage.Client)

	if d.Get("name") == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	// TODO: initialize filter instance and set fields
	tenant := client.Tenants.New(d.Get("name").(string), d.Get("description").(string))

	dn, err := client.Tenants.Save(tenant)
	if err != nil {
		return fmt.Errorf("Error creating tenant id: %s", d.Get("name"))
	}

	d.Set("domain_name", dn)
	d.SetId(dn)

	return nil
}

func resourceAciTenantRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Id() == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	// TODO: initialize filter instance and set fields
	tenant, err := client.Tenants.Get(resource.Id())

	if err != nil {
		return fmt.Errorf("Error updating tenant id: %s", resource.Id())
	}

	// TODO: map collections
	resource.MapFields(GetBaseFieldMap(), tenant)

	return nil
}

func resourceAciTenantUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cage.Client)

	if d.Id() == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	// TODO: initialize filter instance and set fields
	tenant := client.Tenants.New(d.Get("name").(string), d.Get("description").(string))

	_, err := client.Tenants.Save(tenant)

	if err != nil {
		return fmt.Errorf("Error updating tenant id: %s", d.Id())
	}

	return nil
}

func resourceAciTenantDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cage.Client)
	return DeleteAciResource(d, client.Tenants.Delete)
}
