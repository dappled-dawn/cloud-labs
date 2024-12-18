package main // todo: main_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/statecheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestProvider(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"github": providerserver.NewProtocol5WithError(providerFactory()),
		},
		Steps: []resource.TestStep{
			{
				ResourceName: "",
				Taint:        []string{},
				Config: `
				data "github_repository" "example" {
				  full_name = "dappled-dawn/cloud-labs"
			        }
				`,
				Check:              resource.TestCheckResourceAttr("data.github_repository.example", "description", "hats"),
				Destroy:            false,
				ExpectNonEmptyPlan: false,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply:             []plancheck.PlanCheck{},
					PostApplyPreRefresh:  []plancheck.PlanCheck{},
					PostApplyPostRefresh: []plancheck.PlanCheck{},
				},
				RefreshPlanChecks: resource.RefreshPlanChecks{
					PostRefresh: []plancheck.PlanCheck{},
				},
				ConfigStateChecks:                    []statecheck.StateCheck{},
				PlanOnly:                             false,
				PreventDiskCleanup:                   false,
				PreventPostDestroyRefresh:            false,
				ImportState:                          false,
				ImportStateId:                        "",
				ImportStateIdPrefix:                  "",
				ImportStateIdFunc:                    func(*terraform.State) (string, error) { panic("not implemented") },
				ImportStateCheck:                     func([]*terraform.InstanceState) error { panic("not implemented") },
				ImportStateVerify:                    false,
				ImportStateVerifyIdentifierAttribute: "",
				ImportStateVerifyIgnore:              []string{},
				ImportStatePersist:                   false,
				RefreshState:                         false,
			},
		},
	})
}
