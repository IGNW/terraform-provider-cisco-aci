package aci

import "github.com/hashicorp/terraform/helper/schema"

func GetBaseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
	}
}
