// generated by github.com/schigh/slices/internal/base.go.  do not edit.
package slices

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/schigh/slices/internal"
)

//region TESTS
// IndexOf
func TestInt8Slice_IndexOf(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int8
		needle   int8
		expected int
	}{
		{
			name:     "only item",
			slice:    []int8{1},
			needle:   1,
			expected: 0,
		},
		{
			name:     "at index 1",
			slice:    []int8{0, 1, 1},
			needle:   1,
			expected: 1,
		},
		{
			name:     "missing",
			slice:    []int8{1, 2, 3, 4},
			needle:   5,
			expected: NotInSlice,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Int8Slice(test.slice).IndexOf(test.needle)
			if result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		})
	}
}

// Contains
func TestInt8Slice_Contains(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int8
		needle   int8
		expected bool
	}{
		{
			name:     "present",
			slice:    []int8{1, 2, 3, 4},
			needle:   4,
			expected: true,
		},
		{
			name:     "not present",
			slice:    []int8{1, 2, 3, 4},
			needle:   5,
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Int8Slice(test.slice).Contains(test.needle)
			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// SortAsc
func TestInt8Slice_SortAsc(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int8
		expected []int8
	}{
		{
			name:     "empty",
			slice:    []int8{},
			expected: []int8{},
		},
		{
			name:     "already sorted",
			slice:    []int8{0, 1, 2, 3, 4, 5},
			expected: []int8{0, 1, 2, 3, 4, 5},
		},
		{
			name:     "reversed",
			slice:    []int8{5, 4, 3, 2, 1, 0},
			expected: []int8{0, 1, 2, 3, 4, 5},
		},
		{
			name:     "random",
			slice:    []int8{3, 1, 4, 5, 0, 2},
			expected: []int8{0, 1, 2, 3, 4, 5},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Int8Slice(test.slice).SortAsc().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// SortDesc
func TestInt8Slice_SortDesc(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int8
		expected []int8
	}{
		{
			name:     "empty",
			slice:    []int8{},
			expected: []int8{},
		},
		{
			name:     "already sorted",
			slice:    []int8{5, 4, 3, 2, 1, 0},
			expected: []int8{5, 4, 3, 2, 1, 0},
		},
		{
			name:     "reversed",
			slice:    []int8{0, 1, 2, 3, 4, 5},
			expected: []int8{5, 4, 3, 2, 1, 0},
		},
		{
			name:     "random",
			slice:    []int8{3, 1, 4, 5, 0, 2},
			expected: []int8{5, 4, 3, 2, 1, 0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Int8Slice(test.slice).SortDesc().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Unique
func TestInt8Slice_Unique(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int8
		expected []int8
	}{
		{
			name:     "unaffected",
			slice:    []int8{0, 1, 2, 3, 4, 5},
			expected: []int8{0, 1, 2, 3, 4, 5},
		},
		{
			name:     "one extra five",
			slice:    []int8{5, 0, 1, 2, 3, 4, 5},
			expected: []int8{5, 0, 1, 2, 3, 4},
		},
		{
			name:     "extras everywhere",
			slice:    []int8{0, 0, 1, 0, 1, 2, 2, 2, 3, 0, 3, 4, 2, 3, 4, 4, 2, 1, 0},
			expected: []int8{0, 1, 2, 3, 4},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Int8Slice(test.slice).Unique().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Reverse
func TestInt8Slice_Reverse(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int8
		expected []int8
	}{
		{
			name:     "even length",
			slice:    []int8{0, 1, 2, 3, 4, 5},
			expected: []int8{5, 4, 3, 2, 1, 0},
		},
		{
			name:     "odd length",
			slice:    []int8{0, 1, 2, 3, 4, 5, 6},
			expected: []int8{6, 5, 4, 3, 2, 1, 0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Int8Slice(test.slice).Reverse().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Filter
func TestInt8Slice_Filter(t *testing.T) {
	tests := []struct {
		name       string
		slice      []int8
		expected   []int8
		filterFunc func(int8) bool
	}{
		{
			name:       "gt 10",
			slice:      []int8{1, 2, 5, 11, 13, 15},
			expected:   []int8{11, 13, 15},
			filterFunc: func(n int8) bool { return n > 10 },
		},
		{
			name:       "mod 3",
			slice:      []int8{1, 2, 6, 11, 12, 15, 17},
			expected:   []int8{6, 12, 15},
			filterFunc: func(n int8) bool { return n%3 == 0 },
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Int8Slice(test.slice).Filter(test.filterFunc).Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Each
func TestInt8Slice_Each(t *testing.T) {

	var rabbit int8
	tests := []struct {
		name     string
		slice    []int8
		expected int8
		eachFunc func(int8)
	}{
		{
			name:     "add n",
			slice:    []int8{1, 2, 5, 11, 13, 15},
			expected: 47,
			eachFunc: func(n int8) { rabbit += n },
		},
		{
			name:     "subtract n",
			slice:    []int8{1, 2, 6, 8, 12},
			expected: 18,
			eachFunc: func(n int8) { rabbit -= n },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			Int8Slice(test.slice).Each(test.eachFunc)
			if test.expected != rabbit {
				t.Errorf("expected %v, got %v", test.expected, rabbit)
			}
		})
	}
}

// Map
func TestInt8Slice_Map(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int8
		expected []int8
		mapFunc  func(int8) int8
	}{
		{
			name:     "add 3",
			slice:    []int8{1, 2, 5, 11, 13, 15},
			expected: []int8{4, 5, 8, 14, 16, 18},
			mapFunc:  func(n int8) int8 { return n + 3 },
		},
		{
			name:     "set mod 2",
			slice:    []int8{1, 2, 6, 8, 12, 15, 17},
			expected: []int8{1, 0, 0, 0, 0, 1, 1},
			mapFunc:  func(n int8) int8 { return int8(n % 2) },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			Int8Slice(test.slice).Map(test.mapFunc)
			if !reflect.DeepEqual(test.expected, test.slice) {
				t.Errorf("expected %v, got %v", test.expected, test.slice)
			}
		})
	}
}

// endregion

// region BENCHMARKS
// IndexOf
func BenchmarkInt8Slice_IndexOf(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []int8
	}{
		{
			name:  "10 elements",
			slice: internal.GenInt8Slice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenInt8Slice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenInt8Slice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenInt8Slice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenInt8Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int8Slice(bm.slice).IndexOf(1)
			}
		})
	}
}

// Contains
func BenchmarkInt8Slice_Contains(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []int8
	}{
		{
			name:  "10 elements",
			slice: internal.GenInt8Slice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenInt8Slice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenInt8Slice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenInt8Slice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenInt8Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int8Slice(bm.slice).Contains(1)
			}
		})
	}
}

// SortAsc
func BenchmarkInt8Slice_SortAsc(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []int8
	}{
		{
			name:  "10 elements",
			slice: internal.GenInt8Slice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenInt8Slice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenInt8Slice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenInt8Slice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenInt8Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int8Slice(bm.slice).SortAsc()
			}
		})
	}
}

// SortDesc
func BenchmarkInt8Slice_SortDesc(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []int8
	}{
		{
			name:  "10 elements",
			slice: internal.GenInt8Slice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenInt8Slice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenInt8Slice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenInt8Slice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenInt8Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int8Slice(bm.slice).SortDesc()
			}
		})
	}
}

// Reverse
func BenchmarkInt8Slice_Reverse(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []int8
	}{
		{
			name:  "10 elements",
			slice: internal.GenInt8Slice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenInt8Slice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenInt8Slice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenInt8Slice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenInt8Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int8Slice(bm.slice).Reverse()
			}
		})
	}
}

// Filter
func BenchmarkInt8Slice_Filter(b *testing.B) {
	benchFunc := func(n int8) bool {
		return n%2 == 0
	}
	benchmarks := []struct {
		name  string
		slice []int8
	}{
		{
			name:  "10 elements",
			slice: internal.GenInt8Slice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenInt8Slice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenInt8Slice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenInt8Slice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenInt8Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int8Slice(bm.slice).Filter(benchFunc)
			}
		})
	}
}

// Each
func BenchmarkInt8Slice_Each(b *testing.B) {
	var rabbit int8
	benchFunc := func(n int8) {
		rabbit = n
	}
	benchmarks := []struct {
		name  string
		slice []int8
	}{
		{
			name:  "10 elements",
			slice: internal.GenInt8Slice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenInt8Slice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenInt8Slice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenInt8Slice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenInt8Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int8Slice(bm.slice).Each(benchFunc)
			}
		})
	}

	fmt.Fprintf(ioutil.Discard, "%v", rabbit)
}

// Map
func BenchmarkInt8Slice_Map(b *testing.B) {
	benchFunc := func(n int8) int8 {
		n++
		return n
	}
	benchmarks := []struct {
		name  string
		slice []int8
	}{
		{
			name:  "10 elements",
			slice: internal.GenInt8Slice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenInt8Slice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenInt8Slice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenInt8Slice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenInt8Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int8Slice(bm.slice).Map(benchFunc)
			}
		})
	}
}

// endregion
