resource "githubenterprise_enterprise_team_member" "example" {
  enterprise = "example-enterprise"
  team_slug  = "example-team"
  username   = "octocat"
}
