package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/v67/github"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type repositoryResource struct {
	client *github.Client
}

// Metadata should return the full name of the resource, such as
// examplecloud_thing.
func (r *repositoryResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_repository"
}

// Configure is called when the provider is configured. This is where
// the provider should register its resources.
func (r *repositoryResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	providerData := req.ProviderData.(*ProviderData)
	r.client = providerData.client
}

// Schema should return the schema for this resource.
func (r *repositoryResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema.Attributes = map[string]schema.Attribute{
		"full_name": schema.StringAttribute{
			Description: "The full name of the repository.",
			Computed:    true,
		},
		"visibility": schema.StringAttribute{
			Description: "The visibility of the repository.",
			Computed:    true,
		},
	}
}

type ARepository struct {
	FullName   types.String `tfsdk:"full_name"`
	Visibility types.String `tfsdk:"visibility"`
}

func (a ARepository) Owner() string {
	var segments = strings.Split(a.FullName.ValueString(), "/")
	if len(segments) < 2 {
		return ""
	}
	return segments[0]
}

func (a ARepository) Name() string {
	var segments = strings.Split(a.FullName.ValueString(), "/")
	if len(segments) < 2 {
		return ""
	}
	return segments[1]
}

// ImportState is called when the provider must import a resource.
func (r *repositoryResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root("full_name"), req, resp)
}

// Create is called when the provider must create a new resource. Config
// and planned state values should be read from the
// CreateRequest and new state values set on the CreateResponse.
func (r *repositoryResource) Create(_ context.Context, _ resource.CreateRequest, _ *resource.CreateResponse) {
	panic("not implemented") // TODO: Implement
}

// Read is called when the provider must read resource values in order
// to update state. Planned state values should be read from the
// ReadRequest and new state values set on the ReadResponse.
func (r *repositoryResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var repo ARepository
	resp.Diagnostics.Append(req.State.Get(ctx, &repo)...)

	raw_repo, _, err := r.client.Repositories.Get(ctx, repo.Owner(), repo.Name())
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("failed to get repository %s/%s", repo.Owner(), repo.Name()), err.Error())
	}

	repo.FullName = types.StringValue(raw_repo.GetFullName())
	repo.Visibility = types.StringValue(raw_repo.GetVisibility())
	resp.Diagnostics.Append(resp.State.Set(ctx, repo)...)
}

// Update is called to update the state of the resource. Config, planned
// state, and prior state values should be read from the
// UpdateRequest and new state values set on the UpdateResponse.
func (r *repositoryResource) Update(_ context.Context, _ resource.UpdateRequest, _ *resource.UpdateResponse) {
	panic("not implemented") // TODO: Implement
}

// Delete is called when the provider must delete the resource. Config
// values may be read from the DeleteRequest.
//
// If execution completes without error, the framework will automatically
// call DeleteResponse.State.RemoveResource(), so it can be omitted
// from provider logic.
func (r *repositoryResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	panic("not implemented") // TODO: Implement
}

func repositoryResourceFactory() resource.Resource {
	return &repositoryResource{}
}
