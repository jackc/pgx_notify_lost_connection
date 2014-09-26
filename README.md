# PGX Notify Lost Connection Test

This file tests pgx's behavior when in a WaitForNotification when the
connection silently goes away.

    DB_HOST=remote.pg.server DB_USER=user DB_PASSWORD=password godep go run main.go

Use godep to test against the vendored version of pgx. Remove it to test a local
copy.

The best way to test behavior by connecting to a nearby machine, then physically
unplugging that remote machine. kill -9 on the PostgreSQL process is
insufficient as the connection is still closed.
