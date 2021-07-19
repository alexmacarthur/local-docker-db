# Aerospike w/ Docker Compose

with default empty password

## Connecting with client

```
node=`docker ps | grep aerospike | cut -f 1 -d ' '`
docker exec -it $node aql
```

## Other docs

[aerospike docker](https://docs.aerospike.com/docs/deploy_guides/docker/orchestrate/)
