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
				Required: true,
			},
			"alias": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"tags": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
