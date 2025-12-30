---
title: Overview
page_title: "GitHub Enterprise Provider"
description: "The GitHub Enterprise provider manages enterprise teams and SCIM group lookups."
slug: provider_overview
category:
  uri: GitHub Enterprise Terraform Provider
position: 0
privacy:
  view: public
---
# GitHub Enterprise Provider

The GitHub Enterprise provider manages enterprise teams and SCIM group lookups.

## Usage

```terraform
terraform {
  required_providers {
    githubenterprise = {
      source = "easytofu/githubenterprise"
    }
  }
}
```

```terraform
provider "githubenterprise" {
  github_token = var.github_token
}
```

## Authentication

The GitHub API requires a classic personal access token with enterprise admin permissions. Configure it in the provider block or via environment variables:

```shell
export GITHUB_TOKEN="..."
```

## Schema

### Optional

- `github_token` (String) GitHub classic personal access token for enterprise team APIs. Can also be set with the `GITHUB_TOKEN` environment variable.
- `github_api_base_url` (String) Optional override for the GitHub API base URL (default: `https://api.github.com`). Can also be set with the `GITHUB_API_BASE_URL` environment variable.
- `github_api_version` (String) Optional override for the GitHub API version header (default: `2022-11-28`). Can also be set with the `GITHUB_API_VERSION` environment variable.
