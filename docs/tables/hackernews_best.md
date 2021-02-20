# Table: hackernews_best

500 of the latest "best" stories.

## Examples

### Best stories by score

```sql
select
  *
from
  hackernews_best
order by
  score desc
```

### Best stories with most comments

```sql
select
  *
from
  hackernews_best
order by
  descendants desc
```
