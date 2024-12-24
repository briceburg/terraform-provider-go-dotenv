locals {
  # Read the contents of a .env file and unmarshal into a map
  env_map = provider::godotenv::read("${path.root}/.env")

  # Merge several .env files and encode as a yaml string
  env_yaml = yamlencode(provider::godotenv::read("base.env", "production.env"))
}
