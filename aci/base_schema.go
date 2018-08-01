package aci

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	models "github.com/ignw/cisco-aci-go-sdk/src/models"
	service "github.com/ignw/cisco-aci-go-sdk/src/service"
	"reflect"
)

func GetBaseSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"domain_name": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
	}
}

func (d *AciResource) MapFields(fields map[string]string, obj interface{}) {

	original := reflect.Indirect(reflect.ValueOf(obj))

	for sourceField, destField := range fields {
		d.Set(destField, original.FieldByName(sourceField).String())
	}

	d.SetId(original.FieldByName("DomainName").Interface().(string))
}

func (d *AciResource) ConvertToBaseMap(obj *models.ResourceAttributes) map[string]interface{} {
	fields := GetBaseFieldMap()
	original := reflect.ValueOf(obj)

	mapResource := make(map[string]interface{})

	for destField, sourceField := range fields {
		mapResource[destField] = original.FieldByName(sourceField).String()
	}

	return mapResource
}

func (d *AciResource) SetIdArray(key string, items []*models.ResourceAttributes) {
	ids := make([]string, len(items))

	for i, item := range items {
		ids[i] = item.Name
	}

	d.Set(key, ids)
}

func (d *AciResource) CreateSDKResource(obj interface{}) interface{} {
	fields := GetBaseFieldMap()

	data := reflect.ValueOf(obj)

	for sourceField, destField := range fields {
		data.FieldByName(destField).SetString(d.Get(sourceField).(string))
	}
	return data
}

func GetBaseFieldMap() map[string]string {
	return map[string]string{"Name": "name", "Description": "description", "DomainName": "domain_name"}
}

func MergeStringMaps(maps ...map[string]string) map[string]string {
	size := len(maps)

	if size == 0 {
		return nil
	}

	if size == 1 {
		return maps[0]
	}
	output := make(map[string]string)

	for _, m := range maps {
		for k, v := range m {
			output[k] = v
		}
	}

	return output
}

func ValidateAndFetchTenant(d *schema.ResourceData, meta interface{}) (*models.Tenant, error) {

	client := meta.(*service.Client)

	if d.Get("name") == "" {
		return nil, fmt.Errorf("Error missing resource identifier")
	}

	if d.Get("tenant_id") == "" {
		return nil, fmt.Errorf("Error missing tenant resource identifier")
	}

	// TODO: initialize filter instance and set fields
	tenant, err := client.Tenants.Get(d.Get("tenant_id").(string))

	if err != nil {
		return nil, err
	}

	return tenant, nil
}

func DeleteAciResource(d *schema.ResourceData, del func(id string) error) error {
	if d.Id() == "" {
		return fmt.Errorf("Error missing resource identifier")
	}

	err := del(d.Id())
	if err != nil {
		return fmt.Errorf("Error deleting tenant id: %s", d.Id())
	}

	return nil
}

type AciResource struct {
	*schema.ResourceData
}

// type AciResourceDel interface func del(id string) error
