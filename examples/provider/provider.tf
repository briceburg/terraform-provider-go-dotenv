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
