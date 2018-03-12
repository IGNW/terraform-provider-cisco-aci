package aci

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceCiscoAciFilter() *schema.Resource {
	return &schema.Resource{
		Create: resourceCiscoAciFilterCreate,
		Update: resourceCiscoAciFilterUpdate,
		Read:   resourceCiscoAciFilterRead,
		Delete: resourceCiscoAciFilterDelete,

		SchemaVersion: 1,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
			},
			"alias": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},

			"entry": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},

						"source-from": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"source-to": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"destination-from": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"destination-to": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
		},
	}
}

func resourceCiscoAciFilterCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCiscoAciFilterRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCiscoAciFilterUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resourceCiscoAciFilterDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}