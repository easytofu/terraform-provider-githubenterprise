---
title: githubenterprise_enterprise_team_organization
page_title: "Resource: githubenterprise_enterprise_team_organization"
description: "Manages a single organization assignment for a GitHub enterprise team. Requires a GitHub token configured on the provider."
slug: provider_resource_githubenterprise_enterprise_team_organization
category:
  uri: GitHub Enterprise Terraform Provider
parent:
  uri: provider_resources
privacy:
  view: public
position: 6
---
## Resource: githubenterprise_enterprise_team_organization

Manages a single organization assignment for a GitHub enterprise team. Requires a GitHub token configured on the provider.

## Example Usage

```terraform
resource "githubenterprise_enterprise_team_organization" "example" {
  enterprise   = var.github_enterprise_slug
  team_slug    = var.github_team_slug
  organization = "example-org"
}
```

## Schema

### Required

- `enterprise` (String) The enterprise slug.
- `team_slug` (String) The enterprise team slug or ID.
- `organization` (String) The organization slug.

### Read-Only

- `id` (String) Composite ID of the assignment.

## Import

Import is supported using the following syntax:

```shell
terraform import githubenterprise_enterprise_team_organization.example <enterprise>/<team_slug>/<organization>
```
