terraform {
  required_providers {
    godotenv = {
      source  = "briceburg/go-dotenv"
      version = "~> 0.0.0"
    }
  }
}

output "env_map" {
  # Read the contents of a .env file and unmarshal into a map
  value = provider::godotenv::read("${path.root}/.env")
}

/*
# example .env file
###################

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
*/
