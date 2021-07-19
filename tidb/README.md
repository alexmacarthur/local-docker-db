# TiDB Cluster w/ Docker Compose

with default empty username and password

## Setup

```
git clone --depth 1 https://github.com/pingcap/tidb-docker-compose.git
cd tidb-docker-compose 
docker-compose up
```

## Connecting with client

```
mysql -u root -h 127.0.0.1 -P 4000
```

## Other docs

[tidb docker compose](https://github.com/pingcap/tidb-docker-compose)
