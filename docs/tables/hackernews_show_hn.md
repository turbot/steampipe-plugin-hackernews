# Table: hackernews_show_hn

Up to 200 of the latest Show HN stories.

## Examples

### Show HN stories by score

```sql
select
  *
from
  hackernews_show_hn
order by
  score desc
```

### Show HN stories with most comments

```sql
select
  *
from
  hackernews_show_hn
order by
  descendants desc
```
