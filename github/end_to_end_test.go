package github_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/integrations/terraform-provider-github/v6/github"
)

var realProviderFactories map[string]func() (tfprotov5.ProviderServer, error)

func init() {
	realProviderFactories = map[string]func() (tfprotov5.ProviderServer, error){
		"github": func() (tfprotov5.ProviderServer, error) {
			ctx := context.Background()
			return github.ProviderServer(ctx)
		},
	}
}

const anonymous = "anonymous"
const individual = "individual"
const organization = "organization"
const enterprise = "enterprise"

func skipUnlessMode(t *testing.T, providerMode string) {
	switch providerMode {
	case anonymous:
		if os.Getenv("GITHUB_BASE_URL") != "" &&
			os.Getenv("GITHUB_BASE_URL") != "https://api.github.com/" {
			t.Log("anonymous mode not supported for GHES deployments")
			break
		}

		if os.Getenv("GITHUB_TOKEN") == "" {
			return
		} else {
			t.Log("GITHUB_TOKEN environment variable should be empty")
		}
	case enterprise:
		if os.Getenv("GITHUB_TOKEN") == "" {
			t.Log("GITHUB_TOKEN environment variable should be set")
		} else {
			return
		}

	case individual:
		if os.Getenv("GITHUB_TOKEN") != "" && os.Getenv("GITHUB_OWNER") != "" {
			return
		} else {
			t.Log("GITHUB_TOKEN and GITHUB_OWNER environment variables should be set")
		}
	case organization:
		if os.Getenv("GITHUB_TOKEN") != "" && os.Getenv("GITHUB_ORGANIZATION") != "" {
			return
		} else {
			t.Log("GITHUB_TOKEN and GITHUB_ORGANIZATION environment variables should be set")
		}
	}

	t.Skipf("Skipping %s which requires %s mode", t.Name(), providerMode)
}

func TestAccEndToEnd(t *testing.T) {

	randomID := acctest.RandStringFromCharSet(5, acctest.CharSetAlphaNum)

	t.Run("creates a team configured with defaults", func(t *testing.T) {

		config := fmt.Sprintf(`
			resource "github_team" "test" {
				name         = "tf-acc-%s"
			}
		`, randomID)

		check := resource.ComposeTestCheckFunc(
			resource.TestCheckResourceAttrSet("github_team.test", "slug"),
		)

		testCase := func(t *testing.T, mode string) {
			resource.Test(t, resource.TestCase{
				PreCheck:                 func() { skipUnlessMode(t, mode) },
				ProtoV5ProviderFactories: realProviderFactories,
				Steps: []resource.TestStep{
					{
						Config: config,
						Check:  check,
					},
				},
			})
		}

		t.Run("with an anonymous account", func(t *testing.T) {
			t.Skip("anonymous account not supported for this operation")
		})

		t.Run("with an individual account", func(t *testing.T) {
			t.Skip("individual account not supported for this operation")
		})

		t.Run("with an organization account", func(t *testing.T) {
			testCase(t, organization)
		})

	})
}
