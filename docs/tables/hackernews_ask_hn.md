# Table: hackernews_ask_hn

Up to 200 of the latest Ask HN stories.

## Examples

### Ask HN stories by score

```sql
select
  *
from
  hackernews_ask_hn
order by
  score desc
```

### Ask HN stories with most comments

```sql
select
  *
from
  hackernews_ask_hn
order by
  descendants desc
```
