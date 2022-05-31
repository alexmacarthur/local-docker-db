# ScyllaDB w/ Docker Compose

with default empty username and password

## Connecting with client

```
node=`docker ps | grep /scylla: | cut -f 1 -d ' '`
docker exec -it $node cqlsh
```

## Other docs

[scylla docker](https://docs.scylladb.com/operating-scylla/manager/1.3/run-in-docker/)
