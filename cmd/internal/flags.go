package internal

import (
	"fmt"
	"regexp"
)

const (
	mapOperation     = "map"
	eachOperation    = "each"
	tryEachOperation = "tryeach"
	ifEachOperation  = "ifeach"
	filterOperation  = "filter"
	allOperation     = "all"
)

var re *regexp.Regexp

func init() {
	// https://regex101.com/r/ydgkha/1
	re = regexp.MustCompile(`(?m)(map|ifeach|tryeach|each|filter|all)(\(((?P<bv>byvalue)|(?P<cp>copy)|,)+\))?`)
}

// OperationsFromFlags returns all operations in the flag
func OperationsFromFlags(flags string) ([]Operation, error) {
	var ops []Operation
	matches := re.FindAllSubmatch([]byte(flags), -1)
	for _, m := range matches {
		operation := Operation{}
		operation.ByRef = len(m[4]) <= 0
		operation.Copy = len(m[5]) > 0
		op := string(m[1])
		switch op {
		case allOperation:
			return allOps(operation.ByRef, operation.Copy), nil
		case mapOperation:
			operation.Template = MapTmpl
			operation.Name = "map"
		case filterOperation:
			operation.Template = FilterTmpl
			operation.Name = "filter"
		case eachOperation:
			operation.Template = EachTmpl
			operation.Name = "each"
		case tryEachOperation:
			operation.Template = TryEachTmpl
			operation.Name = "tryeach"
		case ifEachOperation:
			operation.Template = IfEachTmpl
			operation.Name = "ifeach"
		default:
			return ops, fmt.Errorf("unknown operation: '%s'", op)
		}

		ops = append(ops, operation)
	}

	return ops, nil
}

func allOps(br, cp bool) []Operation {
	return []Operation{
		{
			Template: MapTmpl,
			ByRef:    br,
			Copy:     cp,
			Name:     "map",
		},
		{
			Template: FilterTmpl,
			ByRef:    br,
			Copy:     cp,
			Name:     "filter",
		},
		{
			Template: EachTmpl,
			ByRef:    br,
			Copy:     cp,
			Name:     "each",
		},
		{
			Template: TryEachTmpl,
			ByRef:    br,
			Copy:     cp,
			Name:     "tryeach",
		},
		{
			Template: IfEachTmpl,
			ByRef:    br,
			Copy:     cp,
			Name:     "ifeach",
		},
	}
}
