# lesiw.io/plain

[![Go Reference](https://pkg.go.dev/badge/lesiw.io/plain.svg)](https://pkg.go.dev/lesiw.io/plain)

A micro-framework for building applications with plain, simple, SQL.

<p align="center">
    <img src="/../media/simple.jpg" />
</p>

Currently, only [pgx/v5][pgx] (PostgreSQL) is supported.

## Structure

Statements go in `sql/statements`. One statement per file. Only files ending in
`.sql` will be used.

Migrations go in `sql/migrations`. Refer to the [golang-migrate
documentation][migrate].

## Configuration

Use the [libpq environment variables][envars] to configure your database
connection. These are automatically understood by `pgx`.

## Usage

Add this to the top of your application's `main.go`:

``` go
//go:generate go run lesiw.io/plain/cmd/plaingen@latest
```

Run `go generate`.

Use `plain.ConnectPgx(context.Context) *pgxPool.pool` to establish a connection
to the database.

Import `internal/stmt` to access the SQL statements written in `sql/statements`.

## Example

See [lesiw.io/smol][smol].

[migrate]: https://github.com/golang-migrate/migrate?tab=readme-ov-file#migration-files
[envars]: https://www.postgresql.org/docs/current/libpq-envars.html
[smol]: https://lesiw.io/smol
[pgx]: https://github.com/jackc/pgx
