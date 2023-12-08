---
title: "Steampipe Table: hackernews_item - Query Hacker News Items using SQL"
description: "Allows users to query Hacker News Items, specifically to retrieve details about posts, comments, jobs, Ask HNs, and Show HNs, providing insights into the content and metadata of these items."
---

# Table: hackernews_item - Query Hacker News Items using SQL

Hacker News is a social news website focusing on computer science and entrepreneurship. It is run by investment fund and startup incubator, Y Combinator. In essence, it's a social news aggregation site that primarily concentrates on information technology and startups.

## Table Usage Guide

The `hackernews_item` table provides insights into the items posted on Hacker News. As a data analyst or a researcher, explore item-specific details through this table, including the type of post, author, score, and associated metadata. Utilize it to uncover information about the most popular or controversial posts, the frequency of posts by specific authors, or trends in the topics being discussed.

**Important Notes**
- `max_items` in the connection configuration defines the number of items returned by a list query to this table.

## Examples

### List recent items
Explore the latest items in HackerNews to stay updated with the most recent discussions, stories, and comments. This is useful for those who want to keep up with the latest trends and topics in the technology and startup space.

```sql+postgres
select * from hackernews_item;
```

```sql+sqlite
select * from hackernews_item;
```

### List all recent stories
Explore the recent narratives shared on Hackernews to stay updated with the latest discussions and trends. This can be useful to monitor popular topics and engage with the community effectively.

```sql+postgres
select * from hackernews_item where type = 'story';
```

```sql+sqlite
select * from hackernews_item where type = 'story';
```

### Recent stories with score > 5
Discover the segments that consist of popular stories on HackerNews. This can be useful to identify trending topics or high-interest areas, helping to guide your content strategy or research focus.

```sql+postgres
select
  *
from
  hackernews_item
where
  type = 'story'
  and score > 5
order by
  score desc;
```

```sql+sqlite
select
  *
from
  hackernews_item
where
  type = 'story'
  and score > 5
order by
  score desc;
```

### Recent stories with no comments
Explore the recent posts on Hacker News that have not received any comments. This can help identify under-discussed topics and potentially overlooked content.

```sql+postgres
select
  *
from
  hackernews_item
where
  type = 'story'
  and kids is null;
```

```sql+sqlite
select
  *
from
  hackernews_item
where
  type = 'story'
  and kids is null;
```

### Which users have made more than 5 submissions recently
Explore which users have been particularly active by identifying those who have made more than five submissions recently. This can be useful for understanding user engagement and identifying key contributors or influencers within your community.

```sql+postgres
select
  by,
  count(*)
from
  hackernews_item
group by
  by
having
  count(*) > 5
order by
  count desc;
```

```sql+sqlite
select
  by,
  count(*)
from
  hackernews_item
group by
  by
having
  count(*) > 5
order by
  count(*) desc;
```