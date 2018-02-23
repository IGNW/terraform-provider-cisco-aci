# aci_tenant

The Tenant resource allows the creation and management of an ACI tenant.

## Example Usage

```hcl
# Create a Tenant
resource "aci_tenant" "example" {
    name = "terraform-example"
    alias = "tf-example"
    description = "This is a Terraform example tenant"
    status = "created"
    tags = []
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) A unique name for the tenant.
* `alias` - (Optional) The alternate name for the tenant.
* `description` - (Optional) A longer, human-readable description for the tenant.
* `status` - (Required) The tenant status, one of: `created`, `created,modified`, `modified`, `deleted`.  
* `tags` - (Optional) User defined meta data that can be applied to the tenant.