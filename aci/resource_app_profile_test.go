package aci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	cage "github.com/ignw/cisco-aci-go-sdk/src/service"
)

// TODO: Validate - this one hasn't been checked, needs data in test case

func TestAccAciAppProfile_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciAppProfileDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciAppProfileConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciAppProfileExists("aci_app_profile.basic"),
					resource.TestCheckResourceAttr(
						"aci_app_profile.basic", "name", "app_profile1"),
					resource.TestCheckResourceAttr(
						"aci_app_profile.basic", "description", "terraform test App Profile"),
					resource.TestCheckResourceAttr(
						"aci_app_profile.basic", "domain_name", "uni/tn-IGNW-tenant1/ap-app_profile1"),
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

		app_profile, err := client.AppProfiles.Get(id)

		if err != nil || app_profile == nil {
			return fmt.Errorf("Error retreiving app_profile id: %s", id)
		}

		return nil
	}
}

func testAccCheckAciAppProfileDestroy(state *terraform.State) error {

	client := testAccProvider.Meta().(*cage.Client)

	err := checkDestroy("aci_app_profile.basic", state, func(s string) (interface{}, error) {
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

resource "aci_app_profile" "basic" {
	name = "app_profile1"
	description = "terraform test App Profile"
	tenant_id = "${aci_tenant.basic.id}"	
}

`
