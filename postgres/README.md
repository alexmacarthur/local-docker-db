# Enter the Container w/ Bash

`docker-compose exec --user root db /bin/bash`

# Enter the Postgres Shell

`docker-compose exec --user root db psql -h localhost -U root`

# Create a DB

While inside the shell, run the following:

```
CREATE DATABASE mydatabase;
```
