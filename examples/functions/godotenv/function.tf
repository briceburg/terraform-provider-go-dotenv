# Copyright (c) HashiCorp, Inc.

output "yaml" {
  value = yamlencode(provider::godotenv::read("${path.root}/base.env", "production.env"))
}
