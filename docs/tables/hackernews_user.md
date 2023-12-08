---
title: "Steampipe Table: hackernews_user - Query Hacker News Users using SQL"
description: "Allows users to query Hacker News Users, specifically their ID, creation time, karma points, and submitted posts, providing insights into user activity and engagement on the platform."
---

# Table: hackernews_user - Query Hacker News Users using SQL

Hacker News is a social news website focusing on computer science and entrepreneurship. It is run by Y Combinator, an American seed money startup accelerator. Users can submit content, comment on other's posts, and earn karma points as a measure of their contributions to the community.

## Table Usage Guide

The `hackernews_user` table provides insights into users' activities on Hacker News. As a data analyst, explore user-specific details through this table, including their creation time, karma points, and submitted posts. Utilize it to uncover information about users, such as their engagement level, contribution to the platform, and the evolution of their karma points over time.

## Examples

### Get information about a user
Explore the details of a specific user in the HackerNews platform. This can be useful in understanding user behavior or for user profile management.

```sql+postgres
select * from hackernews_user where id = 'pg';
```

```sql+sqlite
select * from hackernews_user where id = 'pg';
```