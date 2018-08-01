package aci

import (
	"fmt"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	cage "github.com/ignw/cisco-aci-go-sdk/src/service"
	"testing"
)

func TestAccAciBridgeDomain_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciBridgeDomainDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciBridgeDomainConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciBridgeDomainExists("aci_app_profile.basic_app"),
					resource.TestCheckResourceAttr(
						"aci_app_profile.basic_app", "name", "ap1"),
					resource.TestCheckResourceAttr(
						"aci_app_profile.basic_app", "description", "terraform test app profile"),
					resource.TestCheckResourceAttr(
						"aci_app_profile.basic_app", "domain_name", "uni/tn-tenant1/ap-ap1"),
				),
			},
		},
	})
}

func testAccCheckAciBridgeDomainExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*cage.Client)

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Error -> Resource Not found: %s", n)
		}

		id := rs.Primary.Attributes["id"]

		ap, err := client.AppProfiles.Get(id)

		if err != nil || ap == nil {
			return fmt.Errorf("Error retreiving app profile id: %s", id)
		}

		return nil
	}
}

func testAccCheckAciBridgeDomainDestroy(state *terraform.State) error {

	client := testAccProvider.Meta().(*cage.Client)

	err := checkDestroy("aci_app_profile.basic_app", state, func(s string) (interface{}, error) {
		return client.AppProfiles.Get(s)
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

const testAccCheckAciBridgeDomainConfigBasic = `
resource "aci_tenant" "basic_bd" {
    name = "tenant1"
    description = "my first tenant"  
}

resource "aci_bridge_domain" "basic_bd" {
	name = "bd1"
	description = "terraform test bridge domain"
	tenant_id = "${aci_tenant.basic_bd.id}"
}
`
