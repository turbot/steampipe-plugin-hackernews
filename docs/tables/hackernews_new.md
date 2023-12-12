---
title: "Steampipe Table: hackernews_new - Query HackerNews New Stories using SQL"
description: "Allows users to query New Stories on HackerNews, specifically the latest posts, providing insights into the most recent discussions and topics."
---

# Table: hackernews_new - Query HackerNews New Stories using SQL

HackerNews is a social news website focusing on computer science and entrepreneurship. It is run by Y Combinator, a startup accelerator that provides seed funding for startups. HackerNews allows users to submit stories (known as "posts"), which are voted and commented upon, similar to reddit.

## Table Usage Guide

The `hackernews_new` table provides insights into the latest stories posted on HackerNews. As a data analyst or a social media marketer, explore the most recent discussions through this table, including the post's title, author, and score. Utilize it to uncover information about the latest trends, popular topics, and the engagement level of different posts.

## Examples

### New stories by score
Explore the latest stories on Hacker News, sorted by popularity. This query can help you stay updated with trending topics and discussions.

```sql+postgres
select
  *
from
  hackernews_new
order by
  score desc;
```

```sql+sqlite
select
  *
from
  hackernews_new
order by
  score desc;
```

### New stories with most comments
Discover the latest stories that have sparked the most discussion, helping you stay updated on trending topics and popular conversations. This is particularly useful for identifying key interests and engaging discussions within your community.

```sql+postgres
select
  *
from
  hackernews_new
order by
  descendants desc;
```

```sql+sqlite
select
  *
from
  hackernews_new
order by
  descendants desc;
```