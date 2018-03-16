package aci

import (
	"github.com/hashicorp/terraform/helper/schema"
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
				"entry": &schema.Schema{
					Type:     schema.TypeSet,
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
		),
	}
}

func resourceAciFilterCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAciFilterRead(d *schema.ResourceData, m interface{}) error {

	//TODO: replace with client implementation
	d.Set("name", "http-only")
	d.Set("alias", "tf-example")
	d.Set("status", "created")
	d.Set("tags", "[]")
	return nil
}

func resourceAciFilterUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAciFilterDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
