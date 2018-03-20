package aci

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
)

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
				"bridge_domains": &schema.Schema{
					Type:     schema.TypeList,
					Optional: true,
					Elem:     &schema.Schema{Type: schema.TypeString},
				},
			},
		),
	}
}

func resourceAciVrfCreate(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Get("name") == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	// TODO: initialize filter instance and set fields
	// vrf := cage.NewVRF(resource.Get("name").(string), resource.Get("alias").(string), resource.Get("description").(string))

	/*
		response, err := client.VRFs.Save(vrf)
		if err != nil {
			return fmt.Errorf("Error creating VRF id: %s", vrf.name, err)
		}

		resource.SetBaseFields(response)
		resource.SetIdArray("bridge_domains", response.BridgeDomains)
	*/

	return nil
}

func resourceAciVrfRead(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Get("name") == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	// TODO: initialize filter instance and set fields
	/*
		m := map[string]string{"id": resource.Id()}

			response, err := client.VRFs.Get(&m)
			if err != nil {
				return fmt.Errorf("Error creating VRF id: %s", vrf.name, err)
			}

			resource.SetBaseFields(response)
			resource.SetIdArray("bridge_domains", response.BridgeDomains)
	*/

	return nil
}

func resourceAciVrfUpdate(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Id() == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	// TODO: initialize filter instance and set fields
	// vrf := cage.NewVRF(resource.Get("name").(string), resource.Get("alias").(string), resource.Get("description").(string))

	/*
		response, err := client.VRFs.Save(vrf)
		if err != nil {
			return fmt.Errorf("Error updating VRF id: %s", resource.Id(), err)
		}

		resource.SetBaseFields(response)
		resource.SetIdArray("bridge_domains", response.BridgeDomains)
	*/

	return nil
}

func resourceAciVrfDelete(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Id() == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	// TODO: initialize filter instance and set fields
	/*
		response, err := client.VRFs.Delete(resource.Id())
		if err != nil {
			return fmt.Errorf("Error deleting VRF id: %s", resource.Id(), err)
		}

		resource.SetBaseFields(response)
	*/

	return nil
}
