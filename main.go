package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-mux/tf5muxserver"
	"github.com/integrations/terraform-provider-github/v6/github"
)

func main() {
	ctx := context.Background()
	server, err := github.MuxServer(ctx)
	if err != nil {
		panic(err)
	}

	tf5muxserver.Serve(ctx, server)
}
