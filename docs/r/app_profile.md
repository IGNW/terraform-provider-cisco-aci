# aci_app_profile

The App_Profile Profile resource allows the creation and management of an ACI application profiles.

## Example Usage

```hcl
# Create an Application Profile
resource "aci_app_profile" "example" {
    name = "terraform-example"
    alias = "tf-example"    
    status = "created"
    tags = []
    endpoint_groups = ["${aci.endpoint_groups.web.name}", "${aci.endpoint_groups.db.name}"]
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