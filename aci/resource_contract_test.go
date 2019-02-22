package aci

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	cage "github.com/ignw/cisco-aci-go-sdk/src/service"
)

// TODO: Validate - this one hasn't been checked, needs tenant data in test case

//TODO: Add advanced test cases to cover setting properties and EPGs

func TestAccAciContract_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAciContractDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciContractConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciContractExists("aci_contract.basic"),
					resource.TestCheckResourceAttr(
						"aci_contract.basic", "name", "http-only"),
					resource.TestCheckResourceAttr(
						"aci_contract.basic", "description", "terraform test contract"),
					resource.TestCheckResourceAttr(
						"aci_contract.basic", "domain_name", "uni/tn-IGNW-tenant1/brc-http-only"),
					resource.TestCheckResourceAttr(
						"aci_contract.basic", "scope", "context"),
					resource.TestCheckResourceAttr(
						"aci_contract.basic", "dscp", "unspecified"),
				),
			},
		},
	})
}

func testAccCheckAciContractExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*cage.Client)

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Error -> Resource Not found: %s", n)
		}

		id := rs.Primary.Attributes["id"]

		contract, err := client.Contracts.Get(id)

		if err != nil || contract == nil {
			return fmt.Errorf("Error retreiving contract id: %s", id)
		}

		return nil
	}
}

// TODO: Refactor, bad pattern need to catch and merge erros and not return. We are leaving dangling resources.
func testAccCheckAciContractDestroy(state *terraform.State) error {

	client := testAccProvider.Meta().(*cage.Client)

	err := checkDestroy("aci_contract.basic", state, func(s string) (interface{}, error) {
		return client.Contracts.Get(s)
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

const testAccCheckAciContractConfigBasic = `

resource "aci_tenant" "basic" {
    name = "IGNW-tenant1"
    description = "my first tenant"
}

resource "aci_contract" "basic" {
    name = "http-only" 
	description = "terraform test contract"
	tenant_id = "${aci_tenant.basic.id}"
}
`
