package nodecheck

import (
	"regexp"
	"strings"

	"github.com/vektah/gqlparser/v2/ast"

	"github.com/gqlgo/gqlanalysis"
)

func Analyzer(excludes string) *gqlanalysis.Analyzer {
	return &gqlanalysis.Analyzer{
		Name: "nodecheck",
		Doc:  "nodecheck will find any GraphQL schema that is not conform to Node interface",
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

		for _, t := range needToNodeTypes {
			// Skip private type. e.g) __Directive, __Enum ...
			if strings.HasPrefix(t.Name, "__") {
				break
			}
			pass.Reportf(t.Position, "%+v should conform to Node", t.Name)
		}

		return nil, nil
	}
}
