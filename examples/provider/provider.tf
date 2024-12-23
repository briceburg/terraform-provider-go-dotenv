# Copyright (c) HashiCorp, Inc.

terraform {
  required_providers {
    godotenv = {
      source  = "briceburg/godotenv"
      version = "~> 0.0.0"
    }
  }
}

output "yaml" {
  value = yamlencode(provider::godotenv::read("${path.root}/base.env", "production.env"))
}
