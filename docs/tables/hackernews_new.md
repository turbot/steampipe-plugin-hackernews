# Table: hackernews_new

The 500 newest stories.

## Examples

### New stories by score

```sql
select
  *
from
  hackernews_new
order by
  score desc
```

### New stories with most comments

```sql
select
  *
from
  hackernews_new
order by
  descendants desc
```
