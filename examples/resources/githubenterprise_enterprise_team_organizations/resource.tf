resource "githubenterprise_enterprise_team_organizations" "example" {
  enterprise         = "example-enterprise"
  team_slug          = "example-team"
  organization_slugs = ["example-org", "another-org"]
}
