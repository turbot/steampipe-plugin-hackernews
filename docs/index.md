---
organization: Turbot
category: ["media"]
icon_url: "/images/plugins/turbot/hackernews.svg"
brand_color: "#FF6600"
display_name: Hacker News
name: hackernews
description: Steampipe plugin to query stories, items and users from Hacker News.
og_description: "Query Hacker News with SQL! Open source CLI. No DB required."
og_image: "/images/plugins/turbot/hackernews-social-graphic.png"
engines: ["steampipe", "sqlite", "postgres", "export"]
---

# Hacker News + Steampipe

[Hacker News](https://news.ycombinator.com) is a social news website focusing on computer science and entrepreneurship. Steampipe marshalls the HN API data into queryable tables letting you interactivly explore it via our command line interface or your favorite SQL client. 

[Steampipe](https://steampipe.io) is an open-source zero-ETL engine to instantly query cloud APIs using SQL.

Example query:

```sql
select
  score,
  descendants as comments,
  title
from 
  hackernews_top
where
  type = 'story'
  and lower(title) like '%sql%'
order by
  score desc;
```
standard output (can use `.output` to change to `csv` or `json`):
```text
+-------+----------+---------------------------------------------------------------------+
| score | comments | title                                                               |
+-------+----------+---------------------------------------------------------------------+
| 242   | 300      | Query Hacker News API with SQL                                      |
| 121   | 127      | Why Uber Engineering Switched from Postgres to MySQL (2016)         |
| 70    | 12       | Show HN: QueryCal – calculate metrics from your calendars using SQL |
| 17    | 10       | Global Associative Arrays in PostgreSQL                             |
+-------+----------+---------------------------------------------------------------------+
```
## Documentation

- **[Table definitions & examples →](/plugins/turbot/hackernews/tables)**

## Get started

### Install

Download and install the latest Hacker News plugin:

```bash
steampipe plugin install hackernews
```

### Credentials

The [Hacker News API](https://github.com/HackerNews/API) is open to the public and does not require any credentials.


### Configuration

Connection configurations are defined using HCL in one or more Steampipe config files. Steampipe will load ALL configuration files from `~/.steampipe/config` that have a `.spc` extension. A config file may contain multiple connections.

Installing the latest hackernews plugin will create a default connection named `hackernews` in the `~/.steampipe/config/hackernews.spc` file. You may edit this connection to set options:

```hcl
connection "hackernews" {
  plugin    = "hackernews"
  max_items = 5000
}
```


