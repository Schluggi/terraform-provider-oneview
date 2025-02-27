provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 3600
  ov_ifmatch    = "*"
}

# Reads Storage Pool resource
data "oneview_storage_pool" "storage_pool" {
  name = "CPG-SSD-AO"
}

output "oneview_storage_pool_value" {
  value = data.oneview_storage_pool.storage_pool.uri
}

