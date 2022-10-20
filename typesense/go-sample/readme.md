# RUNNING `docker-compose`
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
# CONNECT TO TYPESENSE
```bash
curl -f http://localhost:9108/health
```

* the result should be 
```bash
{"ok":true}%
```
# RUNNING TEST CONNECTION WITH GO
```bash
go run main.go
```
* the output should be
```bash
connection successfully true
```