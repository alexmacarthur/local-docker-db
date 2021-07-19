# ScyllaDB w/ Docker Compose

with default empty username and password

## Username and Password

```
set the enviroment 
TARANTOOL_USER_NAME
TARANTOOL_USER_PASSWORD
```

## Connecting with client

```
node=`docker ps | grep /scylla: | cut -f 1 -d ' '`
docker exec -it $node cqlsh
```

## Other docs

[scylla docker](https://docs.scylladb.com/operating-scylla/manager/1.3/run-in-docker/)
