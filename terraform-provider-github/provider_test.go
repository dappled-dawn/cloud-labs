package main // todo: main_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestProvider(t *testing.T) {
	resource.Test(t, resource.TestCase{
		ProtoV5ProviderFactories: map[string]func() (tfprotov5.ProviderServer, error){
			"github": providerserver.NewProtocol5WithError(providerFactory()),
		},
		Steps: []resource.TestStep{},
	})
}
