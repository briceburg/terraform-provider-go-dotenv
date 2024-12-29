# Terraform Provider GoDotEnv

Unmarshal .env files for use in Terraform using [godotenv](https://github.com/joho/godotenv) - A Go port of Ruby's dotenv library.

Say yes to parsing of diverse `.env` formats including;

* unquoted values
* multiline values
* comments
* export declarations

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
mDMEZy7NRhYJKwYBBAHaRw8BAQdAcB7cXN8uImfM4u74hu8z+XrEdDf3k61NiTm8
qfodop20JHRvcnZhbGRzQGtlcm5lbC5vcmcgPExpbnVzIFRvcnZhbGRzPohyBBMW
CAAaBAsJCAcCFQgCFgECGQAFgmcuzYACngECmwMACgkQOO+F4ZFUo9uA1gEAxntS
RQlsnKscWz6Ze2JGiy4HZdPz3mVYjB/kYB9PMzoBAI+tfYeAEj/89BunEgg23lQR
FefKwrLMI5KDa3C2peMLtAQgPCA+iGEEMBYIAAkCHQAFgmcuzVYACgkQOO+F4ZFU
o9tSxgD/acJj3LGDqC1XGlSk0tMJIujyfyoNevSlR8R40nqmVl8A/3HG4ThHza2d
xyLBjTj9EyoFDSvAr304e13SttGbKPECiHIEExYIABoECwkIBwIVCAIWAQIZAQWC
Zy7NRgKeAQKbAwAKCRA474XhkVSj24GfAP4vEti7AesQZEMzl4HnYs/bmVCs9XaX
DTNSwMEUCTyfGQD/QElmAdS6VokDyP95QdwFCml5vkFFJEWC8CkuPrBR0Aa4OARn
Ls1GEgorBgEEAZdVAQUBAQdANY5NF1sHqv+0q3pzIK56Aug7lkswAXnaDq7CHZ3H
WU0DAQgHiGEEGBYIAAkFgmcuzUYCmwwACgkQOO+F4ZFUo9uLrgEAkS4dxxwEM7ow
ipkcQ6YYFDzGF6DxEEmgVbqtH2OEps8A/RZvhadroI4ZNm2DqF1pzRpWr995Fmqc
4oCtZwT43W0A
=0pl0
-----END PGP PUBLIC KEY BLOCK-----"
```

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
      source = "briceburg/go-dotenv"
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
