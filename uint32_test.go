package slices

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/schigh/slices/internal"
)

/* generated by github.com/schigh/slices/base/base.go.  do not edit. */

//region TESTS
// IndexOf
func TestUInt32Slice_IndexOf(t *testing.T) {
	tests := []struct {
		name     string
		slice    []uint32
		needle   uint32
		expected int
	}{
		{
			name:     "only item",
			slice:    []uint32{1},
			needle:   1,
			expected: 0,
		},
		{
			name:     "at index 1",
			slice:    []uint32{0, 1, 1},
			needle:   1,
			expected: 1,
		},
		{
			name:     "missing",
			slice:    []uint32{1, 2, 3, 4},
			needle:   5,
			expected: NotInSlice,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UInt32Slice(test.slice).IndexOf(test.needle)
			if result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		})
	}
}

// Contains
func TestUInt32Slice_Contains(t *testing.T) {
	tests := []struct {
		name     string
		slice    []uint32
		needle   uint32
		expected bool
	}{
		{
			name:     "present",
			slice:    []uint32{1, 2, 3, 4},
			needle:   4,
			expected: true,
		},
		{
			name:     "not present",
			slice:    []uint32{1, 2, 3, 4},
			needle:   5,
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UInt32Slice(test.slice).Contains(test.needle)
			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// SortAsc
func TestUInt32Slice_SortAsc(t *testing.T) {
	tests := []struct {
		name     string
		slice    []uint32
		expected []uint32
	}{
		{
			name:     "empty",
			slice:    []uint32{},
			expected: []uint32{},
		},
		{
			name:     "already sorted",
			slice:    []uint32{0, 1, 2, 3, 4, 5},
			expected: []uint32{0, 1, 2, 3, 4, 5},
		},
		{
			name:     "reversed",
			slice:    []uint32{5, 4, 3, 2, 1, 0},
			expected: []uint32{0, 1, 2, 3, 4, 5},
		},
		{
			name:     "random",
			slice:    []uint32{3, 1, 4, 5, 0, 2},
			expected: []uint32{0, 1, 2, 3, 4, 5},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UInt32Slice(test.slice).SortAsc().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// SortDesc
func TestUInt32Slice_SortDesc(t *testing.T) {
	tests := []struct {
		name     string
		slice    []uint32
		expected []uint32
	}{
		{
			name:     "empty",
			slice:    []uint32{},
			expected: []uint32{},
		},
		{
			name:     "already sorted",
			slice:    []uint32{5, 4, 3, 2, 1, 0},
			expected: []uint32{5, 4, 3, 2, 1, 0},
		},
		{
			name:     "reversed",
			slice:    []uint32{0, 1, 2, 3, 4, 5},
			expected: []uint32{5, 4, 3, 2, 1, 0},
		},
		{
			name:     "random",
			slice:    []uint32{3, 1, 4, 5, 0, 2},
			expected: []uint32{5, 4, 3, 2, 1, 0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UInt32Slice(test.slice).SortDesc().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Unique
func TestUInt32Slice_Unique(t *testing.T) {
	tests := []struct {
		name     string
		slice    []uint32
		expected []uint32
	}{
		{
			name:     "unaffected",
			slice:    []uint32{0, 1, 2, 3, 4, 5},
			expected: []uint32{0, 1, 2, 3, 4, 5},
		},
		{
			name:     "one extra five",
			slice:    []uint32{5, 0, 1, 2, 3, 4, 5},
			expected: []uint32{5, 0, 1, 2, 3, 4},
		},
		{
			name:     "extras everywhere",
			slice:    []uint32{0, 0, 1, 0, 1, 2, 2, 2, 3, 0, 3, 4, 2, 3, 4, 4, 2, 1, 0},
			expected: []uint32{0, 1, 2, 3, 4},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UInt32Slice(test.slice).Unique().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Reverse
func TestUInt32Slice_Reverse(t *testing.T) {
	tests := []struct {
		name     string
		slice    []uint32
		expected []uint32
	}{
		{
			name:     "even length",
			slice:    []uint32{0, 1, 2, 3, 4, 5},
			expected: []uint32{5, 4, 3, 2, 1, 0},
		},
		{
			name:     "odd length",
			slice:    []uint32{0, 1, 2, 3, 4, 5, 6},
			expected: []uint32{6, 5, 4, 3, 2, 1, 0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UInt32Slice(test.slice).Reverse().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Filter
func TestUInt32Slice_Filter(t *testing.T) {
	tests := []struct {
		name       string
		slice      []uint32
		expected   []uint32
		filterFunc func(uint32) bool
	}{
		{
			name:       "gt 10",
			slice:      []uint32{1, 2, 5, 11, 13, 15},
			expected:   []uint32{11, 13, 15},
			filterFunc: func(n uint32) bool { return n > 10 },
		},
		{
			name:       "mod 3",
			slice:      []uint32{1, 2, 6, 11, 12, 15, 17},
			expected:   []uint32{6, 12, 15},
			filterFunc: func(n uint32) bool { return n%3 == 0 },
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UInt32Slice(test.slice).Filter(test.filterFunc).Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Each
func TestUInt32Slice_Each(t *testing.T) {

	var rabbit uint32
	tests := []struct {
		name     string
		slice    []uint32
		expected uint32
		eachFunc func(uint32)
	}{
		{
			name:     "add n",
			slice:    []uint32{1, 2, 5, 11, 13, 15},
			expected: 47,
			eachFunc: func(n uint32) { rabbit += n },
		},
		{
			name:     "subtract n",
			slice:    []uint32{1, 2, 6, 8, 12},
			expected: 18,
			eachFunc: func(n uint32) { rabbit -= n },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			UInt32Slice(test.slice).Each(test.eachFunc)
			if test.expected != rabbit {
				t.Errorf("expected %v, got %v", test.expected, rabbit)
			}
		})
	}
}

// Map
func TestUInt32Slice_Map(t *testing.T) {
	tests := []struct {
		name     string
		slice    []uint32
		expected []uint32
		mapFunc  func(uint32) uint32
	}{
		{
			name:     "add 3",
			slice:    []uint32{1, 2, 5, 11, 13, 15},
			expected: []uint32{4, 5, 8, 14, 16, 18},
			mapFunc:  func(n uint32) uint32 { return n + 3 },
		},
		{
			name:     "set mod 2",
			slice:    []uint32{1, 2, 6, 8, 12, 15, 17},
			expected: []uint32{1, 0, 0, 0, 0, 1, 1},
			mapFunc:  func(n uint32) uint32 { return uint32(n % 2) },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			UInt32Slice(test.slice).Map(test.mapFunc)
			if !reflect.DeepEqual(test.expected, test.slice) {
				t.Errorf("expected %v, got %v", test.expected, test.slice)
			}
		})
	}
}

// endregion

// region BENCHMARKS
// IndexOf
func BenchmarkUInt32Slice_IndexOf(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []uint32
	}{
		{
			name:  "10 elements",
			slice: internal.GenUInt32Slice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenUInt32Slice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenUInt32Slice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenUInt32Slice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenUInt32Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UInt32Slice(bm.slice).IndexOf(1)
			}
		})
	}
}

// Contains
func BenchmarkUInt32Slice_Contains(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []uint32
	}{
		{
			name:  "10 elements",
			slice: internal.GenUInt32Slice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenUInt32Slice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenUInt32Slice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenUInt32Slice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenUInt32Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UInt32Slice(bm.slice).Contains(1)
			}
		})
	}
}

// SortAsc
func BenchmarkUInt32Slice_SortAsc(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []uint32
	}{
		{
			name:  "10 elements",
			slice: internal.GenUInt32Slice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenUInt32Slice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenUInt32Slice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenUInt32Slice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenUInt32Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UInt32Slice(bm.slice).SortAsc()
			}
		})
	}
}

// SortDesc
func BenchmarkUInt32Slice_SortDesc(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []uint32
	}{
		{
			name:  "10 elements",
			slice: internal.GenUInt32Slice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenUInt32Slice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenUInt32Slice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenUInt32Slice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenUInt32Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UInt32Slice(bm.slice).SortDesc()
			}
		})
	}
}

// Reverse
func BenchmarkUInt32Slice_Reverse(b *testing.B) {
	benchmarks := []struct {
		name  string
		slice []uint32
	}{
		{
			name:  "10 elements",
			slice: internal.GenUInt32Slice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenUInt32Slice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenUInt32Slice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenUInt32Slice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenUInt32Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UInt32Slice(bm.slice).Reverse()
			}
		})
	}
}

// Filter
func BenchmarkUInt32Slice_Filter(b *testing.B) {
	benchFunc := func(n uint32) bool {
		return n%2 == 0
	}
	benchmarks := []struct {
		name  string
		slice []uint32
	}{
		{
			name:  "10 elements",
			slice: internal.GenUInt32Slice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenUInt32Slice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenUInt32Slice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenUInt32Slice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenUInt32Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UInt32Slice(bm.slice).Filter(benchFunc)
			}
		})
	}
}

// Each
func BenchmarkUInt32Slice_Each(b *testing.B) {
	var rabbit uint32
	benchFunc := func(n uint32) {
		rabbit = n
	}
	benchmarks := []struct {
		name  string
		slice []uint32
	}{
		{
			name:  "10 elements",
			slice: internal.GenUInt32Slice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenUInt32Slice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenUInt32Slice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenUInt32Slice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenUInt32Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UInt32Slice(bm.slice).Each(benchFunc)
			}
		})
	}

	fmt.Fprintf(ioutil.Discard, "%v", rabbit)
}

// Map
func BenchmarkUInt32Slice_Map(b *testing.B) {
	benchFunc := func(n uint32) uint32 {
		n++
		return n
	}
	benchmarks := []struct {
		name  string
		slice []uint32
	}{
		{
			name:  "10 elements",
			slice: internal.GenUInt32Slice(10),
		},
		{
			name:  "100 elements",
			slice: internal.GenUInt32Slice(100),
		},
		{
			name:  "1000 elements",
			slice: internal.GenUInt32Slice(1000),
		},
		{
			name:  "10000 elements",
			slice: internal.GenUInt32Slice(10000),
		},
		{
			name:  "100000 elements",
			slice: internal.GenUInt32Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UInt32Slice(bm.slice).Map(benchFunc)
			}
		})
	}
}

// endregion
