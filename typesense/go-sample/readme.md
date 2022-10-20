# RUNNING
```bash
docker compose up -d
```
# SHUTTING DOWN
```bash
docker compose down -v
```
# ENV
```bash
set searchkey=<up-to-you>
```
# CONNECT
```bash
curl -f http://localhost:9108/health
```

* the result should be 
```bash
{"ok":true}%
```