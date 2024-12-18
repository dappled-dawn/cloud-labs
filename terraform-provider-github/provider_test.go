package main // todo: main_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestRepositoryDataSource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"github": providerserver.NewProtocol5WithError(providerFactory()),
		},
		Steps: []resource.TestStep{
			{
				ResourceName: "",
				Config: `
				data "github_repository" "example" {
				  full_name = "bbasata/shrinkwrap"
			        }
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.github_repository.example", "description", "As an app, it shortens all the URLs. As a code base, it serves as a sandbox for software design experiments with Ruby and Rails."),
					resource.TestCheckResourceAttr("data.github_repository.example", "visibility", "public"),
				),
				Destroy:            false,
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func TestRepositoryResource_Import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"github": providerserver.NewProtocol5WithError(providerFactory()),
		},
		Steps: []resource.TestStep{
			{
				ResourceName: "github_repository.shrinkwrap",
				Config: `
				import {
				  to = github_repository.shrinkwrap
				  id = "bbasata/shrinkwrap"
				}

				resource "github_repository" "shrinkwrap" {
				  full_name = "bbasata/shrinkwrap"
			        }
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("github_repository.shrinkwrap", "visibility", "public"),
				),
			},
		},
	})
}
