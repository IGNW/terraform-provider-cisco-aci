package aci

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAciEpg() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciEpgCreate,
		Update: resourceAciEpgUpdate,
		Read:   resourceAciEpgRead,
		Delete: resourceAciEpgDelete,

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

func resourceAciEpgCreate(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Get("name") == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	// TODO: initialize filter instance and set fields
	// epg := cage.NewEPG(resource.Get("name").(string), resource.Get("alias").(string), resource.Get("description").(string))

	/*
		response, err := client.EPGs.Save(epg)
		if err != nil {
			return fmt.Errorf("Error creating endpoint group id: %s", epg.name, err)
		}

		resource.SetBaseFields(response)
		resource.SetIdArray("bridge_domains", response.BrideDomains)
	*/

	return nil
}

func resourceAciEpgRead(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Get("name") == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	// TODO: initialize filter instance and set fields
	/*
		m := map[string]string{"id": resource.Id()}

		response, err := client.EPGs.Get(&m)
		if err != nil {
			return fmt.Errorf("Error retreiving endpoint group id: %s", resource.Get("name"), err)
		}


			resource.SetBaseFields(response)
			resource.SetIdArray("bridge_domains", response.BrideDomains)
	*/

	return nil
}

func resourceAciEpgUpdate(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Id() == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	// TODO: initialize filter instance and set fields
	// epg := cage.NewEPG(resource.Get("name").(string), resource.Get("alias").(string), resource.Get("description").(string))

	/*
		response, err := client.EPGs.Save(epg)
		if err != nil {
			return fmt.Errorf("Error updating endpoint group id: %s", resource.Id(), err)
		}

		resource.SetBaseFields(response)
		resource.SetIdArray("bridge_domains", response.BrideDomains)
	*/

	return nil
}

func resourceAciEpgDelete(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Id() == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	// TODO: initialize filter instance and set fields
	/*
		response, err := client.EPGs.Delete(resource.Id())
		if err != nil {
			return fmt.Errorf("Error deleting endpoint group id: %s", resource.Id(), err)
		}

			resource.SetBaseFields(response)
			resource.SetIdArray("bridge_domains", response.BrideDomains)
	*/

	return nil
}
