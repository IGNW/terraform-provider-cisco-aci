package aci

import "github.com/hashicorp/terraform/helper/schema"

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

func resourceAciEpgCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAciEpgRead(d *schema.ResourceData, m interface{}) error {

	//TODO: replace with client implementation
	d.Set("name", "http-only")
	d.Set("alias", "tf-example")
	d.Set("status", "created")
	d.Set("tags", "[]")
	return nil
}

func resourceAciEpgUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAciEpgDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
