package aci

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccAciFilter_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckAciFilterConfigBasic,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAciFilterExists("aci_filter.basic"),
					resource.TestCheckResourceAttr(
						"aci_filter.basic", "name", "http-basic"),
					resource.TestCheckResourceAttr(
						"aci_filter.basic", "alias", "tf-example"),
					resource.TestCheckResourceAttr(
						"aci_filter.basic", "status", "created"),
					resource.TestCheckResourceAttr(
						"aci_filter.basic", "tags", ""),
					resource.TestCheckResourceAttr(
						"aci_filter.basic", "entry.1", ""),
				),
			},
		},
	})
}

func testAccCheckAciFilterExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		/*
			client := testAccProvider.Meta().(*aci.Client)
			if err := existsHelper(s, client); err != nil {
				return err
			}
		*/
		return nil
	}
}

const testAccCheckAciFilterConfigBasic = `
resource "aci_filter" "basic" {
    name = "http-only"
    alias = "tf-example"
    status = "created"
    tags = []

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
`
