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

[Hacker News](https://news.ycombinator.com) is a social news website focusing on computer science and entrepreneurship. Steampipe marshalls the HN API data into queryable tables letting you interactivly explore it via our command line interface or your favorite SQL client. Example query:

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

## Installation

If you are just getting started with Steampipe, head over to https://steampipe.io/downloads to install the CLI (don't worry it will only take a minute). Once that is done to download and install the latest **Hacker News plugin**:

```bash
steampipe plugin install hackernews
```

## Browse the [Available Tables →](https://hub.steampipe.io/plugins/turbot/hackernews/tables)
 - [hackernews_ask_hn](https://hub.steampipe.io/plugins/turbot/hackernews/tables/hackernews_ask_hn) | Up to 200 of the latest Ask HN stories.
 - [hackernews_best](https://hub.steampipe.io/plugins/turbot/hackernews/tables/hackernews_best) | 500 of the latest "best" stories.
 - [hackernews_item](https://hub.steampipe.io/plugins/turbot/hackernews/tables/hackernews_item) | Stories, comments, jobs, Ask HNs and even polls are just items. This table includes the most recent items posted to Hacker News. `max_items` in the connection configuration defines the number of items returned by a list query to this table.
 - [hackernews_job](https://hub.steampipe.io/plugins/turbot/hackernews/tables/hackernews_job) | Up to 200 job listings.
 - [hackernews_new](https://hub.steampipe.io/plugins/turbot/hackernews/tables/hackernews_new) | The 500 newest stories.
 - [hackernews_show_hn](https://hub.steampipe.io/plugins/turbot/hackernews/tables/hackernews_show_hn) | Up to 200 of the latest "Show HN" stories.
 - [hackernews_top](https://hub.steampipe.io/plugins/turbot/hackernews/tables/hackernews_top) | 500 of the latest "top" stories.
 - [hackernews_user](https://hub.steampipe.io/plugins/turbot/hackernews/tables/hackernews_user) | Query information about users with public activity on Hacker News.

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
