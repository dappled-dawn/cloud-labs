.PHONY: build

PROJECT_NAME = terraform-provider-github

build:
	@mkdir -p bin
	go build -o bin/$(PROJECT_NAME)

example:
	@cd examples/repository
	terraform init && terraform apply

toggle:
	@if [[ -f ~/.terraformrc ]]; then mv ~/.terraformrc ~/.terraformrc.bak; elif [[ -f ~/.terraformrc.bak ]]; then  mv ~/.terraformrc.bak ~/.terraformrc; fi
