# NebulaGraph w/ Docker Compose

With default password

## Connecting with client

```
nebula-console -u root -p nebula --address=graphd --port=9669

# or

ip=`ifconfig | grep -Eo 'inet (addr:)?([0-9]*\.){3}[0-9]*' | grep -Eo '([0-9]*\.){3}[0-9]*' | grep -v '127.0.0.1' | head -n 1` 
docker run --rm -it --entrypoint nebula-console vesoft/nebula-console:v2.0.1 -u root -p nebula --address=$ip --port=9669
```


## Other docs

[nebula docker compose](https://docs.nebula-graph.io/2.0.1/2.quick-start/2.deploy-nebula-graph-with-docker-compose/)
