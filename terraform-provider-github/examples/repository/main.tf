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
  full_name = "dappled-dawn/cloud-labs"
}

output "repository" {
  value = data.github_repository.self.ssh_clone_url
}

