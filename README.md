# Gator: Blog Aggregator

## Enter toolbox

```bash
toolbox enter
```

## Start database

```bash
pg_ctl -D ~/postgres-data -l ~/postgres.log start
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
