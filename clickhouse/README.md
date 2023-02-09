# ClickHouse w/ Docker Compose

with default empty password

## Change password and other config

```
volumes:
  - ./local.xml:/etc/clickhouse-server/config.d/local.xml
```

## Connecting with client

```
clickhouse-client
```

## Other docs

[clickhouse server docker](https://hub.docker.com/r/clickhouse/clickhouse-server/)
