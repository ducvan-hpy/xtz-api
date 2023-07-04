# Introduction

The *xtz-api* is an API to retrieve XTZ's delegations from [TzKT
API](https://api.tzkt.io/v1/operations/delegations).

# Workflow

When starting the *xtz-api*, it will poll the delegations from the
TzKT API using a *poller*.

## Poller

The poller is used to periodically retrieve the delegations to update
the application's database (in memory in our current implementation).

It will poll every seconds if there are new updates.
Otherwise, it will poll every 2 minutes.

It needs around a minute to retrieve delegations from 2018 to 2023 as
we are limited to 10000 returned delegations by the TzKT API.

## API

The *xtz-api* API's schema is defined in _cmd/xtz-api/xtz_oapi.yml_.

It exposes a single endpoint: */xtz/delegations* which can handle a
*year* optional parameter.

# Software Architecture

The implementation of this project tried to follow a domain driven
design.

The goal is to keep the code clean and make it easy to test (even if I
did not have the time to add unit tests).

The logic of the application is implemented in the *domain* layer.

The *application* layer is meant to interface the outer layers to the
domain.

The *infrastructure* layer contains the package to communicate with
the outer environment.

# How to start

To start the application, you must set GOPATH variable in your
environment (_.bashrc_, _.zshrc_).
```
export GOPATH=`go env GOPATH`
```

Then you can call:
```
> `make run`
```

To compile the binary, you will need to install *oapi-codegen*
(v1.13.0) binary. It helps us to generate the API layer from our
OpenAPI schema.

```
> oapi-codegen -version
github.com/deepmap/oapi-codegen/cmd/oapi-codegen
v1.13.0
```