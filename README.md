# Blacktube GraphQL Backend

![Build](https://github.com/ProjectBlacktube/blacktube-graphql/actions/workflows/BuildImage/badge.svg)

## Requirements

- Go 1.11+
- Docker (optional: for easy development & deployment)
- [Any RDBMS supported by Pop](https://gobuffalo.io/en/docs/db/getting-started#supported-databases) (also optional, you can use Postgres from docker-compose)
- [Soda CLI](https://gobuffalo.io/en/docs/db/toolbox/#installing-cli-support)

## Installation

```
go get github.com/ProjectBlacktube/blacktube-graphql
```

## Development Step

1. Modify GraphQL schema in `graphql/schema.graphql` accordingly.
2. (Optional) Add custom type introduced in schema inside `models`. If you're using simple type, you may skip this step.
3. Trigger stub generation by `go generate ./...`.
4. Tweak and fill `graphql/resolver.go` (somehow there're type mismatch on parser).
5. Implement logic on `manager`.
6. Run `server/server.go` and try it yourself.
