package aci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	cage "github.com/ignw/cisco-aci-go-sdk/src/service"
)

func TestAccAciAppProfile_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciAppProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciAppProfileConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciAppProfileExists("aci_app_profile.basic_app"),
					resource.TestCheckResourceAttr(
						"aci_app_profile.basic_app", "name", "IGNW-ap1"),
					resource.TestCheckResourceAttr(
						"aci_app_profile.basic_app", "description", "terraform test app profile"),
					resource.TestCheckResourceAttr(
						"aci_app_profile.basic_app", "domain_name", "uni/tn-IGNW-tenant1/ap-IGNW-ap1"),
				),
			},
		},
	})
}

func testAccCheckAciAppProfileExists(n string) resource.TestCheckFunc {
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

func testAccCheckAciAppProfileDestroy(state *terraform.State) error {

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

const testAccCheckAciAppProfileConfigBasic = `
resource "aci_tenant" "basic" {
    name = "IGNW-tenant1"
    description = "my first tenant"
}

resource "aci_app_profile" "basic_app" {
	name = "IGNW-ap1"
	description = "terraform test app profile"
	tenant_id = "${aci_tenant.basic.id}"
}
`
