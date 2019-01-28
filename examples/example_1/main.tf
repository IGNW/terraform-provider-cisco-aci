provider "aci" {
  url            = "https://url:port"
  username       = "admin"
  password       = "password"
  allow_insecure = true
}

resource "aci_tenant" "tennant_1" {
  name        = "ignw-tenant-1"
  description = "IGNW Tennant 1"
}
