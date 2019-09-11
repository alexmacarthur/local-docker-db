# FaunaDB w/ Docker Compose

This setup contains two separate services -- one for running the DB platform itself, and one for the shell used to interface with that platform. 

# Using the Dockerized Fauna Shell

Running `docker-compose up` will activate both the DB and Fauna shell services, removing any need to deal with dependencies on your own machine.

## Accessing the Shell

Enter the container to use the shell with the following:

```
docker-compose exec --user root shell /bin/bash
```

## Running Only the DB Service
To run _only_ the DB, use `docker-compose up shell`. In order to use the Fauna shell with that DB, you'll need to set it up on your machine: 

### Set Up the Fauna Shell

Outside your container, you'll need to separately install the `fauna-shell` for interacting with FaunaDB via command line. Dig into the package more [here](https://github.com/fauna/fauna-shell).

### Install the CLI

`npm install -g fauna-shell`

### Configure the Shell
Create a ~/.fauna-shell configuration file. 

```
touch ~/.fauna-shell
```

Place the following in that file. These values will set up the shell to interface with your running FaunaDB container.

```
default=localhost

[localhost]
domain=127.0.0.1
port=8443
scheme=http
secret=secret
```

## Create a DB

Run the following:

```
fauna create-database mydatabase
```

Start up a shell with your newly created database with the following:

```
fauna shell mydatabase
```

## Use a Pretty UI Locally

For a pretty web interface through which to managed your local databases, see [FaunaDB's Developer Dashboard](https://github.com/fauna/dashboard). 
