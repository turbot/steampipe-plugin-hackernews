# Table: hackernews_job

Up to 200 job listings.

## Examples

### Job stories by score

```sql
select
  *
from
  hackernews_job
order by
  score desc
```

### Job stories with most comments

```sql
select
  *
from
  hackernews_job
order by
  descendants desc
```
