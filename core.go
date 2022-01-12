package nodecheck

import (
	"flag"
	"fmt"
	"regexp"
	"strings"

	"github.com/vektah/gqlparser/v2/ast"

	"github.com/gqlgo/gqlanalysis"
)

var excludes string
var Analyzer = &gqlanalysis.Analyzer{
	Name: "nodecheck",
	Doc:  "nodecheck finds invalid GraphQL type that type does not conform Node interface",
	Flags: func() flag.FlagSet {
		f := flag.NewFlagSet("node check", flag.ExitOnError)
		f.StringVar(&excludes, "exclude", "", "exclude GraphQL types for node check. it can specify multiple values separated by `,` and it can use regex(e.g *Connection")
		return *f
	}(),
	Run: run,
}

func run(pass *gqlanalysis.Pass) (interface{}, error) {
	allTypes := map[string]*ast.Definition{}
	for _, def := range pass.Schema.Types {
		if def.Kind == ast.Object {
			allTypes[def.Name] = def
		}
	}

	allNodeImplements := map[string]*ast.Definition{}
	for name, t := range allTypes {
		for _, typeInterface := range pass.Schema.Implements[name] {
			if typeInterface.Kind == ast.Interface && typeInterface.Name == "Node" {
				allNodeImplements[name] = t
			}
		}
	}

	unconformedTypes := map[string]*ast.Definition{}
	for k, v := range allTypes {
		if _, ok := allNodeImplements[k]; !ok {
			unconformedTypes[v.Name] = v
		}
	}

	needToNodeTypes := []*ast.Definition{}

	for k, v := range unconformedTypes {
		ok := false
		for _, rules := range strings.Split(excludes, ",") {
			regex := regexp.MustCompile(rules)

			if ok {
				break
			}
			if regex.MatchString(k) {
				ok = true
			}
		}

		if !ok {
			needToNodeTypes = append(needToNodeTypes, v)
		}
	}

	if len(needToNodeTypes) > 0 {
		return nil, fmt.Errorf("GraphQL types need to conform to Node type %s", needToNodeTypes)
	}

	return nil, nil
}
