terraform {
  required_version = ">= 0.12"
  backend "pg" {
    conn_str = "${var.backend_url}"
  }
}
variable "backend_url" {
  default = "postgres://postgres:admin@127.0.0.1/postgres?sslmode=disable"
}