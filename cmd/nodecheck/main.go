package main

import (
	"github.com/bannzai/nodecheck"
	"github.com/gqlgo/gqlanalysis/multichecker"
)

func main() {
	multichecker.Main(
		nodecheck.Analyzer,
	)
}
