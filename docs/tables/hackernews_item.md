# Table: hackernews_item

Stories, comments, jobs, Ask HNs and even polls are just items. This table
includes the most recent items posted to Hacker News.

`max_items` in the connection configuration defines the number of items
returned by a list query to this table.

## Examples

### List recent items

```sql
select * from hackernews_item
```

### List all recent stories

```sql
select * from hackernews_item where type = 'story'
```

### Recent stories with score > 5

```sql
select
  *
from
  hackernews_item
where
  type = 'story'
  and score > 5
order by
  score desc
```

### Recent stories with no comments

```sql
select
  *
from
  hackernews_item
where
  type = 'story'
  and kids is null
```

### Which users have made more than 5 submissions recently

```sql
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
  count desc
```
