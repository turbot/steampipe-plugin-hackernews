---
title: "Steampipe Table: hackernews_top - Query Hacker News Top Stories using SQL"
description: "Allows users to query Top Stories in Hacker News, specifically the most popular posts, providing insights into trending topics and discussions."
---

# Table: hackernews_top - Query Hacker News Top Stories using SQL

Hacker News is a social news website focusing on computer science and entrepreneurship. It is run by Paul Graham's investment fund and startup incubator, Y Combinator. In general, content that can be submitted is defined as "anything that gratifies one's intellectual curiosity".

## Table Usage Guide

The `hackernews_top` table provides insights into the most popular posts within Hacker News. As a data analyst or a content strategist, explore post-specific details through this table, including post title, author, and score. Utilize it to uncover information about trending topics, popular authors, and the overall user engagement in the platform.

## Examples

### Top stories by score
Discover the most popular stories on HackerNews by assessing their score. This allows you to identify trending topics and understand what content resonates most with the community.

```sql
select
  *
from
  hackernews_top
order by
  score desc
```

### Top stories with most comments
Discover the most discussed topics on Hackernews. This query helps you identify the stories that are generating the most conversation, enabling you to stay informed about trending topics.

```sql
select
  *
from
  hackernews_top
order by
  descendants desc
```