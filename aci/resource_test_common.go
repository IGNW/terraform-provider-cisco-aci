package aci

import (
	"fmt"
	"github.com/hashicorp/terraform/terraform"
)

func checkDestroy(n string, s *terraform.State, get func(string) (interface{}, error)) error {
	rs, ok := s.RootModule().Resources[n]

	if !ok {
		return fmt.Errorf("Error -> Resource Not found: %s", n)
	}

	id := rs.Primary.Attributes["id"]

	r, err := get(id)

	if err != nil || r == nil {
		return fmt.Errorf("Error retreiving resource (%s) id: %s", n, id)
	}

	return nil
}
