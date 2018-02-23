# aci_contract

The Contract resource allows the creation and management of an ACI contract policies.

## Example Usage

```hcl
# Create a Contract
resource "aci_contract" "example" {
    name = "terraform-example"
    alias = "tf-example"    
    status = "created"
    tags = []
    subject = "http traffic"
    filters = ["${aci.filters.http}"]
    endpoint_groups = ["${aci.endpoint_groups.web.name}", "${aci.endpoint_groups.db.name}"]
}

# Create filter
resource "aci_filter" "http" {
    name = "http-only"
    protocol = "6"
    to-port = "80"
    ...
}

# Create web tier endpoint group
resource "aci_endpoint_group" "web" {
    name = "web"
    ...
}

# Create database tier endpoint group
resource "aci_endpoint_group" "db" {
    name = "db"
    ...
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required) A unique name for the tenant.
* `alias` - (Optional) The alternate name for the tenant.
* `status` - (Required) The tenant status, one of: `created`, `created,modified`, `modified`, `deleted`.  
* `tags` - (Optional) User defined meta data that can be applied to the tenant.
* `endpoint_groups` - (Required) Endpoint groups assigned to this application profile.