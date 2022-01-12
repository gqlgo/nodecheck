package main

import (
	"flag"

	"github.com/bannzai/nodecheck"
	"github.com/gqlgo/gqlanalysis/multichecker"
)

var excludes string

func main() {
	flag.StringVar(&excludes, "excludes", "", "exclude GraphQL types for node check. it can specify multiple values separated by `,` and it can use regex(e.g *Connection")
	flag.Parse()

	analyzer := nodecheck.Analyzer(excludes)

	multichecker.Main(
		analyzer,
	)
}
