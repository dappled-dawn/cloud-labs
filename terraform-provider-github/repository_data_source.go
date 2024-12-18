package main

import (
	"context"
	"fmt"

	"github.com/google/go-github/v67/github"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
)

type repositoryDataSource struct {
	client  *github.Client
	service *github.RepositoriesService
}

func repositoryDataSourceFactory() datasource.DataSource {
	return &repositoryDataSource{}
}

// Metadata should return the full name of the data source, such as
// examplecloud_thing.
func (r *repositoryDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_repository"
}

func (r *repositoryDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	data, ok := req.ProviderData.(*ProviderData)
	if ok {
		r.client = data.client
		r.service = data.client.Repositories
	}
}

// Schema should return the schema for this data source.
func (r *repositoryDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema.Attributes = map[string]schema.Attribute{
		"full_name": schema.StringAttribute{
			CustomType:  OwnerWithName{},
			Description: "The repository name",
			Optional:    true,
		},
		"description": schema.StringAttribute{
			Description: "The repository description",
			Computed:    true,
		},
		"visibility": schema.StringAttribute{
			Description: "The repository visibility",
			Computed:    true,
		},
	}
}

// Read is called when the provider must read data source values in
// order to update state. Config values should be read from the
// ReadRequest and new state values set on the ReadResponse.
func (r *repositoryDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var repository Repository = Repository{}
	resp.Diagnostics.Append(req.Config.Get(ctx, &repository)...)

	repo, _, err := r.service.Get(ctx, repository.Owner(), repository.Name())
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("failed to get repository %s", repository.FullName.ValueString()), err.Error())
		return
	}

	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("description"), repo.GetDescription())...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("visibility"), repo.GetVisibility())...)
}
