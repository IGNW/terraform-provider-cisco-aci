# aci_filter

The Filter resource allows the creation and management of an ACI filters.

## Example Usage

```hcl
# Create filter
resource "aci_filter" "http" {
    name = "http-only"
    alias = "tf-example"    
    status = "created"
    tags = []
    entries {    
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
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) A unique name for the filter.
* `alias` - (Optional) The alternate name for the filter.
* `status` - (Required) The tenant status, one of: `created`, `created,modified`, `modified`, `deleted`.  
* `tags` - (Optional) User defined meta data that can be applied to the filter.
* `entry` - (Required) At least one filter entry must be specified. See entry below.

### Entry
* `protocol` - (Required) User defined meta data that can be applied to the filter.
* `source-from` - (Optional) Source host beginning port
* `source-to` - (Optional) Source host allowed ending port
* `destination-from` - (Optional) Destination host beginning port
* `destination-to` - (Optional) Destination host ending port