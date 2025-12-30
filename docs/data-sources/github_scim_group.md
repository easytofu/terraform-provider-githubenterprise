---
title: githubenterprise_github_scim_group
page_title: "Data Source: githubenterprise_github_scim_group"
description: "Fetches a GitHub SCIM group by display name. Requires a GitHub token configured on the provider."
slug: provider_datasource_githubenterprise_github_scim_group
category:
  uri: GitHub Enterprise Terraform Provider
parent:
  uri: provider_datasources
privacy:
  view: public
position: 7
---
## Data Source: githubenterprise_github_scim_group

Fetches a GitHub SCIM group by display name. Requires a GitHub token configured on the provider.

## Example Usage

```terraform
data "githubenterprise_github_scim_group" "example" {
  enterprise   = "example-enterprise"
  display_name = "GitHub Example Repository Reader"
}

output "github_scim_group_id" {
  value = data.githubenterprise_github_scim_group.example.id
}
```

## Schema

### Required

- `enterprise` (String) The enterprise slug.
- `display_name` (String) The SCIM group display name to look up.

### Read-Only

- `id` (String) The SCIM group ID.
- `external_id` (String) The external ID for the SCIM group.
- `members` (List of Object) Members returned by the SCIM group query. (see [below for nested schema](#nestedatt--members))

<a id="nestedatt--members"></a>
### Nested Schema for `members`

Read-Only:

- `value` (String) Member ID.
- `display` (String) Member display name.
- `type` (String) Member type.
- `ref` (String) Member reference URL.
