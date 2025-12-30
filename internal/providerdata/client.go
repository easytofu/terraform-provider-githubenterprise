package providerdata

import "github.com/easytofu/terraform-provider-githubenterprise/internal/githubapi"

// Client holds the GitHub client shared by provider resources and datasources.
type Client struct {
	GitHub *githubapi.Client
}
