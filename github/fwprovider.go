package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

type FrameworkProvider struct {
}

func NewFrameworkProvider() provider.Provider {
	return &FrameworkProvider{}
}

// Metadata should return the metadata for the provider, such as
// a type name and version data.
//
// Implementing the MetadataResponse.TypeName will populate the
// datasource.MetadataRequest.ProviderTypeName and
// resource.MetadataRequest.ProviderTypeName fields automatically.
func (f *FrameworkProvider) Metadata(_ context.Context, _ provider.MetadataRequest, _ *provider.MetadataResponse) {
}

// Schema should return the schema for this provider.
func (f *FrameworkProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema.Attributes = map[string]schema.Attribute{}
}

// Configure is called at the beginning of the provider lifecycle, when
// Terraform sends to the provider the values the user specified in the
// provider configuration block. These are supplied in the
// ConfigureProviderRequest argument.
// Values from provider configuration are often used to initialise an
// API client, which should be stored on the struct implementing the
// Provider interface.
func (f *FrameworkProvider) Configure(_ context.Context, _ provider.ConfigureRequest, _ *provider.ConfigureResponse) {
	panic("not implemented") // TODO: Implement
}

// DataSources returns a slice of functions to instantiate each DataSource
// implementation.
//
// The data source type name is determined by the DataSource implementing
// the Metadata method. All data sources must have unique names.
func (f *FrameworkProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

// Resources returns a slice of functions to instantiate each Resource
// implementation.
//
// The resource type name is determined by the Resource implementing
// the Metadata method. All resources must have unique names.
func (f *FrameworkProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}
