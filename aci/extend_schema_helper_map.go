package aci

import "github.com/hashicorp/terraform/helper/schema"

func MergeSchemaMaps(maps ...map[string]*schema.Schema) map[string]*schema.Schema {
	size := len(maps)

	if size == 0 {
		return nil
	}

	if size == 1 {
		return maps[0]
	}
	output := make(map[string]*schema.Schema)

	for _, m := range maps {
		for k, v := range m {
			output[k] = v
		}
	}

	return output
}
