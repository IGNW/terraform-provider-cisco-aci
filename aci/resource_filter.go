package aci

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/ignw/cisco-aci-go-sdk"
)

// TODO: update docs
// TODO: implement client calls
// TODO: create acceptance tests

func resourceAciFilter() *schema.Resource {

	return &schema.Resource{
		Create: resourceAciFilterCreate,
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
					Elem:     &schema.Schema{Type: schema.TypeString},
				},
				"entries": &schema.Schema{
					Type:     schema.TypeList,
					Optional: true,
					Elem: &schema.Resource{
						Schema: map[string]*schema.Schema{
							"entry": &schema.Schema{
								Type:     schema.TypeList,
								Optional: true,
								Elem: &schema.Resource{
									Schema: map[string]*schema.Schema{
										"protocol": &schema.Schema{
											Type:     schema.TypeString,
											Required: true,
										},

										"source_from": &schema.Schema{
											Type:     schema.TypeString,
											Optional: true,
										},
										"source_to": &schema.Schema{
											Type:     schema.TypeString,
											Optional: true,
										},
										"destination_from": &schema.Schema{
											Type:     schema.TypeString,
											Required: true,
										},
										"destination_to": &schema.Schema{
											Type:     schema.TypeString,
											Required: true,
										},
									},
								},
							},
						},
					},
				},
			},
		),
	}
}

func resourceAciFilterCreate(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Get("name") == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	// TODO: initialize filter instance and set fields
	// filter := cage.NewFilter(resource.Get("name").(string), resource.Get("alias").(string), resource.Get("description").(string))

	/*
		response, err := client.Filters.Save(filter)
		if err != nil {
			return fmt.Errorf("Error creating filter id: %s", resource.Get("name"), err)
		}

		resource.SetBaseFields(response)
		resource.SetEntries(response.entries)
		resource.SetIdArray("subjects", response.subjects)
	*/

	return nil
}

func resourceAciFilterRead(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Id() == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	/*
		m := map[string]string{"id": resource.Id()}

		response, err := client.Filters.Get(&m)
		if err != nil {
			return fmt.Errorf("Error retrieving filter id: %s", d.Id(), err)
		}

			resource.SetBaseFields(filter)
			resource.SetEntries(filter.entries)
			resource.SetIdArray("subjects", response.subjects)
	*/

	return nil
}

func resourceAciFilterUpdate(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Id() == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	// TODO: initialize filter instance and set fields
	// filter := cage.NewFilter(resource.Get("name").(string), resource.Get("alias").(string), resource.Get("description").(string))

	/*
		response, err := client.Filters.Save(filter)
		if err != nil {
			return fmt.Errorf("Error creating filter id: %s", resource.Id(), err)
		}

		resource.SetBaseFields(response)
		resource.SetEntries(response.entries)
		resource.SetIdArray("subjects", response.subjects)
	*/

	return nil
}

func resourceAciFilterDelete(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Id() == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	// TODO: initialize filter instance and set fields

	/*
		response, err := client.Filters.Delete(resource.Id())
		if err != nil {
			return fmt.Errorf("Error creating filter id: %s", resource.Id(), err)
		}

		resource.SetBaseFields(response)
	*/

	return nil
}

func (d *AciResource) SetEntries(entries []*cage.Entry) {
	resources := make([]map[string]interface{}, len(entries))

	for i, entry := range entries {
		resourceEntry := make(map[string]interface{})
		resourceEntry["protocol"] = entry.Protocol
		resourceEntry["source_from"] = entry.Source.From
		resourceEntry["source_to"] = entry.Source.To
		resourceEntry["destination_from"] = entry.Desination.From
		resourceEntry["destination_to"] = entry.Desination.To

		resources[i] = resourceEntry
	}
	d.Set("entries", resources)
}
