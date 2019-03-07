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
func TestUInt8Slice_IndexOf(t *testing.T) {
	tests := []struct {
		name     string
		slice    []uint8
		needle   uint8
		expected int
	}{
		{
			name:     "only item",
			slice:    []uint8{1},
			needle:   1,
			expected: 0,
		},
		{
			name:     "at index 1",
			slice:    []uint8{0, 1, 1},
			needle:   1,
			expected: 1,
		},
		{
			name:     "missing",
			slice:    []uint8{1, 2, 3, 4},
			needle:   5,
			expected: NotInSlice,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UInt8Slice(test.slice).IndexOf(test.needle)
			if result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		})
	}
}

// Contains
func TestUInt8Slice_Contains(t *testing.T) {
	tests := []struct {
		name     string
		slice    []uint8
		needle   uint8
		expected bool
	}{
		{
			name:     "present",
			slice:    []uint8{1, 2, 3, 4},
			needle:   4,
			expected: true,
		},
		{
			name:     "not present",
			slice:    []uint8{1, 2, 3, 4},
			needle:   5,
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UInt8Slice(test.slice).Contains(test.needle)
			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// SortAsc
func TestUInt8Slice_SortAsc(t *testing.T) {
	tests := []struct {
		name     string
		slice    []uint8
		expected []uint8
	}{
		{
			name:     "empty",
			slice:    []uint8{},
			expected: []uint8{},
		},
		{
			name:     "already sorted",
			slice:    []uint8{0, 1, 2, 3, 4, 5},
			expected: []uint8{0, 1, 2, 3, 4, 5},
		},
		{
			name:     "reversed",
			slice:    []uint8{5, 4, 3, 2, 1, 0},
			expected: []uint8{0, 1, 2, 3, 4, 5},
		},
		{
			name:     "random",
			slice:    []uint8{3, 1, 4, 5, 0, 2},
			expected: []uint8{0, 1, 2, 3, 4, 5},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UInt8Slice(test.slice).SortAsc().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// SortDesc
func TestUInt8Slice_SortDesc(t *testing.T) {
	tests := []struct {
		name     string
		slice    []uint8
		expected []uint8
	}{
		{
			name:     "empty",
			slice:    []uint8{},
			expected: []uint8{},
		},
		{
			name:     "already sorted",
			slice:    []uint8{5, 4, 3, 2, 1, 0},
			expected: []uint8{5, 4, 3, 2, 1, 0},
		},
		{
			name:     "reversed",
			slice:    []uint8{0, 1, 2, 3, 4, 5},
			expected: []uint8{5, 4, 3, 2, 1, 0},
		},
		{
			name:     "random",
			slice:    []uint8{3, 1, 4, 5, 0, 2},
			expected: []uint8{5, 4, 3, 2, 1, 0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UInt8Slice(test.slice).SortDesc().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Unique
func TestUInt8Slice_Unique(t *testing.T) {
	tests := []struct {
		name     string
		slice    []uint8
		expected []uint8
	}{
		{
			name:     "unaffected",
			slice:    []uint8{0, 1, 2, 3, 4, 5},
			expected: []uint8{0, 1, 2, 3, 4, 5},
		},
		{
			name:     "one extra five",
			slice:    []uint8{5, 0, 1, 2, 3, 4, 5},
			expected: []uint8{5, 0, 1, 2, 3, 4},
		},
		{
			name:     "extras everywhere",
			slice:    []uint8{0, 0, 1, 0, 1, 2, 2, 2, 3, 0, 3, 4, 2, 3, 4, 4, 2, 1, 0},
			expected: []uint8{0, 1, 2, 3, 4},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UInt8Slice(test.slice).Unique().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Reverse
func TestUInt8Slice_Reverse(t *testing.T) {
	tests := []struct {
		name     string
		slice    []uint8
		expected []uint8
	}{
		{
			name:     "even length",
			slice:    []uint8{0, 1, 2, 3, 4, 5},
			expected: []uint8{5, 4, 3, 2, 1, 0},
		},
		{
			name:     "odd length",
			slice:    []uint8{0, 1, 2, 3, 4, 5, 6},
			expected: []uint8{6, 5, 4, 3, 2, 1, 0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UInt8Slice(test.slice).Reverse().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Filter
func TestUInt8Slice_Filter(t *testing.T) {
	tests := []struct {
		name       string
		slice      []uint8
		expected   []uint8
		filterFunc func(uint8) bool
	}{
		{
			name:       "gt 10",
			slice:      []uint8{1, 2, 5, 11, 13, 15},
			expected:   []uint8{11, 13, 15},
			filterFunc: func(n uint8) bool { return n > 10 },
		},
		{
			name:       "mod 3",
			slice:      []uint8{1, 2, 6, 11, 12, 15, 17},
			expected:   []uint8{6, 12, 15},
			filterFunc: func(n uint8) bool { return n%3 == 0 },
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UInt8Slice(test.slice).Filter(test.filterFunc).Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Each
func TestUInt8Slice_Each(t *testing.T) {

	var rabbit uint8
	tests := []struct {
		name     string
		slice    []uint8
		expected uint8
		eachFunc func(uint8)
	}{
		{
			name:     "add n",
			slice:    []uint8{1, 2, 5, 11, 13, 15},
			expected: 47,
			eachFunc: func(n uint8) { rabbit += n },
		},
		{
			name:     "subtract n",
			slice:    []uint8{1, 2, 6, 8, 12},
			expected: 18,
			eachFunc: func(n uint8) { rabbit -= n },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			UInt8Slice(test.slice).Each(test.eachFunc)
			if test.expected != rabbit {
				t.Errorf("expected %v, got %v", test.expected, rabbit)
			}
		})
	}
}

// TryEach
func TestUInt8Slice_TryEach(t *testing.T) {

	var rabbit uint8
	myErr := errors.New("i am an error")
	tests := []struct {
		name      string
		slice     []uint8
		expected  int
		expected2 error
		before    func()
		eachFunc  func(uint8) error
	}{
		{
			name:      "add n",
			slice:     []uint8{1, 2, 5, 11, 13, 15},
			expected:  NotInSlice,
			expected2: nil,
			eachFunc: func(n uint8) error {
				rabbit += n
				return nil
			},
		},
		{
			name:      "subtract n",
			slice:     []uint8{1, 2, 6, 8, 12},
			expected:  NotInSlice,
			expected2: nil,
			eachFunc: func(n uint8) error {
				rabbit -= n
				return nil
			},
		},
		{
			name:      "errors",
			slice:     []uint8{1, 2, 5, 11, 13, 15},
			expected:  3,
			expected2: myErr,
			before:    func() { rabbit = 0 },
			eachFunc: func(n uint8) error {
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
			e, i := UInt8Slice(test.slice).TryEach(test.eachFunc)
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
func TestUInt8Slice_IfEach(t *testing.T) {

	var rabbit uint8
	tests := []struct {
		name      string
		slice     []uint8
		expected  int
		expected2 bool
		before    func()
		err       error
		eachFunc  func(uint8) bool
	}{
		{
			name:      "all return true",
			slice:     []uint8{1, 2, 5, 11, 13, 15},
			expected:  NotInSlice,
			expected2: true,
			eachFunc: func(n uint8) bool {
				rabbit += n
				return true
			},
		},
		{
			name:      "subtract n",
			slice:     []uint8{1, 2, 6, 8, 12},
			expected:  NotInSlice,
			expected2: true,
			eachFunc: func(n uint8) bool {
				rabbit -= n
				return true
			},
		},
		{
			name:      "breaking",
			slice:     []uint8{1, 2, 5, 11, 13, 15},
			expected:  3,
			expected2: false,
			before:    func() { rabbit = 0 },
			eachFunc: func(n uint8) bool {
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
			e, i := UInt8Slice(test.slice).IfEach(test.eachFunc)
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
func TestUInt8Slice_Map(t *testing.T) {
	tests := []struct {
		name     string
		slice    []uint8
		expected []uint8
		mapFunc  func(uint8) uint8
	}{
		{
			name:     "add 3",
			slice:    []uint8{1, 2, 5, 11, 13, 15},
			expected: []uint8{4, 5, 8, 14, 16, 18},
			mapFunc:  func(n uint8) uint8 { return n + 3 },
		},
		{
			name:     "set mod 2",
			slice:    []uint8{1, 2, 6, 8, 12, 15, 17},
			expected: []uint8{1, 0, 0, 0, 0, 1, 1},
			mapFunc:  func(n uint8) uint8 { return uint8(n % 2) },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			UInt8Slice(test.slice).Map(test.mapFunc)
			if !reflect.DeepEqual(test.expected, test.slice) {
				t.Errorf("expected %v, got %v", test.expected, test.slice)
			}
		})
	}
}

// Chunk
func TestUInt8Slice_Chunk(t *testing.T) {
	tests := []struct {
		name     string
		size     int
		slice    []uint8
		expected [][]uint8
	}{
		{
			name:     "chunks of 2 no remainder",
			size:     2,
			slice:    []uint8{1, 2, 5, 11, 13, 15},
			expected: [][]uint8{{1, 2}, {5, 11}, {13, 15}},
		},
		{
			name:     "chunks of 2 with remainder",
			size:     2,
			slice:    []uint8{1, 2, 5, 11, 13, 15, 17},
			expected: [][]uint8{{1, 2}, {5, 11}, {13, 15}, {17}},
		},
		{
			name:     "chunks of 100",
			size:     100,
			slice:    []uint8{1, 2, 5, 11, 13, 15, 17},
			expected: [][]uint8{{1, 2, 5, 11, 13, 15, 17}},
		},
		{
			name:     "chunks of 4",
			size:     4,
			slice:    []uint8{1, 2, 5, 11, 13, 15, 17},
			expected: [][]uint8{{1, 2, 5, 11}, {13, 15, 17}},
		},
		{
			name:     "chunks of 5",
			size:     5,
			slice:    []uint8{1, 2, 5, 11, 13, 15, 17},
			expected: [][]uint8{{1, 2, 5, 11, 13}, {15, 17}},
		},
		{
			name:     "empty slice",
			size:     5,
			slice:    []uint8{},
			expected: [][]uint8{},
		},
		{
			name:     "invalid chunk size",
			size:     -1,
			slice:    []uint8{1, 2, 5, 11, 13, 15, 17},
			expected: [][]uint8{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out := UInt8Slice(test.slice).Chunk(test.size)
			if !reflect.DeepEqual(test.expected, out) {
				t.Errorf("expected %v, got %v", test.expected, test.slice)
			}
		})
	}
}

// endregion

// region BENCHMARKS
// IndexOf
func BenchmarkUInt8Slice_IndexOf(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []uint8
	}{
		{
			name:  "10 elements",
			slice: internal.GenUInt8Slice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenUInt8Slice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenUInt8Slice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenUInt8Slice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenUInt8Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UInt8Slice(bm.slice).IndexOf(1)
			}
		})
	}
}

// Contains
func BenchmarkUInt8Slice_Contains(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []uint8
	}{
		{
			name:  "10 elements",
			slice: internal.GenUInt8Slice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenUInt8Slice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenUInt8Slice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenUInt8Slice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenUInt8Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UInt8Slice(bm.slice).Contains(1)
			}
		})
	}
}

// SortAsc
func BenchmarkUInt8Slice_SortAsc(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []uint8
	}{
		{
			name:  "10 elements",
			slice: internal.GenUInt8Slice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenUInt8Slice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenUInt8Slice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenUInt8Slice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenUInt8Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UInt8Slice(bm.slice).SortAsc()
			}
		})
	}
}

// SortDesc
func BenchmarkUInt8Slice_SortDesc(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []uint8
	}{
		{
			name:  "10 elements",
			slice: internal.GenUInt8Slice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenUInt8Slice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenUInt8Slice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenUInt8Slice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenUInt8Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UInt8Slice(bm.slice).SortDesc()
			}
		})
	}
}

// Reverse
func BenchmarkUInt8Slice_Reverse(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []uint8
	}{
		{
			name:  "10 elements",
			slice: internal.GenUInt8Slice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenUInt8Slice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenUInt8Slice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenUInt8Slice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenUInt8Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UInt8Slice(bm.slice).Reverse()
			}
		})
	}
}

// Filter
func BenchmarkUInt8Slice_Filter(b *testing.B) {
	benchFunc := func(n uint8) bool {
		return n%2 == 0
	}
	benchmarks := []struct {
		name  string
		slice []uint8
	}{
		{
			name:  "10 elements",
			slice: internal.GenUInt8Slice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenUInt8Slice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenUInt8Slice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenUInt8Slice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenUInt8Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UInt8Slice(bm.slice).Filter(benchFunc)
			}
		})
	}
}

// Each
func BenchmarkUInt8Slice_Each(b *testing.B) {
	var rabbit uint8
	benchFunc := func(n uint8) {
		rabbit = n
	}
	benchmarks := []struct {
		name  string
		slice []uint8
	}{
		{
			name:  "10 elements",
			slice: internal.GenUInt8Slice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenUInt8Slice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenUInt8Slice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenUInt8Slice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenUInt8Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UInt8Slice(bm.slice).Each(benchFunc)
			}
		})
	}

	fmt.Fprintf(ioutil.Discard, "%v", rabbit)
}

// Map
func BenchmarkUInt8Slice_Map(b *testing.B) {
	benchFunc := func(n uint8) uint8 {
		n++
		return n
	}
	benchmarks := []struct {
		name  string
		slice []uint8
	}{
		{
			name:  "10 elements",
			slice: internal.GenUInt8Slice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenUInt8Slice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenUInt8Slice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenUInt8Slice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenUInt8Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UInt8Slice(bm.slice).Map(benchFunc)
			}
		})
	}
}

// endregion
