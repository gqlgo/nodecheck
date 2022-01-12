# nodecheck

[![pkg.go.dev][gopkg-badge]][gopkg]

**nodecheck** will find any GraphQL schema that is not conform to Node interface.

```graphql
# Valid
type User implements Node {
  id: ID!
}

# Invalid. id is not exists
type Community {
  name: String!
}
# Invalid. id is exists but not conform to Node 
type Item {
  id: ID!
}
```

## How to use

A runnable linter can be created with multichecker package.
You can create own linter with your favorite Analyzers.
And nodecheck has independ flag for allow exclude types of nodecheck.

Full example

```go
	flag.StringVar(&excludes, "excludes", "", "exclude GraphQL types for node check. it can specify multiple values separated by `,` and it can use regex(e.g .+Connection")
	flag.Parse()

	analyzer := nodecheck.Analyzer(excludes)

	multichecker.Main(
		analyzer,
	)
```

`lackid` provides a typical main function and you can install with `go install` command.

```sh
$ go install github.com/gqlgo/lackid/cmd/lackid@latest
```

The `lackid` command has two flags, `schema` and `query` which will be parsed and analyzed by lackid's Analyzer.

```sh
$ lackid -schema="server/graphql/schema/**/*.graphql" -query="client/**/*.graphql"
```

The default value of `schema` is "schema/*/**.graphql" and `query` is `query/*/**.graphql`.

`schema` flag accepts URL for a endpoint of GraphQL server.
`lackid` will get schemas by an introspection query via the endpoint.

```sh
$ lackid -schema="https://example.com" -query="client/**/*.graphql"
```

## Author

[![Appify Technologies, Inc.](appify-logo.png)](http://github.com/appify-technologies)

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/gqlgo/lackid
[gopkg-badge]: https://pkg.go.dev/badge/github.com/gqlgo/lackid?status.svg

