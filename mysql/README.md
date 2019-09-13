## Enter the Container w/ Bash

`docker-compose exec --user root db /bin/bash`

## Enter the MySQL Shell

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
## Using SequelPro?

You might get a `Authentication plugin 'caching_sha2_password' cannot be loaded` error. If that's the case, [enter the MySQL shell](#enter-the-mysql-shell) and run the following command: 

```mysql
ALTER USER 'root'@'%' IDENTIFIED WITH mysql_native_password BY 'password'
```
