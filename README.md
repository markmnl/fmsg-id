[![Go](https://github.com/markmnl/fmsgid/actions/workflows/go.yml/badge.svg)](https://github.com/markmnl/fmsgid/actions/workflows/go.yml)

# fmsgid

fmsgid is an implementation of the [https://github.com/markmnl/fmsg/blob/main/standards/fmsgid.md](fmsg Id standard) written in go! An fmsg host uses this API to determine if an address exists and associated attributes such as display name and messaging limits.

## Environment

PostgreSQL environment variables must be set for the database to use, refer to: https://www.postgresql.org/docs/current/libpq-envars.html. 

```
GIN_MODE=release
```

## Running

```
./fmsgid
```
