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
func TestUIntPtrSlice_IndexOf(t *testing.T) {
	tests := []struct {
		name     string
		slice    []uintptr
		needle   uintptr
		expected int
	}{
		{
			name:     "only item",
			slice:    []uintptr{1},
			needle:   1,
			expected: 0,
		},
		{
			name:     "at index 1",
			slice:    []uintptr{0, 1, 1},
			needle:   1,
			expected: 1,
		},
		{
			name:     "missing",
			slice:    []uintptr{1, 2, 3, 4},
			needle:   5,
			expected: NotInSlice,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UIntPtrSlice(test.slice).IndexOf(test.needle)
			if result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		})
	}
}

// Contains
func TestUIntPtrSlice_Contains(t *testing.T) {
	tests := []struct {
		name     string
		slice    []uintptr
		needle   uintptr
		expected bool
	}{
		{
			name:     "present",
			slice:    []uintptr{1, 2, 3, 4},
			needle:   4,
			expected: true,
		},
		{
			name:     "not present",
			slice:    []uintptr{1, 2, 3, 4},
			needle:   5,
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UIntPtrSlice(test.slice).Contains(test.needle)
			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// SortAsc
func TestUIntPtrSlice_SortAsc(t *testing.T) {
	tests := []struct {
		name     string
		slice    []uintptr
		expected []uintptr
	}{
		{
			name:     "empty",
			slice:    []uintptr{},
			expected: []uintptr{},
		},
		{
			name:     "already sorted",
			slice:    []uintptr{0, 1, 2, 3, 4, 5},
			expected: []uintptr{0, 1, 2, 3, 4, 5},
		},
		{
			name:     "reversed",
			slice:    []uintptr{5, 4, 3, 2, 1, 0},
			expected: []uintptr{0, 1, 2, 3, 4, 5},
		},
		{
			name:     "random",
			slice:    []uintptr{3, 1, 4, 5, 0, 2},
			expected: []uintptr{0, 1, 2, 3, 4, 5},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UIntPtrSlice(test.slice).SortAsc().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// SortDesc
func TestUIntPtrSlice_SortDesc(t *testing.T) {
	tests := []struct {
		name     string
		slice    []uintptr
		expected []uintptr
	}{
		{
			name:     "empty",
			slice:    []uintptr{},
			expected: []uintptr{},
		},
		{
			name:     "already sorted",
			slice:    []uintptr{5, 4, 3, 2, 1, 0},
			expected: []uintptr{5, 4, 3, 2, 1, 0},
		},
		{
			name:     "reversed",
			slice:    []uintptr{0, 1, 2, 3, 4, 5},
			expected: []uintptr{5, 4, 3, 2, 1, 0},
		},
		{
			name:     "random",
			slice:    []uintptr{3, 1, 4, 5, 0, 2},
			expected: []uintptr{5, 4, 3, 2, 1, 0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UIntPtrSlice(test.slice).SortDesc().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Unique
func TestUIntPtrSlice_Unique(t *testing.T) {
	tests := []struct {
		name     string
		slice    []uintptr
		expected []uintptr
	}{
		{
			name:     "unaffected",
			slice:    []uintptr{0, 1, 2, 3, 4, 5},
			expected: []uintptr{0, 1, 2, 3, 4, 5},
		},
		{
			name:     "one extra five",
			slice:    []uintptr{5, 0, 1, 2, 3, 4, 5},
			expected: []uintptr{5, 0, 1, 2, 3, 4},
		},
		{
			name:     "extras everywhere",
			slice:    []uintptr{0, 0, 1, 0, 1, 2, 2, 2, 3, 0, 3, 4, 2, 3, 4, 4, 2, 1, 0},
			expected: []uintptr{0, 1, 2, 3, 4},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UIntPtrSlice(test.slice).Unique().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Reverse
func TestUIntPtrSlice_Reverse(t *testing.T) {
	tests := []struct {
		name     string
		slice    []uintptr
		expected []uintptr
	}{
		{
			name:     "even length",
			slice:    []uintptr{0, 1, 2, 3, 4, 5},
			expected: []uintptr{5, 4, 3, 2, 1, 0},
		},
		{
			name:     "odd length",
			slice:    []uintptr{0, 1, 2, 3, 4, 5, 6},
			expected: []uintptr{6, 5, 4, 3, 2, 1, 0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UIntPtrSlice(test.slice).Reverse().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Filter
func TestUIntPtrSlice_Filter(t *testing.T) {
	tests := []struct {
		name       string
		slice      []uintptr
		expected   []uintptr
		filterFunc func(uintptr) bool
	}{
		{
			name:       "gt 10",
			slice:      []uintptr{1, 2, 5, 11, 13, 15},
			expected:   []uintptr{11, 13, 15},
			filterFunc: func(n uintptr) bool { return n > 10 },
		},
		{
			name:       "mod 3",
			slice:      []uintptr{1, 2, 6, 11, 12, 15, 17},
			expected:   []uintptr{6, 12, 15},
			filterFunc: func(n uintptr) bool { return n%3 == 0 },
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UIntPtrSlice(test.slice).Filter(test.filterFunc).Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Each
func TestUIntPtrSlice_Each(t *testing.T) {

	var rabbit uintptr
	tests := []struct {
		name     string
		slice    []uintptr
		expected uintptr
		eachFunc func(uintptr)
	}{
		{
			name:     "add n",
			slice:    []uintptr{1, 2, 5, 11, 13, 15},
			expected: 47,
			eachFunc: func(n uintptr) { rabbit += n },
		},
		{
			name:     "subtract n",
			slice:    []uintptr{1, 2, 6, 8, 12},
			expected: 18,
			eachFunc: func(n uintptr) { rabbit -= n },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			UIntPtrSlice(test.slice).Each(test.eachFunc)
			if test.expected != rabbit {
				t.Errorf("expected %v, got %v", test.expected, rabbit)
			}
		})
	}
}

// TryEach
func TestUIntPtrSlice_TryEach(t *testing.T) {

	var rabbit uintptr
	myErr := errors.New("i am an error")
	tests := []struct {
		name      string
		slice     []uintptr
		expected  int
		expected2 error
		before    func()
		eachFunc  func(uintptr) error
	}{
		{
			name:      "add n",
			slice:     []uintptr{1, 2, 5, 11, 13, 15},
			expected:  NotInSlice,
			expected2: nil,
			eachFunc: func(n uintptr) error {
				rabbit += n
				return nil
			},
		},
		{
			name:      "subtract n",
			slice:     []uintptr{1, 2, 6, 8, 12},
			expected:  NotInSlice,
			expected2: nil,
			eachFunc: func(n uintptr) error {
				rabbit -= n
				return nil
			},
		},
		{
			name:      "errors",
			slice:     []uintptr{1, 2, 5, 11, 13, 15},
			expected:  3,
			expected2: myErr,
			before:    func() { rabbit = 0 },
			eachFunc: func(n uintptr) error {
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
			e, i := UIntPtrSlice(test.slice).TryEach(test.eachFunc)
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
func TestUIntPtrSlice_IfEach(t *testing.T) {

	var rabbit uintptr
	tests := []struct {
		name      string
		slice     []uintptr
		expected  int
		expected2 bool
		before    func()
		err       error
		eachFunc  func(uintptr) bool
	}{
		{
			name:      "all return true",
			slice:     []uintptr{1, 2, 5, 11, 13, 15},
			expected:  NotInSlice,
			expected2: true,
			eachFunc: func(n uintptr) bool {
				rabbit += n
				return true
			},
		},
		{
			name:      "subtract n",
			slice:     []uintptr{1, 2, 6, 8, 12},
			expected:  NotInSlice,
			expected2: true,
			eachFunc: func(n uintptr) bool {
				rabbit -= n
				return true
			},
		},
		{
			name:      "breaking",
			slice:     []uintptr{1, 2, 5, 11, 13, 15},
			expected:  3,
			expected2: false,
			before:    func() { rabbit = 0 },
			eachFunc: func(n uintptr) bool {
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
			e, i := UIntPtrSlice(test.slice).IfEach(test.eachFunc)
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
func TestUIntPtrSlice_Map(t *testing.T) {
	tests := []struct {
		name     string
		slice    []uintptr
		expected []uintptr
		mapFunc  func(uintptr) uintptr
	}{
		{
			name:     "add 3",
			slice:    []uintptr{1, 2, 5, 11, 13, 15},
			expected: []uintptr{4, 5, 8, 14, 16, 18},
			mapFunc:  func(n uintptr) uintptr { return n + 3 },
		},
		{
			name:     "set mod 2",
			slice:    []uintptr{1, 2, 6, 8, 12, 15, 17},
			expected: []uintptr{1, 0, 0, 0, 0, 1, 1},
			mapFunc:  func(n uintptr) uintptr { return uintptr(n % 2) },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			UIntPtrSlice(test.slice).Map(test.mapFunc)
			if !reflect.DeepEqual(test.expected, test.slice) {
				t.Errorf("expected %v, got %v", test.expected, test.slice)
			}
		})
	}
}

// Chunk
func TestUIntPtrSlice_Chunk(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		slice    []uintptr
		expected [][]uintptr
	}{
		{
			name:     "chunks of 2 no remainder",
			size:     2,
			slice:    []uintptr{1, 2, 5, 11, 13, 15},
			expected: [][]uintptr{[]uintptr{1, 2}, []uintptr{5, 11}, []uintptr{13, 15}},
		},
		{
			name:     "chunks of 2 with remainder",
			size:     2,
			slice:    []uintptr{1, 2, 5, 11, 13, 15, 17},
			expected: [][]uintptr{[]uintptr{1, 2}, []uintptr{5, 11}, []uintptr{13, 15}, []uintptr{17}},
		},
		{
			name:     "chunks of 100",
			size:     100,
			slice:    []uintptr{1, 2, 5, 11, 13, 15, 17},
			expected: [][]uintptr{[]uintptr{1, 2, 5, 11, 13, 15, 17}},
		},
		{
			name:     "chunks of 4",
			size:     4,
			slice:    []uintptr{1, 2, 5, 11, 13, 15, 17},
			expected: [][]uintptr{[]uintptr{1, 2, 5, 11}, []uintptr{13, 15, 17}},
		},
		{
			name:     "chunks of 5",
			size:     5,
			slice:    []uintptr{1, 2, 5, 11, 13, 15, 17},
			expected: [][]uintptr{[]uintptr{1, 2, 5, 11, 13}, []uintptr{15, 17}},
		},
		{
			name:     "empty slice",
			size:     5,
			slice:    []uintptr{},
			expected: [][]uintptr{},
		},
		{
			name:     "invalid chunk size",
			size:     -1,
			slice:    []uintptr{1, 2, 5, 11, 13, 15, 17},
			expected: [][]uintptr{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := UIntPtrSlice(test.slice).Chunk(test.size)
			if !reflect.DeepEqual(test.expected, out) {
				t.Errorf("expected %v, got %v", test.expected, test.slice)
			}
		})
	}
}

// endregion

// region BENCHMARKS
// IndexOf
func BenchmarkUIntPtrSlice_IndexOf(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []uintptr
	}{
		{
			name:  "10 elements",
			slice: internal.GenUIntPtrSlice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenUIntPtrSlice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenUIntPtrSlice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenUIntPtrSlice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenUIntPtrSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UIntPtrSlice(bm.slice).IndexOf(1)
			}
		})
	}
}

// Contains
func BenchmarkUIntPtrSlice_Contains(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []uintptr
	}{
		{
			name:  "10 elements",
			slice: internal.GenUIntPtrSlice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenUIntPtrSlice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenUIntPtrSlice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenUIntPtrSlice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenUIntPtrSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UIntPtrSlice(bm.slice).Contains(1)
			}
		})
	}
}

// SortAsc
func BenchmarkUIntPtrSlice_SortAsc(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []uintptr
	}{
		{
			name:  "10 elements",
			slice: internal.GenUIntPtrSlice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenUIntPtrSlice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenUIntPtrSlice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenUIntPtrSlice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenUIntPtrSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UIntPtrSlice(bm.slice).SortAsc()
			}
		})
	}
}

// SortDesc
func BenchmarkUIntPtrSlice_SortDesc(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []uintptr
	}{
		{
			name:  "10 elements",
			slice: internal.GenUIntPtrSlice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenUIntPtrSlice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenUIntPtrSlice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenUIntPtrSlice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenUIntPtrSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UIntPtrSlice(bm.slice).SortDesc()
			}
		})
	}
}

// Reverse
func BenchmarkUIntPtrSlice_Reverse(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []uintptr
	}{
		{
			name:  "10 elements",
			slice: internal.GenUIntPtrSlice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenUIntPtrSlice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenUIntPtrSlice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenUIntPtrSlice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenUIntPtrSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UIntPtrSlice(bm.slice).Reverse()
			}
		})
	}
}

// Filter
func BenchmarkUIntPtrSlice_Filter(b *testing.B) {
	benchFunc := func(n uintptr) bool {
		return n%2 == 0
	}
	benchmarks := []struct {
		name  string
		slice []uintptr
	}{
		{
			name:  "10 elements",
			slice: internal.GenUIntPtrSlice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenUIntPtrSlice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenUIntPtrSlice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenUIntPtrSlice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenUIntPtrSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UIntPtrSlice(bm.slice).Filter(benchFunc)
			}
		})
	}
}

// Each
func BenchmarkUIntPtrSlice_Each(b *testing.B) {
	var rabbit uintptr
	benchFunc := func(n uintptr) {
		rabbit = n
	}
	benchmarks := []struct {
		name  string
		slice []uintptr
	}{
		{
			name:  "10 elements",
			slice: internal.GenUIntPtrSlice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenUIntPtrSlice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenUIntPtrSlice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenUIntPtrSlice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenUIntPtrSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UIntPtrSlice(bm.slice).Each(benchFunc)
			}
		})
	}

	fmt.Fprintf(ioutil.Discard, "%v", rabbit)
}

// Map
func BenchmarkUIntPtrSlice_Map(b *testing.B) {
	benchFunc := func(n uintptr) uintptr {
		n++
		return n
	}
	benchmarks := []struct {
		name  string
		slice []uintptr
	}{
		{
			name:  "10 elements",
			slice: internal.GenUIntPtrSlice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenUIntPtrSlice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenUIntPtrSlice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenUIntPtrSlice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenUIntPtrSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UIntPtrSlice(bm.slice).Map(benchFunc)
			}
		})
	}
}

// endregion
