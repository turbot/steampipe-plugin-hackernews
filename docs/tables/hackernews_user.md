# Table: hackernews_user

Query information about users with public activity on Hacker News.  This table
will look up any user on Hacker News, however you MUST specify an equal (`=`) qualifier for the `id` in a `where` or `join` clause (`where id =` or `join... on id = `) 



## Examples

### Get information about a user

```sql
select * from hackernews_user where id = 'pg'
```

### Get new stories with user information

```sql
select 
  i.id as article_id, 
  i.title as article_title, 
  u.id as user_id, 
  u.karma,  
  u.about 
from 
  hackernews_new as i, 
  hackernews_user as u
where 
  i.by = u.id;
```