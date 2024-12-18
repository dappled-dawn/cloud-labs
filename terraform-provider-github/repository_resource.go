package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
)

type repositoryResource struct {
}

// Metadata should return the full name of the resource, such as
// examplecloud_thing.
func (r *repositoryResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_repository"
}

// Schema should return the schema for this resource.
func (r *repositoryResource) Schema(_ context.Context, _ resource.SchemaRequest, _ *resource.SchemaResponse) {
}

// ImportState is called when the provider must import a resource.
func (r *repositoryResource) ImportState(_ context.Context, _ resource.ImportStateRequest, _ *resource.ImportStateResponse) {
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
func (r *repositoryResource) Read(_ context.Context, _ resource.ReadRequest, _ *resource.ReadResponse) {
	panic("not implemented") // TODO: Implement
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
