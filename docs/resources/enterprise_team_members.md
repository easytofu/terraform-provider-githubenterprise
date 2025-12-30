---
title: githubenterprise_enterprise_team_members
page_title: "Resource: githubenterprise_enterprise_team_members"
description: "Manages enterprise team membership in bulk. Requires a GitHub token configured on the provider."
slug: provider_resource_githubenterprise_enterprise_team_members
category:
  uri: GitHub Enterprise Terraform Provider
parent:
  uri: provider_resources
privacy:
  view: public
position: 9
---
## Resource: githubenterprise_enterprise_team_members

Manages enterprise team membership in bulk. Requires a GitHub token configured on the provider.

## Example Usage

```terraform
resource "githubenterprise_enterprise_team_members" "example" {
  enterprise = var.github_enterprise_slug
  team_slug  = var.github_team_slug
  usernames  = var.github_team_member_usernames
}
```

## Schema

### Required

- `enterprise` (String) The enterprise slug.
- `team_slug` (String) The enterprise team slug or ID.
- `usernames` (Set of String) GitHub usernames to ensure are members of the enterprise team.

### Read-Only

- `id` (String) Composite ID of the team membership set.

## Import

Import is supported using the following syntax:

```shell
terraform import githubenterprise_enterprise_team_members.example <enterprise>/<team_slug>
```
