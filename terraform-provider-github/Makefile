.PHONY: build example toggle clean

PROJECT_NAME = terraform-provider-github

build:
	@mkdir -p bin
	go build -o bin/$(PROJECT_NAME)

example:
	@cd examples/repository && terraform init && terraform apply -auto-approve

test:
	TF_ACC=1 go test -v ./...
toggle:
	@if [[ -f ~/.terraformrc ]]; then mv ~/.terraformrc ~/.terraformrc.bak; elif [[ -f ~/.terraformrc.bak ]]; then  mv ~/.terraformrc.bak ~/.terraformrc; fi

clean:
	@rm -f examples/repository/*.tfstate
	@rm -rf examples/repository/.terraform
	@rm -rf examples/repository/.terraform.lock.hcl
