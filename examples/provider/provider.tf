terraform {
  required_providers {
    godotenv = {
      source  = "briceburg/godotenv"
      version = "~> 0.0.0"
    }
  }
}

output "env_map" {
  # Read the contents of a .env file and unmarshal into a map
  value = provider::godotenv::read("${path.root}/.env")
}
