package aci

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	cage "github.com/ignw/cisco-aci-go-sdk/src/service"
)

//TODO: Add advanced test cases to cover setting properties and bridge domains

func resourceAciVrf() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciVrfCreate,
		Update: resourceAciVrfUpdate,
		Read:   resourceAciVrfRead,
		Delete: resourceAciVrfDelete,

		SchemaVersion: 1,

		Schema: MergeSchemaMaps(
			GetBaseSchema(),
			map[string]*schema.Schema{
				"tenant_id": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"enforce": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					Default:  "unenforced",
				},
				"enforcement_direction": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
					Default:  "ingress",
				},
				"bridge_domains": &schema.Schema{
					Type:     schema.TypeList,
					Optional: true,
					Elem:     &schema.Schema{Type: schema.TypeString},
				},
			},
		),
	}
}

func resourceAciVrfFieldMap() map[string]string {
	return MergeStringMaps(GetBaseFieldMap(),
		map[string]string{
			"Enforce":              "enforce",
			"EnforcementDirection": "enforcement_direction",
		})
}

func resourceAciVrfCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cage.Client)
	resource := &AciResource{d}

	tenant, err := ValidateAndFetchTenant(d, meta)

	if err != nil {
		return fmt.Errorf("Error creating VRF: %s", err.Error())
	}

	vrf := client.VRFs.New(resource.Get("name").(string), resource.Get("description").(string))

	resource.MapFieldsToAci(resourceAciVrfFieldMap(), vrf)

	tenant.AddVRF(vrf)

	dn, err := client.VRFs.Save(vrf)

	if err != nil {
		return fmt.Errorf("Error saving VRF: %s", err.Error())
	}

	d.Set("domain_name", dn)
	d.SetId(dn)

	return nil
}

func resourceAciVrfRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Id() == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	vrf, err := client.VRFs.Get(resource.Id())

	if err != nil {
		return fmt.Errorf("Error reading vrf id: %s", resource.Id())
	}

	resource.MapFields(resourceAciVrfFieldMap(), vrf)

	return nil
}

func resourceAciVrfUpdate(d *schema.ResourceData, meta interface{}) error {
	// HACK: currently same implementation as create
	return resourceAciVrfCreate(d, meta)
}

func resourceAciVrfDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cage.Client)
	return DeleteAciResource(d, client.VRFs.Delete)
}
