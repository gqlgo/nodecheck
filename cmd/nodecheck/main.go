package main

import (
	"flag"

	"github.com/gqlgo/gqlanalysis/multichecker"
	"github.com/gqlgo/nodecheck"
)

func main() {
	var excludes string
	flag.StringVar(&excludes, "excludes", "", "exclude GraphQL types for node check. it can specify multiple values separated by `,` and it can use regex(e.g .+Connection")
	flag.Parse()

	analyzer := nodecheck.Analyzer(excludes)

	multichecker.Main(
		analyzer,
	)
}
