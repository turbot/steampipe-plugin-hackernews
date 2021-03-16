# Table: hackernews_item

Stories, comments, jobs, Ask HNs and even polls are just items. This table
will look up any item posted to Hacker News, however you MUST specify an equal (`=`) qualifier for the `id` in a `where` or `join` clause (`where id =` or `join... on id = `).

## Examples

### Lookup an item by its ID

```sql
select * from hackernews_item where id = '26377203'
```


### List children of an item
```sql
select 
  p.title,
  p.id as parent_id,
  child_id,
  c.id,
  c.text
from 
  hackernews_item as p,
  jsonb_array_elements_text(p.kids) as child_id
left join
  hackernews_item as c on c.id  = child_id
where
  p.id = '26468248'
```

