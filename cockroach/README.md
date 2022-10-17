# CockroachDB w/ Docker Compose

with default empty password

## Connecting with client

```
node=`docker ps | grep cockroach | cut -f 1 -d ' '`
docker exec -it $node cockroach sql --insecure
```

## Connection string

```bash
postgresql://root@localhost:26257/defaultdb?sslmode=disable
```

## Other docs

[cockroach docker](https://kb.objectrocket.com/cockroachdb/docker-compose-and-cockroachdb-1151)
