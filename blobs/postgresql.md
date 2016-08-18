# PostgreSQL #

## Tunning ##
https://wiki.postgresql.org/wiki/Tuning_Your_PostgreSQL_Server

## Konfigurations Parameter ##
https://www.postgresql.org/docs/current/static/runtime-config.html

## Error codes ##
https://www.postgresql.org/docs/9.5/static/errcodes-appendix.html

## Anzeigen von offenen Datenbankverbindungen ##
```sql
SELECT * FROM pg_stat_activity;
```
https://www.postgresql.org/message-id/200404260720.i3Q7KaxE030841%40lurza.secnetix.de

## Kill aktive Verbindungen ##
```sql
SELECT pg_terminate_backend(pid int)
```

Beende alle Verbindungen bis auf die eigene
```sql
SELECT pg_terminate_backend(pg_stat_activity.pid)
FROM pg_stat_activity
WHERE pg_stat_activity.datname = 'TARGET_DB'
  AND pid <> pg_backend_pid();
```

## Verzeichnisstruktur mit ltree ##
* http://www.postgresonline.com/journal/archives/173-Using-LTree-to-Represent-and-Query-Hierarchy-and-Tree-Structures.html
* https://www.postgresql.org/docs/current/static/ltree.html
