# Table: hackernews_new

The 500 newest stories.

## Examples


### List recent items

```sql
select 
  * 
from 
  hackernews_new
```

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





### Recent stories with score > 5

```sql
select
  *
from
  hackernews_new
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
  hackernews_new
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
  hackernews_new
group by
  by
having
  count(*) > 5
order by
  count desc
```
