terraform {
  required_providers {
    github = {
      source = "integrations/github"
    }
  }
}

provider "github" {
}

data "github_repository" "self" {
  full_name = "bbasata/shrinkwrap"
}

output "repository" {
  value = data.github_repository.self
}

