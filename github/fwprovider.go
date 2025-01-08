package github

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
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
	resp.Schema.Attributes = map[string]schema.Attribute{
		"token": schema.StringAttribute{
			Optional: true,
			Description: "The OAuth token used to connect to GitHub. Anonymous mode is " +
				"enabled if both `token` and `app_auth` are not set.",
		},
		"owner": schema.StringAttribute{
			Optional: true,
			Description: "The GitHub owner name to manage. Use this field instead of " +
				"`organization` when managing individual accounts.",
		},
		"retryable_errors": schema.ListAttribute{
			ElementType: types.Int32Type,
			Optional:    true,
			Description: "Allow the provider to retry after receiving an error status code, the max_retries should be set for this to work" +
				"Defaults to [500, 502, 503, 504]",
		},
		"max_retries": schema.Int32Attribute{
			Optional: true,
			Description: "Number of times to retry a request after receiving an error status code" +
				"Defaults to 3",
		},
		"organization": schema.StringAttribute{
			DeprecationMessage: "Use owner (or GITHUB_OWNER) instead of organization (or GITHUB_ORGANIZATION)",
			Optional:           true,
			Description: "The GitHub organization name to manage. Use this field instead " +
				"of `owner` when managing organization accounts.",
		},
		"base_url": schema.StringAttribute{
			Optional:    true,
			Description: "The GitHub Base API URL",
		},
		"insecure": schema.BoolAttribute{
			Optional:    true,
			Description: "Enable `insecure` mode for testing purposes",
		},
		"write_delay_ms": schema.Int32Attribute{
			Optional: true,
			Description: "Amount of time in milliseconds to sleep in between writes to GitHub API. " +
				"Defaults to 1000ms or 1s if not set.",
		},
		"read_delay_ms": schema.Int32Attribute{
			Optional: true,
			Description: "Amount of time in milliseconds to sleep in between non-write requests to GitHub API. " +
				"Defaults to 0ms if not set.",
		},
		"retry_delay_ms": schema.Int32Attribute{
			Optional: true,
			Description: "Amount of time in milliseconds to sleep in between requests to GitHub API after an error response. " +
				"Defaults to 1000ms or 1s if not set, the max_retries must be set to greater than zero.",
		},
		"parallel_requests": schema.BoolAttribute{
			Optional: true,
			Description: "Allow the provider to make parallel API calls to GitHub. " +
				"You may want to set it to true when you have a private Github Enterprise without strict rate limits. " +
				"Although, it is not possible to enable this setting on github.com " +
				"because we enforce the respect of github.com's best practices to avoid hitting abuse rate limits" +
				"Defaults to false if not set",
		},
	}
	resp.Schema.Blocks = map[string]schema.Block{
		"app_auth": schema.ListNestedBlock{
			Description: "The GitHub App credentials used to connect to GitHub. Conflicts with " +
				"`token`. Anonymous mode is enabled if both `token` and `app_auth` are not set.",
			NestedObject: schema.NestedBlockObject{
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Required:    true,
						Description: "The GitHub App ID.",
					},
					"installation_id": schema.StringAttribute{
						Required:    true,
						Description: "The GitHub App installation instance ID.",
					},
					"pem_file": schema.StringAttribute{
						Required:    true,
						Sensitive:   true,
						Description: "The GitHub App PEM file contents.",
					},
				},
			},
		},
	}
}

// Configure is called at the beginning of the provider lifecycle, when
// Terraform sends to the provider the values the user specified in the
// provider configuration block. These are supplied in the
// ConfigureProviderRequest argument.
// Values from provider configuration are often used to initialise an
// API client, which should be stored on the struct implementing the
// Provider interface.
func (f *FrameworkProvider) Configure(_ context.Context, _ provider.ConfigureRequest, _ *provider.ConfigureResponse) {
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
