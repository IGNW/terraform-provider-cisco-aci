package aci

import (
	"github.com/hashicorp/terraform/helper/schema"
	"reflect"
	"github.com/ignw/cisco-aci-go-sdk/src"
)

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

func (d *AciResource) SetBaseFields(obj interface{}) {
	fields := []string{"name", "alias", "status", "tags"}
	original := reflect.ValueOf(obj)

	for _, key := range fields {
		d.Set(key, original.FieldByName(key).Interface().string())
	}

	d.SetId(original.FieldByName("name").Interface().(string))
}

func (obj *cage.ResourceAttributes) ConvertBaseToMap() map[string]interface{} {
	fields := []string{"name", "alias", "status", "tags"}
	original := reflect.ValueOf(obj)

	mapResource := make(map[string]interface{})

	for _, key := range fields {
		mapResource[key] = original.FieldByName(key).Interface().string()
	}

	return mapResource
}


func (d *AciResource) CreateSDKResource(obj interface{}) interface{} {
	fields := []string{"name", "alias", "status", "tags"}
	data := reflect.ValueOf(obj)

	for _, key := range fields {
		data.FieldByName(key).SetString(d.Get(key).(string))
	}
	return data
}

type AciResource struct {
	*schema.ResourceData
}
