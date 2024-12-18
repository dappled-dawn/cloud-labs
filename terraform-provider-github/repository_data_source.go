package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/v67/github"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
)

type repositoryDataSource struct {
}

func repositoryDataSourceFactory() datasource.DataSource {
	return &repositoryDataSource{}
}

// Metadata should return the full name of the data source, such as
// examplecloud_thing.
func (r *repositoryDataSource) Metadata(_ context.Context, _ datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = "github_repository"
}

// Schema should return the schema for this data source.
func (r *repositoryDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema.Attributes = map[string]schema.Attribute{
		"full_name": schema.StringAttribute{
			Description: "The repository name",
			Optional:    true,
		},
		"description": schema.StringAttribute{
			Description: "The repository description",
			Computed:    true,
		},
	}
}

// Read is called when the provider must read data source values in
// order to update state. Config values should be read from the
// ReadRequest and new state values set on the ReadResponse.
func (r *repositoryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var fullName string
	req.Config.GetAttribute(ctx, path.Root("full_name"), &fullName)
	var segments = strings.Split(fullName, "/")

	if len(segments) != 2 {
		resp.Diagnostics.AddError("invalid repository name", "repository name must be in the format 'owner/name'")
		return
	}

	client := github.NewClient(nil)
	repo, _, err := client.Repositories.Get(ctx, segments[0], segments[1])
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("failed to get repository %s", fullName), err.Error())
		return
	}

	var description = repo.GetDescription()
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("description"), description)...)
}
