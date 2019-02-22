package aci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	cage "github.com/ignw/cisco-aci-go-sdk/src/service"
)

// TODO: Validate - this one hasn't been checked, needs tenant data in test case

//TODO: Add advanced test cases to cover setting properties and Entries

func TestAccAciFilter_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciFilterDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciFilterConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciFilterExists("aci_filter.basic"),
					resource.TestCheckResourceAttr(
						"aci_filter.basic", "name", "http-only"),
					resource.TestCheckResourceAttr(
						"aci_filter.basic", "description", "terraform test filter"),
					resource.TestCheckResourceAttr(
						"aci_filter.basic", "domain_name", "uni/tn-IGNW-tenant1/flt-http-only"),
					resource.TestCheckResourceAttr(
						"aci_filter.basic", "entries.#", "1"),
					resource.TestCheckResourceAttr(
						"aci_filter.basic", "entries.0.entry.0.destination_from", "80"),
					resource.TestCheckResourceAttr(
						"aci_filter.basic", "entries.0.entry.0.destination_to", "80"),
					resource.TestCheckResourceAttr(
						"aci_filter.basic", "entries.0.entry.0.protocol", "tcp"),
					resource.TestCheckResourceAttr(
						"aci_filter.basic", "entries.0.entry.0.source_from", "8080"),
					resource.TestCheckResourceAttr(
						"aci_filter.basic", "entries.0.entry.0.source_to", "8080"),
				),
			},
		},
	})
}

func testAccCheckAciFilterExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*cage.Client)

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Error -> Resource Not found: %s", n)
		}

		id := rs.Primary.Attributes["id"]

		filter, err := client.Filters.Get(id)

		if err != nil || filter == nil {
			return fmt.Errorf("Error retreiving filter id: %s", id)
		}

		return nil
	}
}

func testAccCheckAciFilterDestroy(state *terraform.State) error {

	client := testAccProvider.Meta().(*cage.Client)

	err := checkDestroy("aci_filter.basic", state, func(s string) (interface{}, error) {
		return client.Filters.Get(s)
	})

	if err != nil {
		return err
	}

	err = checkDestroy("aci_tenant.basic", state, func(s string) (interface{}, error) {
		return client.Tenants.Get(s)
	})

	if err != nil {
		return err
	}

	return nil
}

const testAccCheckAciFilterConfigBasic = `

resource "aci_tenant" "basic" {
    name = "IGNW-tenant1"
    description = "my first tenant"
}

resource "aci_filter" "basic" {
    name = "http-only" 
	description = "terraform test filter"	
	tenant_id = "${aci_tenant.basic.id}"

	entries {
		entry {
			protocol = "tcp"
			source_from = "8080"
			source_to = "8080"
			destination_from = "80"
			destination_to = "80"
		}
	
		entry {
				protocol = "tcp"
				destination_from = "443"
				destination_to = "443"
		}	
	}
}
`
