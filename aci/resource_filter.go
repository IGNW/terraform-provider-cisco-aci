package aci

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	cage "github.com/ignw/cisco-aci-go-sdk/src/service"
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
				"tenant_id": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
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

func resourceAciFilterFieldMap() map[string]string {
	return MergeStringMaps(GetBaseFieldMap(),
		map[string]string{
			"Subjects": "subjects",
		})
}

func resourceAciFilterCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cage.Client)
	resource := &AciResource{d}

	tenant, err := ValidateAndFetchTenant(d, meta)

	if err != nil {
		return fmt.Errorf("Error creating filter: %s", err.Error())
	}

	filter := client.Filters.New(resource.Get("name").(string), resource.Get("description").(string))

	resource.MapFieldsToAci(resourceAciFilterFieldMap(), filter)

	tenant.AddFilter(filter)

	dn, err := client.Filters.Save(filter)

	if err != nil {
		return fmt.Errorf("Error saving filter: %s", err.Error())
	}

	d.Set("domain_name", dn)
	d.SetId(dn)

	return nil
}

func resourceAciFilterRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Id() == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	filter, err := client.Filters.Get(resource.Id())

	if err != nil {
		return fmt.Errorf("Error reading filter id: %s", resource.Id())
	}

	resource.MapFields(resourceAciFilterFieldMap(), filter)

	return nil
}

func resourceAciFilterUpdate(d *schema.ResourceData, meta interface{}) error {
	// HACK: currently same implementation as create
	return resourceAciFilterCreate(d, meta)
}

func resourceAciFilterDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cage.Client)
	return DeleteAciResource(d, client.Filters.Delete)
}

/*
func (d *AciResource) SetEntries(entries []*cage.Entry) {
	resources := make([]map[string]interface{}, len(entries))

	for i, entry := range entries {
		resourceEntry := make(map[string]interface{})
		resourceEntry["protocol"] = entry.Protocol
		resourceEntry["source_from"] = entry.Source.From
		resourceEntry["source_to"] = entry.Source.To
		resourceEntry["destination_from"] = entry.Destination.From
		resourceEntry["destination_to"] = entry.Destination.To

		resources[i] = resourceEntry
	}
	d.Set("entries", resources)
}
*/
