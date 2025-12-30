package provider

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/easytofu/terraform-provider-githubenterprise/internal/datasources"
	"github.com/easytofu/terraform-provider-githubenterprise/internal/githubapi"
	"github.com/easytofu/terraform-provider-githubenterprise/internal/providerdata"
	"github.com/easytofu/terraform-provider-githubenterprise/internal/resources"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &GitHubEnterpriseProvider{}
)

// GitHubEnterpriseProvider is the provider implementation.
type GitHubEnterpriseProvider struct {
	Version string
}

// GitHubEnterpriseProviderModel describes the provider data model.
type GitHubEnterpriseProviderModel struct {
	GithubToken      types.String `tfsdk:"github_token"`
	GithubAPIBaseURL types.String `tfsdk:"github_api_base_url"`
	GithubAPIVersion types.String `tfsdk:"github_api_version"`
}

// New is a helper function to simplify the provider implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &GitHubEnterpriseProvider{
			Version: version,
		}
	}
}

// Metadata returns the provider type name.
func (p *GitHubEnterpriseProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "githubenterprise"
	resp.Version = p.Version
}

// Schema defines the provider-level schema for configuration.
func (p *GitHubEnterpriseProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "The GitHub Enterprise provider manages enterprise teams and SCIM integrations.",
		Attributes: map[string]schema.Attribute{
			"github_token": schema.StringAttribute{
				Description: "GitHub classic personal access token for enterprise team APIs. Can also be set with the `GITHUB_TOKEN` environment variable.",
				Optional:    true,
				Sensitive:   true,
			},
			"github_api_base_url": schema.StringAttribute{
				Description: "Optional override for the GitHub API base URL (default: `https://api.github.com`). Can also be set with the `GITHUB_API_BASE_URL` environment variable.",
				Optional:    true,
			},
			"github_api_version": schema.StringAttribute{
				Description: "Optional override for the GitHub API version header (default: `2022-11-28`). Can also be set with the `GITHUB_API_VERSION` environment variable.",
				Optional:    true,
			},
		},
	}
}

// Configure prepares a GitHub client for data sources and resources.
func (p *GitHubEnterpriseProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var config GitHubEnterpriseProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	githubToken := os.Getenv("GITHUB_TOKEN")
	if !config.GithubToken.IsNull() {
		githubToken = config.GithubToken.ValueString()
	}

	githubAPIBaseURL := os.Getenv("GITHUB_API_BASE_URL")
	if !config.GithubAPIBaseURL.IsNull() {
		githubAPIBaseURL = strings.TrimSpace(config.GithubAPIBaseURL.ValueString())
	}

	githubAPIVersion := os.Getenv("GITHUB_API_VERSION")
	if !config.GithubAPIVersion.IsNull() {
		githubAPIVersion = strings.TrimSpace(config.GithubAPIVersion.ValueString())
	}

	clientData := &providerdata.Client{}

	if githubToken != "" {
		userAgent := "terraform-provider-githubenterprise"
		if p.Version != "" {
			userAgent = fmt.Sprintf("terraform-provider-githubenterprise/%s", p.Version)
		}

		githubClient, err := githubapi.NewClient(githubToken, githubAPIBaseURL, githubAPIVersion, userAgent, nil)
		if err != nil {
			resp.Diagnostics.AddError(
				"Unable to create GitHub client",
				fmt.Sprintf("An error occurred when creating the GitHub client: %s", err),
			)
			return
		}

		clientData.GitHub = githubClient
	}

	resp.DataSourceData = clientData
	resp.ResourceData = clientData
}

// Resources defines the resources implemented in the provider.
func (p *GitHubEnterpriseProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		resources.NewGithubEnterpriseTeamResource,
		resources.NewGithubEnterpriseTeamOrganizationResource,
		resources.NewGithubEnterpriseTeamOrganizationsResource,
		resources.NewGithubEnterpriseTeamMemberResource,
		resources.NewGithubEnterpriseTeamMembersResource,
	}
}

// DataSources defines the data sources implemented in the provider.
func (p *GitHubEnterpriseProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		datasources.NewGithubEnterpriseTeamsDataSource,
		datasources.NewGithubScimGroupDataSource,
	}
}
