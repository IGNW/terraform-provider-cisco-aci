provider "aci" {
  url            = "https://sandboxapicdc.cisco.com"
  username       = "admin"
  password       = "ciscopsdt"
  allow_insecure = true
}

resource "aci_tenant" "basic" {
  name = "IGNW-tenant1"
  description = "my first tenant"
}
