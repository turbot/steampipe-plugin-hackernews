---
title: "Steampipe Table: hackernews_show_hn - Query Hacker News 'Show HN' Posts using SQL"
description: "Allows users to query 'Show HN' posts on Hacker News, providing access to data about user-submitted posts that showcase new projects, products, or interesting initiatives."
---

# Table: hackernews_show_hn - Query Hacker News 'Show HN' Posts using SQL

Hacker News is a social news website focusing on computer science and entrepreneurship. It is run by Paul Graham's investment fund and startup incubator, Y Combinator. In essence, it's a place where users submit and share content, with 'Show HN' being a specific category where users showcase new projects, products, or interesting initiatives.

## Table Usage Guide

The `hackernews_show_hn` table provides insights into 'Show HN' posts on Hacker News. As a data analyst, explore post-specific details through this table, including post titles, URLs, points, and associated metadata. Utilize it to uncover information about posts, such as the most popular projects, trending topics, and user engagement metrics.

## Examples

### Show HN stories by score
Explore the popularity of Hacker News 'Show HN' stories by ranking them according to their score. This can help identify trending topics and influential discussions.

```sql
select
  *
from
  hackernews_show_hn
order by
  score desc
```

### Show HN stories with most comments
Explore the most discussed 'Show HN' stories on HackerNews. This query is useful for identifying trending topics or popular discussions within the community.

```sql
select
  *
from
  hackernews_show_hn
order by
  descendants desc
```