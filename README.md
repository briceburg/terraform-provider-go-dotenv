# Terraform Provider GoDotEnv

Unmarshal .env files for use in Terraform using [godotenv](https://github.com/joho/godotenv) - A Go port of Ruby's dotenv library.

## Requirements

- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.8

## Using the provider


### Read function 

The `read` function will read from one or more .env files and return a map of key/value pairs. This map can be passed to all Terraform entities.

The example below reads from the `base.env` and `production.env` files and passes output to the yamldecode function to create a yaml representation of the environment.

```terraform
terraform {
  required_providers {
    godotenv = {
      source = "briceburg/godotenv"
      version = "~> 0.0.0"
    }
  }
}

output "yaml" {
  value = yamlencode(provider::godotenv::read("${path.root}/base.env", "production.env"))
}

```

## Developing the Provider

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `make generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.
