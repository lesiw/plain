# lesiw.io/plain

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

Use the [libpq environment variables][envar] to configure your database
connection. These are automatically understood by `pgx`.

## Usage

Add this to the top of your application's `main.go`:

``` go
//go:generate go run lesiw.io/plain/cmd/plaingen@latest
```

Run `go generate`.

Create a new context: `ctx := context.Background()`.

Use `plain.ConnectPgx(ctx, migrate)` to establish a connection to the database.
`migrate` is an auto-generated function. `ConnectPgx` returns a `pgxpool.Pool`
to use in the rest of your application.

Import the statements generated in `internal/stmt` as needed to use the embedded
sql files.

## Example

See [lesiw.io/smol][smol].

[migrate]: https://github.com/golang-migrate/migrate?tab=readme-ov-file#migration-files
[envars]: https://www.postgresql.org/docs/current/libpq-envars.html
[smol]: https://lesiw.io/smol
[pgx]: https://github.com/jackc/pgx
