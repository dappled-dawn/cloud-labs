package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

type githubProvider struct {
}

// Metadata should return the metadata for the provider, such as
// a type name and version data.
//
// Implementing the MetadataResponse.TypeName will populate the
// datasource.MetadataRequest.ProviderTypeName and
// resource.MetadataRequest.ProviderTypeName fields automatically.
func (g *githubProvider) Metadata(_ context.Context, _ provider.MetadataRequest, _ *provider.MetadataResponse) {
}

// Schema should return the schema for this provider.
func (g *githubProvider) Schema(_ context.Context, _ provider.SchemaRequest, _ *provider.SchemaResponse) {
}

// Configure is called at the beginning of the provider lifecycle, when
// Terraform sends to the provider the values the user specified in the
// provider configuration block. These are supplied in the
// ConfigureProviderRequest argument.
// Values from provider configuration are often used to initialise an
// API client, which should be stored on the struct implementing the
// Provider interface.
func (g *githubProvider) Configure(_ context.Context, _ provider.ConfigureRequest, _ *provider.ConfigureResponse) {
	panic("not implemented") // TODO: Implement
}

// DataSources returns a slice of functions to instantiate each DataSource
// implementation.
//
// The data source type name is determined by the DataSource implementing
// the Metadata method. All data sources must have unique names.
func (g *githubProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		repositoryDataSourceFactory,
	}
}

// Resources returns a slice of functions to instantiate each Resource
// implementation.
//
// The resource type name is determined by the Resource implementing
// the Metadata method. All resources must have unique names.
func (g *githubProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func providerFactory() provider.Provider {
	return &githubProvider{}
}
