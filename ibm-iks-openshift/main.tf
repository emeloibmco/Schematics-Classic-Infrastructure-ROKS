resource "ibm_container_cluster" "cluster" {
  name              = "${var.cluster_name}${random_id.name.hex}"
  datacenter        = "${var.datacenter}"
  default_pool_size = 1
  machine_type      = "${var.machine_type}"
  hardware          = "${var.hardware}"
  kube_version      = "${var.kube_version}"
  public_vlan_id    = "${var.public_vlan_id}"
  private_vlan_id   = "${var.private_vlan_id}"
  lifecycle {
  ignore_changes = ["kube_version"]
  }
  resource_group_id = "6764d0853c60419180ba88daafce32c4"
  iaas_classic_username = "2059386_jhoinerrojas@ibm.com"
  iaas_classic_api_key = "a47020190de8b75e3084a1c9ddf3a5016658ca82cffcd17f92060abc7ee81d0b"
  ibmcloud_api_key = "8FynCmLa_2Xy0NXCWU1W6b3gMeeGDjW7ibLCPfBs5Tcy"
}

data "ibm_container_cluster_config" "cluster_config" {
  cluster_name_id = "${ibm_container_cluster.cluster.id}"
}

resource "random_id" "name" {
  byte_length = 4
}
