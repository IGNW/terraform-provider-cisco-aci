package aci

import "github.com/hashicorp/terraform/helper/schema"

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

func resourceAciAppProfileCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAciAppProfileRead(d *schema.ResourceData, m interface{}) error {

	//TODO: replace with client implementation
	d.Set("name", "http-only")
	d.Set("alias", "tf-example")
	d.Set("status", "created")
	d.Set("tags", "[]")
	return nil
}

func resourceAciAppProfileUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAciAppProfileDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
