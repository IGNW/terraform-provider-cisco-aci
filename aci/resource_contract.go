package aci

// TODO: update docs
// TODO: implement client calls
// TODO: create acceptance tests

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	cage "github.com/ignw/cisco-aci-go-sdk/src/service"
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
				"tenant_id": &schema.Schema{
					Type:     schema.TypeString,
					Required: true,
				},
				"scope": &schema.Schema{
					Type:        schema.TypeString,
					Optional:    true,
					Description: "Represents the scope of this contract. If the scope is set as application-profile, the epg can only communicate with epgs in the same application-profile",
				},
				"dscp": &schema.Schema{
					Type:     schema.TypeString,
					Optional: true,
				},
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

func resourceAciContractFieldMap() map[string]string {
	return MergeStringMaps(GetBaseFieldMap(),
		map[string]string{
			"Scope": "scope",
			"DSCP":  "dscp",
		})
}

func resourceAciContractCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cage.Client)

	tenant, err := ValidateAndFetchTenant(d, meta)

	if err != nil {
		return fmt.Errorf("Error creating contract: %s", err.Error())
	}

	contract := client.Contracts.New(d.Get("name").(string), d.Get("description").(string))

	tenant.AddContract(contract)

	dn, err := client.Contracts.Save(contract)

	if err != nil {
		return fmt.Errorf("Error saving contract: %s", err.Error())
	}

	d.Set("domain_name", dn)
	d.SetId(dn)

	return nil
}

func resourceAciContractRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cage.Client)
	resource := &AciResource{d}

	if resource.Id() == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	contract, err := client.Contracts.Get(resource.Id())

	if err != nil {
		return fmt.Errorf("Error updating application profile id: %s", resource.Id())
	}

	resource.MapFields(resourceAciContractFieldMap(), contract)

	return nil
}

func resourceAciContractUpdate(d *schema.ResourceData, meta interface{}) error {
	// HACK: currently same implementation as create
	return resourceAciContractCreate(d, meta)
}

func resourceAciContractDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*cage.Client)
	return DeleteAciResource(d, client.Contracts.Delete)
}

/*
func (d *AciResource) SetSubjects(subjects []*cage.Subject) {
	resources := make([]map[string]interface{}, len(subjects))

	for i, entry := range subjects {
		resources[i] = d.ConvertToBaseMap(&entry.ResourceAttributes)
	}
	d.Set("subjects", resources)
}
*/
