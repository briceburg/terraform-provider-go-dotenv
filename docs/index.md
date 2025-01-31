---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "go-dotenv Provider"
subcategory: ""
description: |-
  Unmarshal .env files into Terraform using godotenv https://github.com/joho/godotenv - A Go port of Ruby's dotenv library.
  Say yes to parsing of diverse .envformats including;
  unquoted valuesmultiline valuescommentsexport declarations
  example .env file
  
  # comments are supported
  PROFILES"local dev" # and trailing comments
  
  # so are unquoted values
  FOO=fizz buzz
  
  # and export declarations
  export BAR="foobar"
  
  # multiline values too (both inline and newline separated)!
  STOOGES="larry\ncurly\nmoe"
  PGP_LINUS="-----BEGIN PGP PUBLIC KEY BLOCK-----
  ...lots of base64 encoded text...
  -----END PGP PUBLIC KEY BLOCK-----"
---

# go-dotenv Provider

Unmarshal .env files into Terraform using [godotenv](https://github.com/joho/godotenv) - A Go port of Ruby's dotenv library.

Say yes to parsing of diverse `.env`formats including;

* unquoted values
* multiline values
* comments
* export declarations

##### example .env file
```sh
# comments are supported
PROFILES"local dev" # and trailing comments

# so are unquoted values
FOO=fizz buzz

# and export declarations
export BAR="foobar"

# multiline values too (both inline and newline separated)!
STOOGES="larry\ncurly\nmoe"
PGP_LINUS="-----BEGIN PGP PUBLIC KEY BLOCK-----
...lots of base64 encoded text...
-----END PGP PUBLIC KEY BLOCK-----"
```

## Example Usage

```terraform
terraform {
  required_providers {
    go-dotenv = {
      source  = "briceburg/go-dotenv"
      version = "~> 0.0.0"
    }
  }
}

output "env_map" {
  # Read the contents of a .env file and unmarshal into a map
  value = provider::godotenv::read("${path.root}/.env")
}
```

<!-- schema generated by tfplugindocs -->
## Schema
