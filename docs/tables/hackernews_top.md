# Table: hackernews_top

500 of the latest "top" stories.

## Examples

### Top stories by score

```sql
select
  *
from
  hackernews_top
order by
  score desc
```

### Top stories with most comments

```sql
select
  *
from
  hackernews_top
order by
  descendants desc
```
