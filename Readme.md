[![codecov](https://codecov.io/gh/danielvladco/go-proto-gql/branch/refactor-and-e2e-tests/graph/badge.svg?token=L3N8kUGpGV)](https://codecov.io/gh/danielvladco/go-proto-gql)

THIS REPO FORKED FROM https://github.com/danielvladco/go-proto-gql
Big thanks for him

Protoc plugins for generating graphql schema and go graphql code

If you use micro-service architecture with grpc for back-end and graphql api gateway for front-end, you will find yourself
repeating a lot of code for translating from one transport layer to another (which many times may be a source of bugs)

This repository aims to simplify working with grpc trough protocol buffers and graphql by generating code.

Install:
-

```sh
go install github.com/akhmadnurmuhammad/go-proto-gql/protoc-gen-gql
go install github.com/akhmadnurmuhammad/go-proto-gql/protoc-gen-gogql
```

Usage Examples:
-
The protoc compiler expects to find plugins named `proto-gen-<PLUGIN_NAME>` on the execution `$PATH`. So first:

```sh
export PATH=${PATH}:${GOPATH}/bin
```

---
`--gql_out` plugin will generate graphql files with extension `.graphqls` 
rather than go code which means it can be further used for any other language or framework.

Example: 
```sh
protoc --gql_out=paths=source_relative:. -I=. -I=./example/ ./example/*.proto
```

If you still want to generate go source code instead of graphql then use 
http://github.com/99designs/gqlgen plugin, and map all the generated go types with all the generated graphql types. 
Luckily `--gqlgencfg_out` plugin does exactly this. 

---
`--gogql_out` plugin generates methods for implementing
`github.com/99designs/gqlgen/graphql.Marshaler` and `github.com/99designs/gqlgen/graphql.Unmarshaler` interfaces. Now proto `enum`s, `oneof`s and `map`s will work fine with graphql. 

This plugin also creates convenience methods that will implement generated by the `gqlgen` `MutationResolver` and `QueryResolver` interfaces.

Example:
```sh
protoc --gogql_out=gogoimport=false,paths=source_relative:. -I=. -I=./example/ ./example/*.proto
``` 

---
See `/example` folder for more examples.

## Gateway (alpha)
A unified gateway is also possible. Right now a gateway can be spawn up 
pointing to a list of grpc endpoints (grpc reflection must be enabled on the grpc servers).
The gateway will query the servers for protobuf descriptors and generate a graphql schema abstract tree.
The requests to the gateway will be transformed on the fly to grpc servers without any additional code generation 
or writing any code at all. See `examples/gateway` for usage more info. 

## Community:
Will be very glad for any contributions so feel free to create issues, forks and PRs.

## License:

`go-proto-gql` is released under the Apache 2.0 license. See the LICENSE file for details.
