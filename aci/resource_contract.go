package aci

// TODO: update docs
// TODO: implement client calls
// TODO: create acceptance tests

import (
	"github.com/hashicorp/terraform/helper/schema"
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
					Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceAciContractCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAciContractRead(d *schema.ResourceData, m interface{}) error {

	//TODO: replace with client implementation
	d.Set("name", "http-only")
	d.Set("alias", "tf-example")
	d.Set("status", "created")
	d.Set("tags", "[]")
	return nil
}

func resourceAciContractUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAciContractDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
