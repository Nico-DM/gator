# Gator: Blog Aggregator

## Enter toolbox

```bash
toolbox enter
```

## Start database

```bash
pg_ctl -D ~/postgres-data -l ~/postgres.log start
```

## Run migrations

### Up

```bash
cd sql/schema && goose postgres "postgres://nico:@localhost:5432/gator" up && cd ../..
```

### Down

```bash
cd sql/schema && goose postgres "postgres://nico:@localhost:5432/gator" down && cd ../..
```

## Connect to database

```bash
psql "postgres://nico:@localhost:5432/gator"
```

## Usage

```bash
go run . <command> [args...]
```

### Commands

* `register <name>`: registers a user in the database and logs in
* `login <name>`: logs in with an existing user
* `reset`: deletes all data from the database
* `users`: lists all users
* `agg`: fetches an example RSS feed
* `addfeed <name> <url>`: adds an RSS feed associated to the current user
