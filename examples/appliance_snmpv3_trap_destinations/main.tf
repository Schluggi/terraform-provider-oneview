provider "oneview" {
  ov_username   = var.username
  ov_password   = var.password
  ov_endpoint   = var.endpoint
  ov_sslverify  = var.ssl_enabled
  ov_apiversion = 3000
  ov_ifmatch    = "*"
}

# Creates SNMPv3 Trap Destination
resource "oneview_snmpv3_trap_destinations" "snmptrap" {
    destination_address = "1.1.1.1"
    port = 162
    user_id = "41b96bbb-8f31-44e1-a3aa-8681e3d7c56c"
}