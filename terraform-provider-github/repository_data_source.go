package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
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
func (r *repositoryDataSource) Read(ctx context.Context, _ datasource.ReadRequest, resp *datasource.ReadResponse) {
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, _path.Root("description"), "This is a test repository")...)
}
