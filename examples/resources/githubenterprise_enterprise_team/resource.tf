resource "githubenterprise_enterprise_team" "example" {
  enterprise = "example-enterprise"
  name       = "GitHub Example Team"

  organization_selection_type = "selected"
  organization_slugs          = ["example-org", "another-org"]

  group_id = "11111111-1111-1111-1111-111111111111"
}
