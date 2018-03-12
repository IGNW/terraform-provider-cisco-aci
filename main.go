package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/ignw/terraform-provider-cisco-aci/aci"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: aci.Provider})
}