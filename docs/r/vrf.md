# aci_vrf

The Virtual Routing and Forwarding object (VRF) resource allows the creation and management of an ACI tenant VRF resources.

## Example Usage

```hcl
# Create a VRF
resource "aci_vrf" "example" {
    name = "terraform-example"
    alias = "tf-example"    
    status = "created"
    tags = []
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) A unique name for the VRF.
* `alias` - (Optional) The alternate name for the VRF.
* `status` - (Required) The VRF status, one of: `created`, `created,modified`, `modified`, `deleted`.  
* `tags` - (Optional) User defined meta data that can be applied to the VRF.