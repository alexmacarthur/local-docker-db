# Redis w/ Docker Compose

Password authentication is achieved through running a custom redis-server command to require the `"password"` password.

## Enter the Container w/ Bash

`docker-compose exec --user root db /bin/bash`

## Enter the Redis Shell

`docker-compose exec --user root db redis-cli`

## Super User Authentication

Password: `password`

While inside the shell, run the following:

```
AUTH "password"
```
