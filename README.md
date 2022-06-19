# Remote Invocation Protocol

**RIP** is a simple request-response protocol for remote function invocation
designed to be used for microservices and cloud functions.

## REST Has No Chill

REST-ful APIs **suck**. I remember having to write lines upon lines of
boilerplate for URL parsing, JSON parsing, statuses, cookies, etc. Enough is
enough! REST APIs should Rest In Peace.

With **RIP**, you can turn any Go `struct` into an API server in one line:

```go
proto := Proto(dbConn, redisConn) // struct that serves as a proto source
addr := "localhost:420"
rip.Use(proto).Server().ListenAndServeTCP(addr)
```

## Examples

To see some action, look at the code in [server](./example/server/) and
[client](./example/client/). Run them as follows and see what happens:

```bash
go run ./example/server/
# and then in another terminal window ...
go run ./example/client/
```

## Learn More

You can examine request and response format specification for RIP/TCP
(pronounced as "RIP over TCP") in [this document](./docs/tcp-format.md).