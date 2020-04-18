variable "iaas_classic_username" {
  description = "2059386_jhoinerrojas@ibm.com"
}

variable "iaas_classic_api_key" {
  description = "a47020190de8b75e3084a1c9ddf3a5016658ca82cffcd17f92060abc7ee81d0b"
}

variable "ibmcloud_api_key" {
  description = "8FynCmLa_2Xy0NXCWU1W6b3gMeeGDjW7ibLCPfBs5Tcy"
}

provider "ibm" {
  iaas_classic_username    = "${var.iaas_classic_username}"
  iaas_classic_api_key     = "${var.iaas_classic_api_key}"
  ibmcloud_api_key    = "${var.ibmcloud_api_key}"
  generation	      = 1
  region             = "us-south"
}
