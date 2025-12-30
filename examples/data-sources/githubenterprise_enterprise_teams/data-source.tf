data "githubenterprise_enterprise_teams" "all" {
  enterprise = "example-enterprise"
}

output "github_enterprise_team_slugs" {
  value = [for team in data.githubenterprise_enterprise_teams.all.teams : team.slug]
}
