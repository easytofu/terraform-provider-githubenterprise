resource "githubenterprise_enterprise_team_members" "example" {
  enterprise = "example-enterprise"
  team_slug  = "example-team"
  usernames  = ["octocat", "hubot"]
}
