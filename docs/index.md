---
organization: Turbot
category: ["media"]
icon_url: "/images/plugins/turbot/hackernews.svg"
brand_color: "#FF6600"
display_name: Hacker News
name: hackernews
description: Steampipe plugin to query stories, items and users from Hacker News.
---

# Hacker News

[Hacker News](https://news.ycombinator.com) is a social news website focusing on computer science and entrepreneurship.

## Installation

To download and install the latest hackernews plugin:

```bash
steampipe plugin install hackernews
```

## Credentials

The [Hacker News API](https://github.com/HackerNews/API) is open to the public and does not require any credentials.

## Connection Configuration

Connection configurations are defined using HCL in one or more Steampipe config files. Steampipe will load ALL configuration files from `~/.steampipe/config` that have a `.spc` extension. A config file may contain multiple connections.

Installing the latest hackernews plugin will create a default connection named `hackernews` in the `~/.steampipe/config/hackernews.spc` file. You may edit this connection to set options:

```hcl
connection "hackernews" {
  plugin    = "hackernews"
  max_items = 5000
}
```
