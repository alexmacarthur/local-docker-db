# Local Docker DB

## What?

A collection of Docker Compose files I've used to quickly spin local databases of various sorts.

## Why?

Because I've oft needed them, particularly when I just don't wanna deal with the hassle of spinning up a DB on my own system.

## How? 

Clone the repo or copy a `docker-compose.yml` file to your system, `cd` into that directory, and turn it on with `docker-compose up`. You may also use a `docker-compose.override.yml` file inside this repository to customize a container. For a full reference on how to use Docker Compose, [go here](https://docs.docker.com/compose/reference/).

## Explore a Container w/ Bash

In any given container, run the following. Most of the time, the service name will just be `db`, but there may be an exception or two.

```
docker-compose exec --user root SERVICE_NAME /bin/bash
```

## Local Persistence

In each setup, a managed volume is created to persist each container's data. This volume can be deleted by passing the `-v` option when deleting the container.

```
docker-compose down -v
```

## Contributions

If you have a Docker Compose configuration for a database not seen here, please consider making a pull request to add it!
