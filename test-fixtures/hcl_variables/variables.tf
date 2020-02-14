variable "alias" {
  description = "The alias of vpc,must be unique in project"
  type = "list"
  default = [
    "别名零",
    "别名一",
    "别名二",
    "别名三",
    "别名四"]
}

variable "tags" {
  description = "defualt值不要用map,hcl转json,map类型会被转成list,导致引用bug,建议使用list类型代替map"
  default = [
    {
      key = "k1",
      value = "v1"
    },
    {
      key = "k2",
      value = "v2"
    },
    {
      key = "k3",
      value = "v3"
    }
  ]
}