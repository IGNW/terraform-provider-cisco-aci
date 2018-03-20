package aci

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/ignw/cisco-aci-go-sdk"
	"reflect"
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
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"tags": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}
}

func (d *AciResource) SetBaseFields(obj interface{}) {
	fields := []string{"name", "alias", "description", "tags"}
	original := reflect.ValueOf(obj)

	for _, key := range fields {
		d.Set(key, original.FieldByName(key).String())
	}

	d.SetId(original.FieldByName("name").Interface().(string))
}

func (d *AciResource) ConvertToBaseMap(obj *cage.ResourceAttributes) map[string]interface{} {
	fields := []string{"name", "alias", "description", "tags"}
	original := reflect.ValueOf(obj)

	mapResource := make(map[string]interface{})

	for _, key := range fields {
		mapResource[key] = original.FieldByName(key).String()
	}

	return mapResource
}

func (d *AciResource) SetIdArray(key string, items []*cage.ResourceAttributes) {
	ids := make([]string, len(items))

	for i, item := range items {
		ids[i] = item.Name
	}

	d.Set(key, ids)
}

func (d *AciResource) CreateSDKResource(obj interface{}) interface{} {
	fields := []string{"name", "alias", "description", "tags"}
	data := reflect.ValueOf(obj)

	for _, key := range fields {
		data.FieldByName(key).SetString(d.Get(key).(string))
	}
	return data
}

type AciResource struct {
	*schema.ResourceData
}
