package aci

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/ignw/cisco-aci-go-sdk"
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
				"endpoint_groups": &schema.Schema{
					Type:     schema.TypeList,
					Optional: true,
					Elem:     &schema.Schema{Type: schema.TypeString},
				},
			},
		),
	}
}

func resourceAciAppProfileCreate(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Get("name") == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	// TODO: initialize filter instance and set fields
	// appProfile := cage.NewAppProfile(resource.Get("name").(string), resource.Get("alias").(string), resource.Get("description").(string))

	/*
		response, err := client.AppProfiles.Save(appProfile)
		if err != nil {
			return fmt.Errorf("Error creating app profile id: %s", resource.Get("name"), err)
		}
	*/

	// resource.SetBaseFields(response)
	// resource.SetIdArray("endpoint_groups", response.EPGs.([]cage.ResourceAttributes))

	return nil
}

func resourceAciAppProfileRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Get("name") == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	// TODO: initialize filter instance and set fields
	m := map[string]string{"id": resource.Id()}

	response, err := client.AppProfiles.Get(&m)
	if err != nil {
		return fmt.Errorf("Error creating app profile id: %s", resource.Id(), err)
	}

	resource.SetBaseFields(response)
	// resource.SetIdArray("endpoint_groups", response.EPGs)

	return nil
}

func resourceAciAppProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Id() == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	// TODO: initialize filter instance and set fields
	// appProfile := cage.NewAppProfile(resource.Get("name").(string), resource.Get("alias").(string), resource.Get("description").(string))

	/*
		response, err := client.AppProfiles.Save(appProfile)
		if err != nil {
			return fmt.Errorf("Error creating app profile id: %s", resource.Id(), err)
		}

		resource.SetBaseFields(response)
		resource.SetIdArray("endpoint_groups", response.EPGs)
	*/

	return nil
}

func resourceAciAppProfileDelete(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Id() == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	/*
		response, err := client.AppProfiles.Delete(resource.Id())
		if err != nil {
			return fmt.Errorf("Error creating app profile id: %s", resource.Id(), err)
		}
	*/

	return nil
}
