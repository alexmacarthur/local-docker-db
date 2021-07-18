# DyanmoDB w/ Docker Compose

## Starting the Container

Run `docker-compose up db` to turn on the DynamoDB container. After starting, the DB instance will be running at `http://localhost:8000`.

## Accessing the DB w/ the AWS CLI

Run `cp .env-sample .env` to create a new `.env` file, and fill the variables with your AWS credentials.

In order to run `aws` CLI commands against the locally-running DyanmoDB instance, use `docker-compose run aws SOME_AWS_COMMAND`. For example, referencing the DyanmoDB [developer guide](https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/getting-started-step-1.html), you could create a table by running the following:

```bash
docker-compose run aws dynamodb create-table --table-name Music --attribute-definitions AttributeName=Artist,AttributeType=S AttributeName=SongTitle,AttributeType=S --key-schema AttributeName=Artist,KeyType=HASH AttributeName=SongTitle,KeyType=RANGE --provisioned-throughput ReadCapacityUnits=10,WriteCapacityUnits=5
```

And then list all tables by running this:

```bash
docker-compose run aws dynamodb list-tables
```
