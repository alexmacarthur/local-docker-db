# Local Docker DB

A collection of Docker Compose files I've used to quickly spin up local databases of various sorts.

# Included Databases

Database    | Docker Compose Configuration    | Website
----------- | ------------------------------- | ----------------------------------
ClickHouse  / [./clickhouse](./clickhouse)    / <https://https://clickhouse.tech/>
DynamoDB    | [./dynamo](./dynamo/)           | <https://aws.amazon.com/dynamodb/>
Fauna       | [./fauna](./fauna/)             | <https://fauna.com/>
MariaDB     | [./maria](./maria/)             | <https://mariadb.org/>
MeiliSearch | [./meilisearch](./meilisearch/) | <https://meilisearch.com/>
MongoDB     | [./mongo](./mongo/)             | <https://www.mongodb.com/>
MySQL       | [./mysql](./mysql/)             | <https://www.mysql.com/>
PostgreSQL  | [./postgres](./postgres/)       | <https://www.postgresql.org/>
Redis       | [./redis](./redis/)             | <https://redis.io/>
ScyllDB     | [./scylla](./scylla)            | <https://www.scylladb.com/>
Tarantool   | [./tarantool](./tarantool/)    | <https://tarantool.io/>
TiDB        | [./tidb](./tidb/)               | <https://pingcap.com/>
YugaByteDB  | [./yugabyte](./yugabyte)        | <https://www.yugabyte.com/>

## Usage

Clone the repo or copy a `docker-compose.yml` file to your system, `cd` into that directory, and turn it on with `docker-compose up` (unless otherwise noted by the directory's `README.md`). You may also use a `docker-compose.override.yml` file inside this repository to customize a container.

For a full reference on how to use Docker Compose, [go here](https://docs.docker.com/compose/reference/).

## Local Persistence

In each setup, a managed volume is created to persist each container's data. This volume can be deleted by passing the `-v` option when deleting the container.

```
docker-compose down -v
```

## Contributions

If you have a Docker Compose configuration for a database not seen here, please consider making a pull request to add it!

## TODO

- add data volume binding for each database
- add all possible environment variables
- add example how to connect with client, with or without docker (have client program installed), and with go
