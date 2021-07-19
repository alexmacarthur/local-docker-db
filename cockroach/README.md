# CockroachDB w/ Docker Compose

with default empty password

## Connecting with client

```
node=`docker ps | grep cockroach | cut -f 1 -d ' '`
docker exec -it $node cockroach sql --insecure
```

## Other docs

[cockroach docker](https://kb.objectrocket.com/cockroachdb/docker-compose-and-cockroachdb-1151)
