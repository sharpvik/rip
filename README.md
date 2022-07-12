# Remote Invocation Protocol

**RIP** is a simple request-response protocol for remote function invocation
designed to be used for microservices and cloud functions.

## REST Has No Chill

REST-ful APIs **suck**. I remember having to write lines upon lines of
boilerplate for URL parsing, JSON parsing, statuses, cookies, etc. Enough is
enough! REST APIs should Rest In Peace.

With **RIP**, you can turn any Go `struct` into an API server in one line:

```go
proto := NewProto(dbConn, redisConn) // struct that serves as a proto source
riptcp.Use(proto).Server().ListenAndServe("localhost:420")
```

## Examples

To see some action, look at the code in [test](./tcp/test/). Run it as follows
and see what happens:

```bash
go test ./tcp/test/ -v
```

## Learn More

You can examine request and response format specification for RIP/TCP
(pronounced as "RIP over TCP") in [this document](./docs/rip-tcp-spec.md).