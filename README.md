# nodecheck

[![pkg.go.dev][gopkg-badge]][gopkg]

**nodecheck** will find any GraphQL schema that is not conform to Node interface.

```graphql
# Valid
type User implements Node {
  id: ID!
}

# Invalid. not conform to Node
type Community {
  name: String!
}

# Invalid. not conform to Node but id is exist 
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
func main() {
	var excludes string
	flag.StringVar(&excludes, "excludes", "", "exclude GraphQL types for node check. it can specify multiple values separated by `,` and it can use regex(e.g .+Connection")
	flag.Parse()

	analyzer := nodecheck.Analyzer(excludes)

	multichecker.Main(
		analyzer,
	)
}
```

`nodecheck` provides a executable binary. So, you can get cmd tool via `go install` command.

```sh
$ go install github.com/gqlgo/nodecheck/cmd/nodecheck@latest
```

The `nodecheck` command receive two flags, `schema` and `excludes`. `excludes` can specify with regex format and it can receive multiple arguments separated by ','.

```sh
$ nodecheck -schema="server/graphql/schema/**/*.graphql" -excludes=.+Connection,.+Edge
```

## Author

[![Appify Technologies, Inc.](appify-logo.png)](http://github.com/appify-technologies)

<!-- links -->
[gopkg]: https://pkg.go.dev/github.com/gqlgo/nodecheck
[gopkg-badge]: https://pkg.go.dev/badge/github.com/gqlgo/nodecheck?status.svg

