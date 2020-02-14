provider "tencentcloud" {
  secret_id  = "${var.secret_id}"
  secret_key = "${var.secret_key}"
  region     = "${var.region}"
}
