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
func TestUIntSlice_IndexOf(t *testing.T) {
	tests := []struct {
		name     string
		slice    []uint
		needle   uint
		expected int
	}{
		{
			name:     "only item",
			slice:    []uint{1},
			needle:   1,
			expected: 0,
		},
		{
			name:     "at index 1",
			slice:    []uint{0, 1, 1},
			needle:   1,
			expected: 1,
		},
		{
			name:     "missing",
			slice:    []uint{1, 2, 3, 4},
			needle:   5,
			expected: NotInSlice,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UIntSlice(test.slice).IndexOf(test.needle)
			if result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		})
	}
}

// Contains
func TestUIntSlice_Contains(t *testing.T) {
	tests := []struct {
		name     string
		slice    []uint
		needle   uint
		expected bool
	}{
		{
			name:     "present",
			slice:    []uint{1, 2, 3, 4},
			needle:   4,
			expected: true,
		},
		{
			name:     "not present",
			slice:    []uint{1, 2, 3, 4},
			needle:   5,
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UIntSlice(test.slice).Contains(test.needle)
			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// SortAsc
func TestUIntSlice_SortAsc(t *testing.T) {
	tests := []struct {
		name     string
		slice    []uint
		expected []uint
	}{
		{
			name:     "empty",
			slice:    []uint{},
			expected: []uint{},
		},
		{
			name:     "already sorted",
			slice:    []uint{0, 1, 2, 3, 4, 5},
			expected: []uint{0, 1, 2, 3, 4, 5},
		},
		{
			name:     "reversed",
			slice:    []uint{5, 4, 3, 2, 1, 0},
			expected: []uint{0, 1, 2, 3, 4, 5},
		},
		{
			name:     "random",
			slice:    []uint{3, 1, 4, 5, 0, 2},
			expected: []uint{0, 1, 2, 3, 4, 5},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UIntSlice(test.slice).SortAsc().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// SortDesc
func TestUIntSlice_SortDesc(t *testing.T) {
	tests := []struct {
		name     string
		slice    []uint
		expected []uint
	}{
		{
			name:     "empty",
			slice:    []uint{},
			expected: []uint{},
		},
		{
			name:     "already sorted",
			slice:    []uint{5, 4, 3, 2, 1, 0},
			expected: []uint{5, 4, 3, 2, 1, 0},
		},
		{
			name:     "reversed",
			slice:    []uint{0, 1, 2, 3, 4, 5},
			expected: []uint{5, 4, 3, 2, 1, 0},
		},
		{
			name:     "random",
			slice:    []uint{3, 1, 4, 5, 0, 2},
			expected: []uint{5, 4, 3, 2, 1, 0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UIntSlice(test.slice).SortDesc().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Unique
func TestUIntSlice_Unique(t *testing.T) {
	tests := []struct {
		name     string
		slice    []uint
		expected []uint
	}{
		{
			name:     "unaffected",
			slice:    []uint{0, 1, 2, 3, 4, 5},
			expected: []uint{0, 1, 2, 3, 4, 5},
		},
		{
			name:     "one extra five",
			slice:    []uint{5, 0, 1, 2, 3, 4, 5},
			expected: []uint{5, 0, 1, 2, 3, 4},
		},
		{
			name:     "extras everywhere",
			slice:    []uint{0, 0, 1, 0, 1, 2, 2, 2, 3, 0, 3, 4, 2, 3, 4, 4, 2, 1, 0},
			expected: []uint{0, 1, 2, 3, 4},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UIntSlice(test.slice).Unique().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Reverse
func TestUIntSlice_Reverse(t *testing.T) {
	tests := []struct {
		name     string
		slice    []uint
		expected []uint
	}{
		{
			name:     "even length",
			slice:    []uint{0, 1, 2, 3, 4, 5},
			expected: []uint{5, 4, 3, 2, 1, 0},
		},
		{
			name:     "odd length",
			slice:    []uint{0, 1, 2, 3, 4, 5, 6},
			expected: []uint{6, 5, 4, 3, 2, 1, 0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UIntSlice(test.slice).Reverse().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Filter
func TestUIntSlice_Filter(t *testing.T) {
	tests := []struct {
		name       string
		slice      []uint
		expected   []uint
		filterFunc func(uint) bool
	}{
		{
			name:       "gt 10",
			slice:      []uint{1, 2, 5, 11, 13, 15},
			expected:   []uint{11, 13, 15},
			filterFunc: func(n uint) bool { return n > 10 },
		},
		{
			name:       "mod 3",
			slice:      []uint{1, 2, 6, 11, 12, 15, 17},
			expected:   []uint{6, 12, 15},
			filterFunc: func(n uint) bool { return n%3 == 0 },
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UIntSlice(test.slice).Filter(test.filterFunc).Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Each
func TestUIntSlice_Each(t *testing.T) {

	var rabbit uint
	tests := []struct {
		name     string
		slice    []uint
		expected uint
		eachFunc func(uint)
	}{
		{
			name:     "add n",
			slice:    []uint{1, 2, 5, 11, 13, 15},
			expected: 47,
			eachFunc: func(n uint) { rabbit += n },
		},
		{
			name:     "subtract n",
			slice:    []uint{1, 2, 6, 8, 12},
			expected: 18,
			eachFunc: func(n uint) { rabbit -= n },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			UIntSlice(test.slice).Each(test.eachFunc)
			if test.expected != rabbit {
				t.Errorf("expected %v, got %v", test.expected, rabbit)
			}
		})
	}
}

// TryEach
func TestUIntSlice_TryEach(t *testing.T) {

	var rabbit uint
	myErr := errors.New("i am an error")
	tests := []struct {
		name      string
		slice     []uint
		expected  int
		expected2 error
		before    func()
		eachFunc  func(uint) error
	}{
		{
			name:      "add n",
			slice:     []uint{1, 2, 5, 11, 13, 15},
			expected:  NotInSlice,
			expected2: nil,
			eachFunc: func(n uint) error {
				rabbit += n
				return nil
			},
		},
		{
			name:      "subtract n",
			slice:     []uint{1, 2, 6, 8, 12},
			expected:  NotInSlice,
			expected2: nil,
			eachFunc: func(n uint) error {
				rabbit -= n
				return nil
			},
		},
		{
			name:      "errors",
			slice:     []uint{1, 2, 5, 11, 13, 15},
			expected:  3,
			expected2: myErr,
			before:    func() { rabbit = 0 },
			eachFunc: func(n uint) error {
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
			e, i := UIntSlice(test.slice).TryEach(test.eachFunc)
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
func TestUIntSlice_IfEach(t *testing.T) {

	var rabbit uint
	tests := []struct {
		name      string
		slice     []uint
		expected  int
		expected2 bool
		before    func()
		err       error
		eachFunc  func(uint) bool
	}{
		{
			name:      "all return true",
			slice:     []uint{1, 2, 5, 11, 13, 15},
			expected:  NotInSlice,
			expected2: true,
			eachFunc: func(n uint) bool {
				rabbit += n
				return true
			},
		},
		{
			name:      "subtract n",
			slice:     []uint{1, 2, 6, 8, 12},
			expected:  NotInSlice,
			expected2: true,
			eachFunc: func(n uint) bool {
				rabbit -= n
				return true
			},
		},
		{
			name:      "breaking",
			slice:     []uint{1, 2, 5, 11, 13, 15},
			expected:  3,
			expected2: false,
			before:    func() { rabbit = 0 },
			eachFunc: func(n uint) bool {
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
			e, i := UIntSlice(test.slice).IfEach(test.eachFunc)
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
func TestUIntSlice_Map(t *testing.T) {
	tests := []struct {
		name     string
		slice    []uint
		expected []uint
		mapFunc  func(uint) uint
	}{
		{
			name:     "add 3",
			slice:    []uint{1, 2, 5, 11, 13, 15},
			expected: []uint{4, 5, 8, 14, 16, 18},
			mapFunc:  func(n uint) uint { return n + 3 },
		},
		{
			name:     "set mod 2",
			slice:    []uint{1, 2, 6, 8, 12, 15, 17},
			expected: []uint{1, 0, 0, 0, 0, 1, 1},
			mapFunc:  func(n uint) uint { return uint(n % 2) },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			UIntSlice(test.slice).Map(test.mapFunc)
			if !reflect.DeepEqual(test.expected, test.slice) {
				t.Errorf("expected %v, got %v", test.expected, test.slice)
			}
		})
	}
}

// Chunk
func TestUIntSlice_Chunk(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		slice    []uint
		expected [][]uint
	}{
		{
			name:     "chunks of 2 no remainder",
			size:     2,
			slice:    []uint{1, 2, 5, 11, 13, 15},
			expected: [][]uint{{1, 2}, {5, 11}, {13, 15}},
		},
		{
			name:     "chunks of 2 with remainder",
			size:     2,
			slice:    []uint{1, 2, 5, 11, 13, 15, 17},
			expected: [][]uint{{1, 2}, {5, 11}, {13, 15}, {17}},
		},
		{
			name:     "chunks of 100",
			size:     100,
			slice:    []uint{1, 2, 5, 11, 13, 15, 17},
			expected: [][]uint{{1, 2, 5, 11, 13, 15, 17}},
		},
		{
			name:     "chunks of 4",
			size:     4,
			slice:    []uint{1, 2, 5, 11, 13, 15, 17},
			expected: [][]uint{{1, 2, 5, 11}, {13, 15, 17}},
		},
		{
			name:     "chunks of 5",
			size:     5,
			slice:    []uint{1, 2, 5, 11, 13, 15, 17},
			expected: [][]uint{{1, 2, 5, 11, 13}, {15, 17}},
		},
		{
			name:     "empty slice",
			size:     5,
			slice:    []uint{},
			expected: [][]uint{},
		},
		{
			name:     "invalid chunk size",
			size:     -1,
			slice:    []uint{1, 2, 5, 11, 13, 15, 17},
			expected: [][]uint{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := UIntSlice(test.slice).Chunk(test.size)
			if !reflect.DeepEqual(test.expected, out) {
				t.Errorf("expected %v, got %v", test.expected, test.slice)
			}
		})
	}
}

// endregion

// region BENCHMARKS
// IndexOf
func BenchmarkUIntSlice_IndexOf(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []uint
	}{
		{
			name:  "10 elements",
			slice: internal.GenUIntSlice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenUIntSlice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenUIntSlice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenUIntSlice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenUIntSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UIntSlice(bm.slice).IndexOf(1)
			}
		})
	}
}

// Contains
func BenchmarkUIntSlice_Contains(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []uint
	}{
		{
			name:  "10 elements",
			slice: internal.GenUIntSlice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenUIntSlice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenUIntSlice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenUIntSlice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenUIntSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UIntSlice(bm.slice).Contains(1)
			}
		})
	}
}

// SortAsc
func BenchmarkUIntSlice_SortAsc(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []uint
	}{
		{
			name:  "10 elements",
			slice: internal.GenUIntSlice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenUIntSlice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenUIntSlice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenUIntSlice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenUIntSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UIntSlice(bm.slice).SortAsc()
			}
		})
	}
}

// SortDesc
func BenchmarkUIntSlice_SortDesc(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []uint
	}{
		{
			name:  "10 elements",
			slice: internal.GenUIntSlice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenUIntSlice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenUIntSlice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenUIntSlice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenUIntSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UIntSlice(bm.slice).SortDesc()
			}
		})
	}
}

// Reverse
func BenchmarkUIntSlice_Reverse(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []uint
	}{
		{
			name:  "10 elements",
			slice: internal.GenUIntSlice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenUIntSlice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenUIntSlice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenUIntSlice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenUIntSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UIntSlice(bm.slice).Reverse()
			}
		})
	}
}

// Filter
func BenchmarkUIntSlice_Filter(b *testing.B) {
	benchFunc := func(n uint) bool {
		return n%2 == 0
	}
	benchmarks := []struct {
		name  string
		slice []uint
	}{
		{
			name:  "10 elements",
			slice: internal.GenUIntSlice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenUIntSlice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenUIntSlice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenUIntSlice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenUIntSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UIntSlice(bm.slice).Filter(benchFunc)
			}
		})
	}
}

// Each
func BenchmarkUIntSlice_Each(b *testing.B) {
	var rabbit uint
	benchFunc := func(n uint) {
		rabbit = n
	}
	benchmarks := []struct {
		name  string
		slice []uint
	}{
		{
			name:  "10 elements",
			slice: internal.GenUIntSlice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenUIntSlice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenUIntSlice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenUIntSlice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenUIntSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UIntSlice(bm.slice).Each(benchFunc)
			}
		})
	}

	fmt.Fprintf(ioutil.Discard, "%v", rabbit)
}

// Map
func BenchmarkUIntSlice_Map(b *testing.B) {
	benchFunc := func(n uint) uint {
		n++
		return n
	}
	benchmarks := []struct {
		name  string
		slice []uint
	}{
		{
			name:  "10 elements",
			slice: internal.GenUIntSlice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenUIntSlice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenUIntSlice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenUIntSlice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenUIntSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UIntSlice(bm.slice).Map(benchFunc)
			}
		})
	}
}

// endregion
