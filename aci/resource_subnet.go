
package aci

import "github.com/hashicorp/terraform/helper/schema"

func resourceAciSubnet() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciSubnetCreate,
		Update: resourceAciSubnetUpdate,
		Read:   resourceAciSubnetRead,
		Delete: resourceAciSubnetDelete,

		SchemaVersion: 1,

		Schema: GetBaseSchema(),
	}
}

func resourceAciSubnetCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAciSubnetRead(d *schema.ResourceData, m interface{}) error {

	//TODO: replace with client implementation
	d.Set("name", "http-only")
	d.Set("alias", "tf-example")
	d.Set("status", "created")
	d.Set("tags", "[]")
	return nil
}

func resourceAciSubnetUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAciSubnetDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}