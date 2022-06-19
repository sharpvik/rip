# RIP/TCP Specification

## Request

```
24 Greet {"hello": "world"}
 ^   ^   ^               ^
 |   |   |               |
 |   |   *----- ARG -----*
 |   |
 |   *-- FUNCTION NAME
 |
 *----- CONTENT LENGTH
```

## Response

Functions invoked through RIP _always_ return data in binary format with prior
specification of body length like so:

```
0 31 "Glad to see you there, my boi"
^  ^ ^                             ^
|  | |                             |
|  | *----- RESPONSE --------------*
|  |
|  *----- CONTENT LENGTH
|
*----- STATUS
```

