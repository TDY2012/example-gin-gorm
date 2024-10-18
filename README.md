# example-gin-gorm

## Local setup

1. Install VS Code Dev Containers extension.
2. Install Docker.
3. Install Postgres and create one empty database.
4. Prepare `.env` file. (See next section)
5. Open this repository in container via a menu from VS Code Dev Containers extension.

## Environment variables
Put these environment variables in `.env` and place it at the top level directory of this repository.
|Variable Name|Example Value|Note|
|-|-|-|
|DB_HOST|127.0.0.1|Use `host.docker.internal`, if you run Postgres on host.|
|DB_PORT|5432|Postgres uses 5432 by default.|
|DB_USER|postgres||
|DB_PASSWORD|12345678||
|DB_NAME|test-gorm-local||