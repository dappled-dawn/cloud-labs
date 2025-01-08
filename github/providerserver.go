package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
	"github.com/hashicorp/terraform-plugin-go/tfprotov5"
	"github.com/hashicorp/terraform-plugin-mux/tf5muxserver"
)

func ProviderServer(ctx context.Context) (tfprotov5.ProviderServer, error) {
	sdkServerFunc := func() tfprotov5.ProviderServer {
		return Provider().GRPCProvider()
	}

	frameworkServerFunc := providerserver.NewProtocol5(NewFrameworkProvider())

	return tf5muxserver.NewMuxServer(ctx, sdkServerFunc, frameworkServerFunc)
}
