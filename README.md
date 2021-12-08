# link
Yet a simple link redirect app.


## Make links.id case-sensitive
```sql
alter table links change id id varchar(255) character set utf8mb4 collate utf8mb4_bin;
```