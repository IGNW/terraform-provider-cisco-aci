provider "aci" {
  url            = "https://sandboxapicdc.cisco.com"
  username       = "admin"
  password       = "ciscopsdt"
  allow_insecure = true
}

resource "aci_tenant" "tennant_1" {
  name        = "ignw-tenant-1"
  description = "IGNW Tennant 1"
}
