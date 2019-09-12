# MariaDB w/ Docker Compose

## Enter the Container w/ Bash

`docker-compose exec --user root db /bin/bash`

# Enter the MySQL Shell

`docker-compose exec --user root db mysql -u root -p`

When prompted, enter `password` as the password.

## Super User Authentication

Username: `root`

Password: `password`

## Create a DB

While inside the shell, run the following:

```
CREATE DATABASE mydatabase;
```
