package aci

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	cage "github.com/ignw/cisco-aci-go-sdk/src/service"
)

func resourceAciAppProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciAppProfileCreate,
		Update: resourceAciAppProfileUpdate,
		Read:   resourceAciAppProfileRead,
		Delete: resourceAciAppProfileDelete,

		SchemaVersion: 1,

		Schema: MergeSchemaMaps(
			GetBaseSchema(),
			map[string]*schema.Schema{
				"tenant_id": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
			},
			map[string]*schema.Schema{
				"endpoint_groups": &schema.Schema{
					Type:     schema.TypeList,
					Optional: true,
					Elem:     &schema.Schema{Type: schema.TypeString},
				},
			},
		),
	}
}

func resourceAciAppProfileFieldMap() map[string]string {

	return MergeStringMaps(GetBaseFieldMap(),
		map[string]string{
			//HACK: need to fix, this wont work right now
			// Gonna need to do model.GetParent()....GetPath() or GetDomainName()
			//"Parent.id": "tenant_id",
		})
}

func resourceAciAppProfileCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cage.Client)

	tenant, err := ValidateAndFetchTenant(d, meta)

	if err != nil {
		return fmt.Errorf("Error creating app profile id: %s", d.Get("name"), err)
	}

	appProfile := client.AppProfiles.New(d.Get("name").(string), d.Get("description").(string))

	tenant.AddAppProfile(appProfile)

	dn, err := client.AppProfiles.Save(appProfile)

	if err != nil {
		return fmt.Errorf("Error creating app profile id: %s", d.Get("name"), err)
	}

	d.Set("domain_name", dn)
	d.SetId(dn)

	return nil
}

func resourceAciAppProfileRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Id() == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	appProfile, err := client.AppProfiles.Get(resource.Id())

	if err != nil {
		return fmt.Errorf("Error updating application profile id: %s", resource.Id())
	}

	resource.MapFields(resourceAciAppProfileFieldMap(), appProfile)

	return nil
}

func resourceAciAppProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	// HACK: currently same implementation as create
	return resourceAciAppProfileCreate(d, meta)
}

func resourceAciAppProfileDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cage.Client)
	return DeleteAciResource(d, client.AppProfiles.Delete)
}
