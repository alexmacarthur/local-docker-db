# Microsoft SQL Server w/ Docker Compose

## Super User Authentication

Username: `sa`  
Password: `StrongPassw0rd!`

## Enter the Container w/ Bash

```docker-compose exec --user root db /bin/bash```

## Enter the sqlcmd Shell

```docker-compose exec --user root db /opt/mssql-tools/bin/sqlcmd -S localhost -U sa -P 'StrongPassw0rd!'```

## Create a DB

While inside the sqlcmd shell, run the following:

```sql
CREATE DATABASE MyDatabase;
GO
```

## Other Docs

[Quickstart: Run SQL Server container images with Docker](https://docs.microsoft.com/en-us/sql/linux/quickstart-install-connect-docker)