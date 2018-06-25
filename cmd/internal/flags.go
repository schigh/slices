package internal

import (
	"errors"
	"fmt"
	"regexp"
)

const (
	mapOperation    = "map"
	eachOperation   = "each"
	filterOperation = "filter"
	allOperation    = "all"
)

var re *regexp.Regexp

func init() {
	// https://regex101.com/r/ydgkha/1
	re = regexp.MustCompile(`(?m)(map|each|filter|all)(\(((?P<bv>byvalue)|(?P<cp>copy)|,)+\))?`)
}

// OperationsFromFlags returns all operations in the flag
func OperationsFromFlags(flags string) ([]Operation, error) {
	var ops []Operation
	matches := re.FindAllSubmatch([]byte(flags), -1)
	for _, m := range matches {
		operation := Operation{}
		operation.ByValue = len(m[4]) > 0
		operation.Copy = len(m[5]) > 0
		op := string(m[1])
		switch op {
		case allOperation:
			return allOps(operation.ByValue, operation.Copy), nil
		case mapOperation:
			operation.Template = MapTmpl
		case filterOperation:
			operation.Template = FilterTmpl
		case eachOperation:
			operation.Template = EachTmpl
		default:
			return ops, errors.New(fmt.Sprintf("unknown operation: '%s', op"))
		}

		ops = append(ops, operation)
	}

	return ops, nil
}

func allOps(bv, cp bool) []Operation {
	return []Operation{
		{
			Template: MapTmpl,
			ByValue:  bv,
			Copy:     cp,
		},
		{
			Template: FilterTmpl,
			ByValue:  bv,
			Copy:     cp,
		},
		{
			Template: EachTmpl,
			ByValue:  bv,
			Copy:     cp,
		},
	}
}
