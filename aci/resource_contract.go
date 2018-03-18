package aci

// TODO: update docs
// TODO: implement client calls
// TODO: create acceptance tests

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/ignw/cisco-aci-go-sdk/src"
	"fmt"
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
	client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Get("name") == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	// TODO: initialize filter instance and set fields
	contract := resource.CreateSDKResource(&cage.Contract{})

	response, err := client.Contracts.New(contract)
	if err != nil {
		return fmt.Errorf("Error creating contract id: %s", contract.name, err)
	}

	resource.SetBaseFields(response)
	resource.SetEntries(response.entries)
	resource.SetSubjects(response.subjects)

	return nil
}

func resourceAciContractRead(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Get("name") == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	// TODO: initialize filter instance and set fields
	contract := resource.CreateSDKResource(&cage.Contract{})

	response, err := client.Contracts.Get(contract)
	if err != nil {
		return fmt.Errorf("Error creating contract id: %s", contract.name, err)
	}

	resource.SetBaseFields(response)
	resource.SetSubjects(response.subjects)

	return nil
}

func resourceAciContractUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Get("name") == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	// TODO: initialize filter instance and set fields
	contract := resource.CreateSDKResource(&cage.Contract{})

	response, err := client.Contracts.Update(contract)
	if err != nil {
		return fmt.Errorf("Error updating contract id: %s", contract.name, err)
	}

	resource.SetBaseFields(response)
	resource.SetSubjects(response.subjects)

	return nil
}

func resourceAciContractDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Get("name") == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	// TODO: initialize filter instance and set fields
	contract := resource.CreateSDKResource(&cage.Contract{})

	response, err := client.Contracts.Delete(contract)
	if err != nil {
		return fmt.Errorf("Error deleting contract id: %s", contract.name, err)
	}

	resource.SetBaseFields(response)

	return nil
}

func (d *AciResource) SetSubjects(subjects []*cage.Subject) {
	resources := make([]map[string]interface{}, len(subjects))

	for i, entry := range subjects {
		resources[i] = entry.ResourceAttributes.ConvertToBaseMap()
	}
	d.Set("subjects", resources)
}
