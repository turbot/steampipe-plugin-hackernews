---
title: "Steampipe Table: hackernews_ask_hn - Query Hacker News Ask HN posts using SQL"
description: "Allows users to query Ask HN posts. It provides data about the posts, including the author, score, and the number of comments."
---

# Table: hackernews_ask_hn - Query Hacker News Ask HN posts using SQL

Hacker News is a social news website focusing on computer science and entrepreneurship. It is run by Paul Graham's investment fund and startup incubator, Y Combinator. In particular, the Ask HN section of the site allows users to ask the Hacker News community questions.

## Table Usage Guide

The `hackernews_ask_hn` table provides insights into Ask HN posts on Hacker News. As a data analyst, explore post-specific details through this table, including author, score, and number of comments. Utilize it to uncover information about the most popular posts, the most active authors, and the discussions that generate the most comments.

## Examples

### Ask HN stories by score
Explore the most popular Ask HN stories on Hackernews, organized by their popularity score. This allows you to quickly identify and engage with the most impactful discussions.

```sql
select
  *
from
  hackernews_ask_hn
order by
  score desc
```

### Ask HN stories with most comments
Discover the Ask HN stories that have garnered the most engagement, allowing you to focus on popular topics and trends within the HackerNews community. This could be particularly useful for content creators, marketers, or researchers looking for highly interactive themes.

```sql
select
  *
from
  hackernews_ask_hn
order by
  descendants desc
```