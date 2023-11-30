---
title: "Steampipe Table: hackernews_job - Query Hacker News Jobs using SQL"
description: "Allows users to query Hacker News Jobs, providing insights into job postings and their relevant details."
---

# Table: hackernews_job - Query Hacker News Jobs using SQL

Hacker News is a social news website focusing on computer science and entrepreneurship. It is run by investment fund and startup incubator, Y Combinator. In particular, the 'Jobs' section on Hacker News is a common go-to for tech professionals and startups to post and find job listings.

## Table Usage Guide

The `hackernews_job` table provides insights into job postings on Hacker News. As a job seeker or recruiter, explore job-specific details through this table, including job titles, company names, and URLs. Utilize it to uncover information about the job market in the tech industry, such as trending job roles, popular companies hiring, and the nature of job descriptions.

## Examples

### Job stories by score
Discover the job postings on Hacker News that have received the highest scores. This can be beneficial for understanding which job listings are most popular or garner the most attention.

```sql
select
  *
from
  hackernews_job
order by
  score desc
```

### Job stories with most comments
Explore which job stories on HackerNews are generating the most discussion, helping you to identify popular topics and trends within the tech job market.

```sql
select
  *
from
  hackernews_job
order by
  descendants desc
```