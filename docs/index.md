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

```
> .inspect hackernews
+--------------------+-----------------------------------------------------------------------------+
| TABLE              | DESCRIPTION                                                                 |
+--------------------+-----------------------------------------------------------------------------+
| hackernews_ask_hn  | Latest 200 Ask HN stories.                                                  |
| hackernews_best    | Best 500 stories.                                                           |
| hackernews_item    | Stories, comments, jobs, Ask HNs and even polls are just items. This table  |
|                    | includes the most recent items posted to Hacker News.                       |
| hackernews_job     | Latest 200 Job stories.                                                     |
| hackernews_new     | Newest 500 stories.                                                         |
| hackernews_show_hn | Latest 200 Show HN stories.                                                 |
| hackernews_top     | Top 500 stories.                                                            |
| hackernews_user    | Information about Hacker News registered users who have public activity.    |
+--------------------+-----------------------------------------------------------------------------+
```

```
> .inspect hackernews_item
+-------------+-----------------------------+---------------------------------------------------------------------------+
| COLUMN      | TYPE                        | DESCRIPTION                                                               |
+-------------+-----------------------------+---------------------------------------------------------------------------+
| by          | text                        | The username of the item's author.                                        |
| dead        | boolean                     | True if the item is dead.                                                 |
| deleted     | boolean                     | True if the item is deleted.                                              |
| descendants | bigint                      | In the case of stories or polls, the total comment count.                 |
| id          | bigint                      | The item's unique id.                                                     |
| kids        | jsonb                       | The ids of the item's comments, in ranked display order.                  |
| parent      | bigint                      | The comment's parent: either another comment or the relevant story.       |
| parts       | jsonb                       | A list of related pollopts, in display order.                             |
| poll        | bigint                      | The pollopt's associated poll.                                            |
| score       | bigint                      | The story's score, or the votes for a pollopt.                            |
| text        | text                        | The comment, story or poll text. HTML.                                    |
| time        | timestamp without time zone | Timestamp when the item was created.                                      |
| title       | text                        | The title of the story, poll or job. HTML.                                |
| type        | text                        | The type of item. One of "job", "story", "comment", "poll", or "pollopt". |
| url         | text                        | The URL of the story.                                                     |
+-------------+-----------------------------+---------------------------------------------------------------------------+
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
