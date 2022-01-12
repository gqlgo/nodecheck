package nodecheck

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/vektah/gqlparser/v2/ast"

	"github.com/gqlgo/gqlanalysis"
)

func Analyzer(excludes string) *gqlanalysis.Analyzer {
	return &gqlanalysis.Analyzer{
		Name: "nodecheck",
		Doc:  "nodecheck finds invalid GraphQL type that type does not conform Node interface",
		Run:  run(excludes),
	}
}

func run(excludes string) func(pass *gqlanalysis.Pass) (interface{}, error) {
	return func(pass *gqlanalysis.Pass) (interface{}, error) {
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
			for _, rule := range strings.Split(excludes, ",") {
				if len(rule) > 0 {
					regex := regexp.MustCompile(rule)

					if ok {
						break
					}
					if regex.MatchString(k) {
						ok = true
					}
				}
			}

			if !ok {
				needToNodeTypes = append(needToNodeTypes, v)
			}
		}

		if len(needToNodeTypes) > 0 {
			names := make([]string, len(needToNodeTypes))
			for i, t := range needToNodeTypes {
				names[i] = t.Name
			}

			out := strings.Join(names, "\n")
			return nil, fmt.Errorf("GraphQL types need to conform to Node type %s", out)
		}

		return nil, nil
	}
}
