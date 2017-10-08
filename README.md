
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

# Usage

```
$ short -h
Usage of ./short:
  -baseurl string
    	baseurl URL of short links (default "localhost:port")
  -num int
    	number of characters in shortened url (default 4)
  -port int
    	listen on port (default 8080)
  -postgres string
    	postgres string (default "postgresql://root@localhost:26257/short?sslmode=disable")
```

See [pq docs](https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters) for more information on the postgres string. 
