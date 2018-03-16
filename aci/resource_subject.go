package aci

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// TODO: update docs
// TODO: implement client calls
// TODO: create acceptance tests

func resourceAciSubject() *schema.Resource {
	return &schema.Resource{
		Create: resourceAciFilterCreate,
		Update: resourceAciFilterUpdate,
		Read:   resourceAciFilterRead,
		Delete: resourceAciFilterDelete,

		SchemaVersion: 1,

		Schema: GetBaseSchema(),
	}
}

func resourceAciSubjectCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAciSubjectRead(d *schema.ResourceData, m interface{}) error {

	//TODO: replace with client implementation
	d.Set("name", "http-only")
	d.Set("alias", "tf-example")
	d.Set("status", "created")
	d.Set("tags", "[]")
	return nil
}

func resourceAciSubjectUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceAciSubjectDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
