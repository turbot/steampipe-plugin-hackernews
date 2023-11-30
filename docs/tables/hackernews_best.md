---
title: "Steampipe Table: hackernews_best - Query Hacker News Best Stories using SQL"
description: "Allows users to query Best Stories from Hacker News, providing insights into the most upvoted and popular stories."
---

# Table: hackernews_best - Query Hacker News Best Stories using SQL

Hacker News is a social news website focusing on computer science and entrepreneurship. It is run by investment fund and startup accelerator, Y Combinator. Users submit stories (known as "posts"), which are voted and commented upon, similar to Reddit.

## Table Usage Guide

The `hackernews_best` table provides insights into the best stories within Hacker News. As a data analyst or a data enthusiast, explore story-specific details through this table, including titles, authors, and scores. Utilize it to uncover information about stories, such as those with the highest upvotes, the most engaged authors, and the most popular topics.

## Examples

### Best stories by score
Discover the top-rated stories on Hacker News by sorting them in descending order based on their score. This can be useful for identifying popular trends and high-interest topics within the community.

```sql
select
  *
from
  hackernews_best
order by
  score desc
```

### Best stories with most comments
Discover the most engaging stories on Hackernews by sorting them according to the number of comments they have received. This can be useful for understanding what topics or story types generate the most discussion among users.

```sql
select
  *
from
  hackernews_best
order by
  descendants desc
```