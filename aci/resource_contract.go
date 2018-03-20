package aci

// TODO: update docs
// TODO: implement client calls
// TODO: create acceptance tests

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/ignw/cisco-aci-go-sdk"
)

func resourceAciContract() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciContractCreate,
		Update: resourceAciFilterUpdate,
		Read:   resourceAciFilterRead,
		Delete: resourceAciFilterDelete,

		SchemaVersion: 1,

		Schema: MergeSchemaMaps(
			GetBaseSchema(),
			map[string]*schema.Schema{
				"subjects": &schema.Schema{
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Resource{
						Schema: GetBaseSchema(),
					},
				},
				"endpoint_groups": &schema.Schema{
					Type:     schema.TypeList,
					Optional: true,
					Elem:     &schema.Schema{Type: schema.TypeString},
				},
			},
		),
	}
}

func resourceAciContractCreate(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Get("name") == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	// TODO: initialize filter instance and set fields
	// contract := cage.NewContract(resource.Get("name").(string), resource.Get("alias").(string), resource.Get("description").(string))

	/*
		response, err := client.Contracts.Save(&contract)
		if err != nil {
			return fmt.Errorf("Error creating contract id: %s", contract.name, err)
		}

		resource.SetBaseFields(response)
		resource.SetSubjects(response.subjects)
		resource.SetIdArray("endpoint_groups", response.EPGs)
	*/

	return nil
}

func resourceAciContractRead(d *schema.ResourceData, meta interface{}) error {

	// client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Get("name") == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	// TODO: initialize filter instance and set fields
	/*
		m := map[string]string{"id": resource.Id()}
		response, err := client.AppProfiles.Get(&m)

		if err != nil {
			return fmt.Errorf("Error creating contract id: %s", resource.Id(), err)
		}

			resource.SetBaseFields(response)
			resource.SetSubjects(response.subjects)
			resource.SetIdArray("endpoint_groups", response.EPGs)
	*/

	return nil
}

func resourceAciContractUpdate(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Get("name") == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	// TODO: initialize filter instance and set fields
	// contract := cage.NewContract(resource.Get("name").(string), resource.Get("alias").(string), resource.Get("description").(string))

	/*
		response, err := client.Contracts.Save(contract)
		if err != nil {
			return fmt.Errorf("Error updating contract id: %s", contract.name, err)
		}

		resource.SetBaseFields(response)
		resource.SetSubjects(response.subjects)
		resource.SetIdArray("endpoint_groups", response.EPGs)
	*/

	return nil
}

func resourceAciContractDelete(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Get("name") == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	// TODO: initialize filter instance and set fields

	/*
		response, err := client.Contracts.Delete(resource.Id())
		if err != nil {
			return fmt.Errorf("Error deleting contract id: %s", resource.Id(), err)
		}

		resource.SetBaseFields(response)
	*/

	return nil
}

func (d *AciResource) SetSubjects(subjects []*cage.Subject) {
	resources := make([]map[string]interface{}, len(subjects))

	for i, entry := range subjects {
		resources[i] = d.ConvertToBaseMap(&entry.ResourceAttributes)
	}
	d.Set("subjects", resources)
}
