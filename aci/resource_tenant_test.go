package aci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	cage "github.com/ignw/cisco-aci-go-sdk/src/service"
)

func TestAccAciTenant_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciTenantConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciTenantExists("aci_tenant.basic"),
					resource.TestCheckResourceAttr(
						"aci_tenant.basic", "name", "IGNW-tenant1"),
					resource.TestCheckResourceAttr(
						"aci_tenant.basic", "description", "my first tenant"),
					resource.TestCheckResourceAttr(
						"aci_tenant.basic", "domain_name", "uni/tn-IGNW-tenant1"),
				),
			},
		},
	})
}

func testAccCheckAciTenantExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*cage.Client)

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Error -> Resource Not found: %s", n)
		}

		id := rs.Primary.Attributes["id"]

		tenant, err := client.Tenants.Get(id)

		if err != nil {
			return fmt.Errorf("Error retreiving tenant id: %s", id)
		}

		if tenant == nil {
			return fmt.Errorf("Error retreiving tenant id: %s", id)
		}

		return nil
	}
}

const testAccCheckAciTenantConfigBasic = `
resource "aci_tenant" "basic" {
    name = "IGNW-tenant1"
    description = "my first tenant"
}
`
