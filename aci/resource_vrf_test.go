package aci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	cage "github.com/ignw/cisco-aci-go-sdk/src/service"
)

//TODO: Add advanced test cases to cover setting properties and Bridge Domains

func TestAccAciVrf_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciVrfDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciVrfConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciVrfExists("aci_vrf.basic"),
					resource.TestCheckResourceAttr(
						"aci_vrf.basic", "name", "vrf1"),
					resource.TestCheckResourceAttr(
						"aci_vrf.basic", "description", "terraform test VRF"),
					resource.TestCheckResourceAttr(
						"aci_vrf.basic", "domain_name", "uni/tn-IGNW-tenant1/ctx-vrf1"),
				),
			},
		},
	})
}

func testAccCheckAciVrfExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*cage.Client)

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Error -> Resource Not found: %s", n)
		}

		id := rs.Primary.Attributes["id"]

		vrf, err := client.VRFs.Get(id)

		if err != nil || vrf == nil {
			return fmt.Errorf("Error retreiving vrf id: %s", id)
		}

		return nil
	}
}

func testAccCheckAciVrfDestroy(state *terraform.State) error {

	client := testAccProvider.Meta().(*cage.Client)

	err := checkDestroy("aci_vrf.basic", state, func(s string) (interface{}, error) {
		return client.VRFs.Get(s)
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

const testAccCheckAciVrfConfigBasic = `
resource "aci_tenant" "basic" {
    name = "IGNW-tenant1"
    description = "my first tenant"
}

resource "aci_vrf" "basic" {
	name = "vrf1"
	description = "terraform test VRF"
	tenant_id = "${aci_tenant.basic.id}"	
}
`
