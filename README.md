
# short

A tiny, two-dependency, unauthenticated URL shortener service. 

# Getting started

Start a cockroachdb server:

```
cockroach start --insecure
echo "CREATE DATABASE IF NOT EXISTS short;" | cockroach sql --insecure
```

Build and run short:

```
go build
export POSTGRES=postgresql://root@localhost:26257/short?sslmode=disable
./short
```

# ENV

BASEURL: Base URL used in returning the short address.

PORT: Port to listen on

POSTGRES: [Postgres string](https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters)
