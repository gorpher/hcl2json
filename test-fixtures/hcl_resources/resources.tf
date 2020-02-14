// 创建一个
resource "tencentcloud_instance" "cvm_test" {
  instance_name     = "cvm-test"
  availability_zone = "${var.availability_zone}"
  image_id          = "img-d5bte9sz"
  instance_type     = "S2.SMALL1"
  system_disk_type  = "CLOUD_PREMIUM"
  security_groups = [
    "${tencentcloud_security_group.sg_test.id}",
  ]

  vpc_id                     = "${tencentcloud_vpc.vpc_test.id}"
  subnet_id                  = "${tencentcloud_subnet.subnet_test.id}"
  internet_max_bandwidth_out = 10
  count                      = 3
  password                   = "qwer@123"
  allocate_public_ip         = true

  dynamic "tags" {
    for_each = "${var.tags}"
    content {
      key   = "${tags.value.key}"
      value = "${tags.value.value}"
    }
  }

}

data "tencentcloud_availability_zones" "my_favorate_zones" {}


// CLOUD_BASIC", "CLOUD_PREMIUM", "CLOUD_SSD
resource "tencentcloud_cbs_storage" "my_storage" {
  storage_type      = "CLOUD_PREMIUM"
  storage_size      = 10
  period            = 1
  availability_zone = "${data.tencentcloud_availability_zones.my_favorate_zones.zones.0.name}"
  storage_name      = "${var.storage_name}"
}

// 挂盘

resource "tencentcloud_cbs_storage_attachment" "my-attachment" {
  storage_id  = "${tencentcloud_cbs_storage.my_storage.id}"
  instance_id = "${tencentcloud_instance.cvm_test.0.id}"
}
