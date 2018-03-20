package aci

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/ignw/cisco-aci-go-sdk"
)

func resourceAciBridgeDomain() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciBridgeDomainCreate,
		Update: resourceAciBridgeDomainUpdate,
		Read:   resourceAciBridgeDomainRead,
		Delete: resourceAciBridgeDomainDelete,

		SchemaVersion: 1,

		Schema: MergeSchemaMaps(
			GetBaseSchema(),
			map[string]*schema.Schema{
				"subnets": &schema.Schema{
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

func resourceAciBridgeDomainCreate(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Get("name") == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	// TODO: initialize filter instance and set fields
	// domain := cage.NewBridgeDomain(resource.Get("name").(string), resource.Get("alias").(string), resource.Get("description").(string))

	/*
		response, err := client.BridgeDomains.New(domain)
		if err != nil {
			return fmt.Errorf("Error creating bridge domain id: %s", domain.name, err)
		}

		resource.SetBaseFields(response)
		resource.SetIdArray("endpoint_groups", response.EPGs)
		resource.SetSubnets(response.subnets)
	*/

	return nil
}

func resourceAciBridgeDomainRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Get("name") == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	// TODO: initialize filter instance and set fields
	m := map[string]string{"id": resource.Id()}

	response, err := client.BridgeDomains.Get(&m)
	if err != nil {
		return fmt.Errorf("Error creating bridge domain id: %s", resource.Get("name"), err)
	}

	resource.SetBaseFields(response)
	// resource.SetIdArray("endpoint_groups", response.EPGs)
	// resource.SetSubnets(response.subnets)

	return nil
}

func resourceAciBridgeDomainUpdate(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Id() == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	// TODO: initialize filter instance and set fields
	// domain := cage.NewBridgeDomain(resource.Get("name").(string), resource.Get("alias").(string), resource.Get("description").(string))

	/*
		response, err := client.BridgeDomains.Save(domain)
		if err != nil {
			return fmt.Errorf("Error updating bridge domain id: %s", resource.Id(), err)
		}

		resource.SetBaseFields(response)
		resource.SetIdArray("endpoint_groups", response.EPGs)
		resource.SetSubnets(response.subnets)
	*/

	return nil
}

func resourceAciBridgeDomainDelete(d *schema.ResourceData, meta interface{}) error {
	// client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Id() == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	/*
		response, err := client.BridgeDomains.Delete(resource.Id())
		if err != nil {
			return fmt.Errorf("Error deleting bridge domain id: %s", resource.Id(), err)
		}

		resource.SetBaseFields(response)
		resource.SetIdArray("endpoint_groups", response.EPGs)
		resource.SetSubnets(response.subnets)
	*/

	return nil
}

func (d *AciResource) SetSubnets(items []*cage.Subnet) {
	subnets := make([]map[string]interface{}, len(items))

	for i, item := range items {
		subnets[i] = d.ConvertToBaseMap(&item.ResourceAttributes)
	}
	d.Set("subnets", subnets)
}
