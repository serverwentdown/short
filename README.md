
# short

A tiny, one-dependency URL shortener service. 

# Getting started

Start a cockroachdb server:

```
cockroach start --insecure
echo "CREATE DATABASE IF NOT EXISTS short;" | cockroach sql --insecure
```

Build and run short:

```
go build
POSTGRES=postgresql://root@localhost:26257/short?sslmode=disable
./short
```
