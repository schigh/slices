package slices

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/schigh/slices/internal"
)

/* generated by github.com/schigh/slices/base/base.go.  do not edit. */

//region TESTS
// IndexOf
func TestIntSlice_IndexOf(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		needle   int
		expected int
	}{
		{
			name:     "only item",
			slice:    []int{1},
			needle:   1,
			expected: 0,
		},
		{
			name:     "at index 1",
			slice:    []int{0, 1, 1},
			needle:   1,
			expected: 1,
		},
		{
			name:     "missing",
			slice:    []int{1, 2, 3, 4},
			needle:   5,
			expected: NotInSlice,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IntSlice(test.slice).IndexOf(test.needle)
			if result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		})
	}
}

// Contains
func TestIntSlice_Contains(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		needle   int
		expected bool
	}{
		{
			name:     "present",
			slice:    []int{1, 2, 3, 4},
			needle:   4,
			expected: true,
		},
		{
			name:     "not present",
			slice:    []int{1, 2, 3, 4},
			needle:   5,
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IntSlice(test.slice).Contains(test.needle)
			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// SortAsc
func TestIntSlice_SortAsc(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		expected []int
	}{
		{
			name:     "empty",
			slice:    []int{},
			expected: []int{},
		},
		{
			name:     "already sorted",
			slice:    []int{0, 1, 2, 3, 4, 5},
			expected: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name:     "reversed",
			slice:    []int{5, 4, 3, 2, 1, 0},
			expected: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name:     "random",
			slice:    []int{3, 1, 4, 5, 0, 2},
			expected: []int{0, 1, 2, 3, 4, 5},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IntSlice(test.slice).SortAsc().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// SortDesc
func TestIntSlice_SortDesc(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		expected []int
	}{
		{
			name:     "empty",
			slice:    []int{},
			expected: []int{},
		},
		{
			name:     "already sorted",
			slice:    []int{5, 4, 3, 2, 1, 0},
			expected: []int{5, 4, 3, 2, 1, 0},
		},
		{
			name:     "reversed",
			slice:    []int{0, 1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1, 0},
		},
		{
			name:     "random",
			slice:    []int{3, 1, 4, 5, 0, 2},
			expected: []int{5, 4, 3, 2, 1, 0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IntSlice(test.slice).SortDesc().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Unique
func TestIntSlice_Unique(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		expected []int
	}{
		{
			name:     "unaffected",
			slice:    []int{0, 1, 2, 3, 4, 5},
			expected: []int{0, 1, 2, 3, 4, 5},
		},
		{
			name:     "one extra five",
			slice:    []int{5, 0, 1, 2, 3, 4, 5},
			expected: []int{5, 0, 1, 2, 3, 4},
		},
		{
			name:     "extras everywhere",
			slice:    []int{0, 0, 1, 0, 1, 2, 2, 2, 3, 0, 3, 4, 2, 3, 4, 4, 2, 1, 0},
			expected: []int{0, 1, 2, 3, 4},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IntSlice(test.slice).Unique().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Reverse
func TestIntSlice_Reverse(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		expected []int
	}{
		{
			name:     "even length",
			slice:    []int{0, 1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1, 0},
		},
		{
			name:     "odd length",
			slice:    []int{0, 1, 2, 3, 4, 5, 6},
			expected: []int{6, 5, 4, 3, 2, 1, 0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IntSlice(test.slice).Reverse().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Filter
func TestIntSlice_Filter(t *testing.T) {
	tests := []struct {
		name       string
		slice      []int
		expected   []int
		filterFunc func(int) bool
	}{
		{
			name:       "gt 10",
			slice:      []int{1, 2, 5, 11, 13, 15},
			expected:   []int{11, 13, 15},
			filterFunc: func(n int) bool { return n > 10 },
		},
		{
			name:       "mod 3",
			slice:      []int{1, 2, 6, 11, 12, 15, 17},
			expected:   []int{6, 12, 15},
			filterFunc: func(n int) bool { return n%3 == 0 },
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := IntSlice(test.slice).Filter(test.filterFunc).Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Each
func TestIntSlice_Each(t *testing.T) {

	var rabbit int
	tests := []struct {
		name     string
		slice    []int
		expected int
		eachFunc func(int)
	}{
		{
			name:     "add n",
			slice:    []int{1, 2, 5, 11, 13, 15},
			expected: 47,
			eachFunc: func(n int) { rabbit += n },
		},
		{
			name:     "subtract n",
			slice:    []int{1, 2, 6, 8, 12},
			expected: 18,
			eachFunc: func(n int) { rabbit -= n },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			IntSlice(test.slice).Each(test.eachFunc)
			if test.expected != rabbit {
				t.Errorf("expected %v, got %v", test.expected, rabbit)
			}
		})
	}
}

// TryEach
func TestIntSlice_TryEach(t *testing.T) {

	var rabbit int
	myErr := errors.New("i am an error")
	tests := []struct {
		name      string
		slice     []int
		expected  int
		expected2 error
		before    func()
		eachFunc  func(int) error
	}{
		{
			name:      "add n",
			slice:     []int{1, 2, 5, 11, 13, 15},
			expected:  NotInSlice,
			expected2: nil,
			eachFunc: func(n int) error {
				rabbit += n
				return nil
			},
		},
		{
			name:      "subtract n",
			slice:     []int{1, 2, 6, 8, 12},
			expected:  NotInSlice,
			expected2: nil,
			eachFunc: func(n int) error {
				rabbit -= n
				return nil
			},
		},
		{
			name:      "errors",
			slice:     []int{1, 2, 5, 11, 13, 15},
			expected:  3,
			expected2: myErr,
			before:    func() { rabbit = 0 },
			eachFunc: func(n int) error {
				if n > 5 {
					return myErr
				}
				rabbit += n
				return nil
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.before != nil {
				test.before()
			}
			e, i := IntSlice(test.slice).TryEach(test.eachFunc)
			if test.expected != e {
				t.Errorf("expected %v, got %v", test.expected, e)
			}
			if test.expected2 != i {
				t.Errorf("expected %v, got %v", test.expected2, i)
			}
		})
	}
}

// IfEach
func TestIntSlice_IfEach(t *testing.T) {

	var rabbit int
	tests := []struct {
		name      string
		slice     []int
		expected  int
		expected2 bool
		before    func()
		err       error
		eachFunc  func(int) bool
	}{
		{
			name:      "all return true",
			slice:     []int{1, 2, 5, 11, 13, 15},
			expected:  NotInSlice,
			expected2: true,
			eachFunc: func(n int) bool {
				rabbit += n
				return true
			},
		},
		{
			name:      "subtract n",
			slice:     []int{1, 2, 6, 8, 12},
			expected:  NotInSlice,
			expected2: true,
			eachFunc: func(n int) bool {
				rabbit -= n
				return true
			},
		},
		{
			name:      "breaking",
			slice:     []int{1, 2, 5, 11, 13, 15},
			expected:  3,
			expected2: false,
			before:    func() { rabbit = 0 },
			eachFunc: func(n int) bool {
				if n > 5 {
					return false
				}
				rabbit += n
				return true
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.before != nil {
				test.before()
			}
			e, i := IntSlice(test.slice).IfEach(test.eachFunc)
			if test.expected != e {
				t.Errorf("expected %v, got %v", test.expected, e)
			}
			if test.expected2 != i {
				t.Errorf("expected %v, got %v", test.expected2, i)
			}
		})
	}
}

// Map
func TestIntSlice_Map(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		expected []int
		mapFunc  func(int) int
	}{
		{
			name:     "add 3",
			slice:    []int{1, 2, 5, 11, 13, 15},
			expected: []int{4, 5, 8, 14, 16, 18},
			mapFunc:  func(n int) int { return n + 3 },
		},
		{
			name:     "set mod 2",
			slice:    []int{1, 2, 6, 8, 12, 15, 17},
			expected: []int{1, 0, 0, 0, 0, 1, 1},
			mapFunc:  func(n int) int { return int(n % 2) },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			IntSlice(test.slice).Map(test.mapFunc)
			if !reflect.DeepEqual(test.expected, test.slice) {
				t.Errorf("expected %v, got %v", test.expected, test.slice)
			}
		})
	}
}

// Chunk
func TestIntSlice_Chunk(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		slice    []int
		expected [][]int
	}{
		{
			name:     "chunks of 2 no remainder",
			size:     2,
			slice:    []int{1, 2, 5, 11, 13, 15},
			expected: [][]int{{1, 2}, {5, 11}, {13, 15}},
		},
		{
			name:     "chunks of 2 with remainder",
			size:     2,
			slice:    []int{1, 2, 5, 11, 13, 15, 17},
			expected: [][]int{{1, 2}, {5, 11}, {13, 15}, {17}},
		},
		{
			name:     "chunks of 100",
			size:     100,
			slice:    []int{1, 2, 5, 11, 13, 15, 17},
			expected: [][]int{{1, 2, 5, 11, 13, 15, 17}},
		},
		{
			name:     "chunks of 4",
			size:     4,
			slice:    []int{1, 2, 5, 11, 13, 15, 17},
			expected: [][]int{{1, 2, 5, 11}, {13, 15, 17}},
		},
		{
			name:     "chunks of 5",
			size:     5,
			slice:    []int{1, 2, 5, 11, 13, 15, 17},
			expected: [][]int{{1, 2, 5, 11, 13}, {15, 17}},
		},
		{
			name:     "empty slice",
			size:     5,
			slice:    []int{},
			expected: [][]int{},
		},
		{
			name:     "invalid chunk size",
			size:     -1,
			slice:    []int{1, 2, 5, 11, 13, 15, 17},
			expected: [][]int{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := IntSlice(test.slice).Chunk(test.size)
			if !reflect.DeepEqual(test.expected, out) {
				t.Errorf("expected %v, got %v", test.expected, test.slice)
			}
		})
	}
}

// endregion

// region BENCHMARKS
// IndexOf
func BenchmarkIntSlice_IndexOf(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []int
	}{
		{
			name:  "10 elements",
			slice: internal.GenIntSlice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenIntSlice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenIntSlice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenIntSlice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenIntSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IntSlice(bm.slice).IndexOf(1)
			}
		})
	}
}

// Contains
func BenchmarkIntSlice_Contains(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []int
	}{
		{
			name:  "10 elements",
			slice: internal.GenIntSlice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenIntSlice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenIntSlice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenIntSlice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenIntSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IntSlice(bm.slice).Contains(1)
			}
		})
	}
}

// SortAsc
func BenchmarkIntSlice_SortAsc(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []int
	}{
		{
			name:  "10 elements",
			slice: internal.GenIntSlice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenIntSlice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenIntSlice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenIntSlice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenIntSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IntSlice(bm.slice).SortAsc()
			}
		})
	}
}

// SortDesc
func BenchmarkIntSlice_SortDesc(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []int
	}{
		{
			name:  "10 elements",
			slice: internal.GenIntSlice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenIntSlice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenIntSlice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenIntSlice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenIntSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IntSlice(bm.slice).SortDesc()
			}
		})
	}
}

// Reverse
func BenchmarkIntSlice_Reverse(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []int
	}{
		{
			name:  "10 elements",
			slice: internal.GenIntSlice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenIntSlice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenIntSlice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenIntSlice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenIntSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IntSlice(bm.slice).Reverse()
			}
		})
	}
}

// Filter
func BenchmarkIntSlice_Filter(b *testing.B) {
	benchFunc := func(n int) bool {
		return n%2 == 0
	}
	benchmarks := []struct {
		name  string
		slice []int
	}{
		{
			name:  "10 elements",
			slice: internal.GenIntSlice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenIntSlice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenIntSlice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenIntSlice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenIntSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IntSlice(bm.slice).Filter(benchFunc)
			}
		})
	}
}

// Each
func BenchmarkIntSlice_Each(b *testing.B) {
	var rabbit int
	benchFunc := func(n int) {
		rabbit = n
	}
	benchmarks := []struct {
		name  string
		slice []int
	}{
		{
			name:  "10 elements",
			slice: internal.GenIntSlice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenIntSlice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenIntSlice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenIntSlice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenIntSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IntSlice(bm.slice).Each(benchFunc)
			}
		})
	}

	fmt.Fprintf(ioutil.Discard, "%v", rabbit)
}

// Map
func BenchmarkIntSlice_Map(b *testing.B) {
	benchFunc := func(n int) int {
		n++
		return n
	}
	benchmarks := []struct {
		name  string
		slice []int
	}{
		{
			name:  "10 elements",
			slice: internal.GenIntSlice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenIntSlice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenIntSlice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenIntSlice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenIntSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				IntSlice(bm.slice).Map(benchFunc)
			}
		})
	}
}

// endregion
