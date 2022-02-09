# Remote Function Invocation Protocol

**RFIP** is a simple request-response protocol for remote function invocation
designed to be used for microservices and cloud functions.

## Request

```
54 /url/path/to/specify/function/to/call {"hello":"world"}
 ^ ^                                   ^ ^               ^
 | |                                   | |               |
 | *----- URL -------------------------* *----- ARG -----*
 |
 *----- CONTENT LENGTH
```

## Response

Functions invoked through RFIP _always_ return data in binary format with prior
specification of body length like so:

```
29 Glad to see you there, my boi
 ^ ^                           ^
 | |                           |
 | *----- RESPONSE ------------*
 |
 *----- CONTENT LENGTH
```

## Examples

To see some action, look at code in [server](server/) and [client](client/). Run
them as follows and see what happens:

```bash
go run ./server/
# and then in another terminal window ...
go run ./client/
```
