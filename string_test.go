// generated by github.com/schigh/slices/internal/base.go.  do not edit.
package slices

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
	"testing"

	"github.com/schigh/slices/internal"
)

// region TESTS
// IndexOf

func TestStringSlice_IndexOf(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		needle   string
		expected int
	}{
		{
			name:     "only item",
			slice:    []string{"foo"},
			needle:   "foo",
			expected: 0,
		},
		{
			name:     "at index 1",
			slice:    []string{"foo", "bar", "bar"},
			needle:   "bar",
			expected: 1,
		},
		{
			name:     "missing",
			slice:    []string{"foo", "bar", "fizz", "buzz"},
			needle:   "baz",
			expected: NotInSlice,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := StringSlice(test.slice).IndexOf(test.needle)
			if result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		})
	}
}

// Contains
func TestStringSlice_Contains(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		needle   string
		expected bool
	}{
		{
			name:     "present",
			slice:    []string{"foo", "bar", "fizz", "buzz"},
			needle:   "buzz",
			expected: true,
		},
		{
			name:     "not present",
			slice:    []string{"foo", "bar", "fizz", "buzz"},
			needle:   "baz",
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := StringSlice(test.slice).Contains(test.needle)
			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// SortAsc
func TestStringSlice_SortAsc(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		expected []string
	}{
		{
			name:     "empty",
			slice:    []string{},
			expected: []string{},
		},
		{
			name:     "already sorted",
			slice:    []string{"bar", "baz", "buzz", "fizz", "foo"},
			expected: []string{"bar", "baz", "buzz", "fizz", "foo"},
		},
		{
			name:     "reversed",
			slice:    []string{"foo", "fizz", "buzz", "baz", "bar"},
			expected: []string{"bar", "baz", "buzz", "fizz", "foo"},
		},
		{
			name:     "random",
			slice:    []string{"foo", "bar", "fizz", "buzz", "baz"},
			expected: []string{"bar", "baz", "buzz", "fizz", "foo"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := StringSlice(test.slice).SortAsc().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// SortDesc
func TestStringSlice_SortDesc(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		expected []string
	}{
		{
			name:     "empty",
			slice:    []string{},
			expected: []string{},
		},
		{
			name:     "already sorted",
			slice:    []string{"foo", "fizz", "buzz", "baz", "bar"},
			expected: []string{"foo", "fizz", "buzz", "baz", "bar"},
		},
		{
			name:     "reversed",
			slice:    []string{"bar", "baz", "buzz", "fizz", "foo"},
			expected: []string{"foo", "fizz", "buzz", "baz", "bar"},
		},
		{
			name:     "random",
			slice:    []string{"foo", "bar", "fizz", "buzz", "baz"},
			expected: []string{"foo", "fizz", "buzz", "baz", "bar"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := StringSlice(test.slice).SortDesc().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Unique
func TestStringSlice_Unique(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		expected []string
	}{
		{
			name:     "unaffected",
			slice:    []string{"foo", "bar", "fizz", "buzz", "baz"},
			expected: []string{"foo", "bar", "fizz", "buzz", "baz"},
		},
		{
			name:     "one extra buzz",
			slice:    []string{"foo", "bar", "fizz", "buzz", "baz", "buzz"},
			expected: []string{"foo", "bar", "fizz", "buzz", "baz"},
		},
		{
			name:     "extras everywhere",
			slice:    []string{"foo", "foo", "bar", "foo", "bar", "fizz", "fizz", "fizz", "buzz", "foo", "buzz", "baz", "fizz", "buzz", "baz", "baz", "fizz", "bar", "foo"},
			expected: []string{"foo", "bar", "fizz", "buzz", "baz"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := StringSlice(test.slice).Unique().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Reverse
func TestStringSlice_Reverse(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		expected []string
	}{
		{
			name:     "even length",
			slice:    []string{"foo", "bar", "baz", "fizz", "buzz"},
			expected: []string{"buzz", "fizz", "baz", "bar", "foo"},
		},
		{
			name:     "odd length",
			slice:    []string{"foo", "bar", "baz", "fizz", "buzz", "herp"},
			expected: []string{"herp", "buzz", "fizz", "baz", "bar", "foo"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := StringSlice(test.slice).Reverse().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Filter
func TestStringSlice_Filter(t *testing.T) {
	tests := []struct {
		name       string
		slice      []string
		expected   []string
		filterFunc func(string) bool
	}{
		{
			name:       "starts with f",
			slice:      []string{"foo", "bar", "baz", "fizz", "buzz"},
			expected:   []string{"foo", "fizz"},
			filterFunc: func(s string) bool { return len(s) > 0 && s[0] == 'f' },
		},
		{
			name:     "contains x",
			slice:    []string{"max", "foo", "xray", "homer", "boxy", "blue"},
			expected: []string{"max", "xray", "boxy"},
			filterFunc: func(s string) bool {
				for i := 0; i < len(s); i++ {
					if s[i] == 'x' {
						return true
					}
				}
				return false
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := StringSlice(test.slice).Filter(test.filterFunc).Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Each
func TestStringSlice_Each(t *testing.T) {

	var rabbit string
	tests := []struct {
		name     string
		slice    []string
		expected string
		eachFunc func(string)
	}{
		{
			name:     "concat",
			slice:    []string{"abc", "def", "ghi", "jkl", "mno", "pqr"},
			expected: "abcdefghijklmnopqr",
			eachFunc: func(n string) { rabbit = fmt.Sprintf("%s%s", rabbit, n) },
		},
		{
			name:     "prepend",
			slice:    []string{"aa", "bb", "cc", "dd", "ee"},
			expected: "eeddccbbaaabcdefghijklmnopqr",
			eachFunc: func(n string) { rabbit = fmt.Sprintf("%s%s", n, rabbit) },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			StringSlice(test.slice).Each(test.eachFunc)
			if test.expected != rabbit {
				t.Errorf("expected %s, got %s", test.expected, rabbit)
			}
		})
	}
}

// Map
func TestStringSlice_Map(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		expected []string
		mapFunc  func(string) string
	}{
		{
			name:     "replace vowels",
			slice:    []string{"cats", "dogs", "fish", "birds", "lizards", "mice"},
			expected: []string{"cXts", "dXgs", "fXsh", "bXrds", "lXzXrds", "mXcX"},
			mapFunc: func(n string) string {
				if strings.ContainsAny(n, "aeiou") {
					// hacky!
					n = strings.Replace(
						strings.Replace(
							strings.Replace(
								strings.Replace(
									strings.Replace(n, "a", "X", -1),
									"e", "X", -1,
								),
								"i", "X", -1,
							),
							"o", "X", -1,
						),
						"u", "X", -1,
					)
				}
				return n
			},
		},
		{
			name:     "lowercase",
			slice:    []string{"AbCdE", "aaa", "DDD", "ABx", "12AS"},
			expected: []string{"abcde", "aaa", "ddd", "abx", "12as"},
			mapFunc:  strings.ToLower,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			StringSlice(test.slice).Map(test.mapFunc)
			if !reflect.DeepEqual(test.expected, test.slice) {
				t.Errorf("expected %v, got %v", test.expected, test.slice)
			}
		})
	}
}

// endregion

// region BENCHMARKS
// IndexOf
func BenchmarkStringSlice_IndexOf(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []string
	}{
		{
			name:  "10 elements",
			slice: internal.GenStringSlice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenStringSlice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenStringSlice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenStringSlice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenStringSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				StringSlice(bm.slice).IndexOf("foo")
			}
		})
	}
}

// Contains
func BenchmarkStringSlice_Contains(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []string
	}{
		{
			name:  "10 elements",
			slice: internal.GenStringSlice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenStringSlice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenStringSlice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenStringSlice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenStringSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				StringSlice(bm.slice).Contains("foo")
			}
		})
	}
}

// SortAsc
func BenchmarkStringSlice_SortAsc(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []string
	}{
		{
			name:  "10 elements",
			slice: internal.GenStringSlice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenStringSlice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenStringSlice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenStringSlice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenStringSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				StringSlice(bm.slice).SortAsc()
			}
		})
	}
}

// SortDesc
func BenchmarkStringSlice_SortDesc(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []string
	}{
		{
			name:  "10 elements",
			slice: internal.GenStringSlice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenStringSlice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenStringSlice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenStringSlice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenStringSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				StringSlice(bm.slice).SortDesc()
			}
		})
	}
}

// Reverse
func BenchmarkStringSlice_Reverse(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []string
	}{
		{
			name:  "10 elements",
			slice: internal.GenStringSlice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenStringSlice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenStringSlice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenStringSlice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenStringSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				StringSlice(bm.slice).Reverse()
			}
		})
	}
}

// Filter
func BenchmarkStringSlice_Filter(b *testing.B) {
	benchFunc := func(s string) bool {
		if len(s) == 0 {
			return false
		}
		return s[0] == 'a'
	}
	benchmarks := []struct {
		name  string
		slice []string
	}{
		{
			name:  "10 elements",
			slice: internal.GenStringSlice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenStringSlice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenStringSlice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenStringSlice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenStringSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				StringSlice(bm.slice).Filter(benchFunc)
			}
		})
	}
}

// Each
func BenchmarkStringSlice_Each(b *testing.B) {
	var rabbit string
	benchFunc := func(n string) {
		rabbit = n
	}
	benchmarks := []struct {
		name  string
		slice []string
	}{
		{
			name:  "10 elements",
			slice: internal.GenStringSlice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenStringSlice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenStringSlice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenStringSlice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenStringSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				StringSlice(bm.slice).Each(benchFunc)
			}
		})
	}

	fmt.Fprintf(ioutil.Discard, "%v", rabbit)
}

// Map
func BenchmarkStringSlice_Map(b *testing.B) {
	benchFunc := strings.ToLower
	benchmarks := []struct {
		name  string
		slice []string
	}{
		{
			name:  "10 elements",
			slice: internal.GenStringSlice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenStringSlice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenStringSlice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenStringSlice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenStringSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				StringSlice(bm.slice).Map(benchFunc)
			}
		})
	}
}

// endregion
