# RedPanda w/ Docker Compose

with default no password

## Connecting with client

```
node=`docker ps | grep redpanda | cut -f 1 -d ' '`
docker exec -it redpanda-1 rpk
```

## Other docs

[redpanda docker](https://vectorized.io/docs/quick-start-docker)
