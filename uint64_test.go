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
func TestUInt64Slice_IndexOf(t *testing.T) {
	tests := []struct {
		name string
		slice []uint64
		needle uint64
		expected int
	}{
		{
			name: "only item",
			slice: []uint64{1},
			needle: 1,
			expected: 0,
		},
		{
			name: "at index 1",
			slice: []uint64{0,1,1},
			needle: 1,
			expected: 1,
		},
		{
			name: "missing",
			slice: []uint64{1,2,3,4},
			needle: 5,
			expected: NotInSlice,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UInt64Slice(test.slice).IndexOf(test.needle)
			if result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		})
	}
}

// Contains
func TestUInt64Slice_Contains(t *testing.T) {
	tests := []struct {
		name string
		slice []uint64
		needle uint64
		expected bool
	}{
		{
			name: "present",
			slice: []uint64{1,2,3,4},
			needle: 4,
			expected: true,
		},
		{
			name: "not present",
			slice: []uint64{1,2,3,4},
			needle: 5,
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UInt64Slice(test.slice).Contains(test.needle)
			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// SortAsc
func TestUInt64Slice_SortAsc(t *testing.T) {
	tests := []struct {
		name string
		slice []uint64
		expected []uint64
	}{
		{
			name: "empty",
			slice: []uint64{},
			expected: []uint64{},
		},
		{
			name: "already sorted",
			slice: []uint64{0,1,2,3,4,5},
			expected: []uint64{0,1,2,3,4,5},
		},
		{
			name: "reversed",
			slice: []uint64{5,4,3,2,1,0},
			expected: []uint64{0,1,2,3,4,5},
		},
		{
			name: "random",
			slice: []uint64{3,1,4,5,0,2},
			expected: []uint64{0,1,2,3,4,5},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UInt64Slice(test.slice).SortAsc().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// SortDesc
func TestUInt64Slice_SortDesc(t *testing.T) {
	tests := []struct {
		name string
		slice []uint64
		expected []uint64
	}{
		{
			name: "empty",
			slice: []uint64{},
			expected: []uint64{},
		},
		{
			name: "already sorted",
			slice: []uint64{5,4,3,2,1,0},
			expected: []uint64{5,4,3,2,1,0},
		},
		{
			name: "reversed",
			slice: []uint64{0,1,2,3,4,5},
			expected: []uint64{5,4,3,2,1,0},
		},
		{
			name: "random",
			slice: []uint64{3,1,4,5,0,2},
			expected: []uint64{5,4,3,2,1,0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UInt64Slice(test.slice).SortDesc().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Unique
func TestUInt64Slice_Unique(t *testing.T) {
	tests := []struct {
		name string
		slice []uint64
		expected []uint64
	}{
		{
			name: "unaffected",
			slice: []uint64{0,1,2,3,4,5},
			expected: []uint64{0,1,2,3,4,5},
		},
		{
			name: "one extra five",
			slice: []uint64{5,0,1,2,3,4,5},
			expected: []uint64{5,0,1,2,3,4},
		},
		{
			name: "extras everywhere",
			slice: []uint64{0,0,1,0,1,2,2,2,3,0,3,4,2,3,4,4,2,1,0},
			expected: []uint64{0,1,2,3,4},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UInt64Slice(test.slice).Unique().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Reverse
func TestUInt64Slice_Reverse(t *testing.T) {
	tests := []struct {
		name string
		slice []uint64
		expected []uint64
	}{
		{
			name: "even length",
			slice: []uint64{0,1,2,3,4,5},
			expected: []uint64{5,4,3,2,1,0},
		},
		{
			name: "odd length",
			slice: []uint64{0,1,2,3,4,5,6},
			expected: []uint64{6,5,4,3,2,1,0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UInt64Slice(test.slice).Reverse().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Filter
func TestUInt64Slice_Filter(t *testing.T) {
	tests := []struct {
		name string
		slice []uint64
		expected []uint64
		filterFunc func(uint64) bool
	}{
		{
			name: "gt 10",
			slice: []uint64{1, 2, 5, 11, 13, 15},
			expected: []uint64{11, 13, 15},
			filterFunc: func(n uint64) bool { return n > 10 },
		},
		{
			name: "mod 3",
			slice: []uint64{1, 2, 6, 11, 12, 15, 17},
			expected: []uint64{6, 12, 15},
			filterFunc: func(n uint64) bool { return n%3 == 0 },
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UInt64Slice(test.slice).Filter(test.filterFunc).Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Each
func TestUInt64Slice_Each(t *testing.T) {

	var rabbit uint64
	tests := []struct {
		name string
		slice []uint64
		expected uint64
		eachFunc func(uint64)
	}{
		{
			name: "add n",
			slice: []uint64{1, 2, 5, 11, 13, 15},
			expected: 47,
			eachFunc: func(n uint64) { rabbit += n },
		},
		{
			name: "subtract n",
			slice: []uint64{1, 2, 6, 8, 12},
			expected: 18,
			eachFunc: func(n uint64) { rabbit -= n },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			UInt64Slice(test.slice).Each(test.eachFunc)
			if test.expected != rabbit {
				t.Errorf("expected %v, got %v", test.expected, rabbit)
			}
		})
	}
}

// Map
func TestUInt64Slice_Map(t *testing.T) {
	tests := []struct {
		name string
		slice []uint64
		expected []uint64
		mapFunc func(uint64) uint64
	}{
		{
			name: "add 3",
			slice: []uint64{1, 2, 5, 11, 13, 15},
			expected: []uint64{4, 5, 8, 14, 16, 18},
			mapFunc: func(n uint64) uint64 { return n + 3 },
		},
		{
			name: "set mod 2",
			slice: []uint64{1, 2, 6, 8, 12, 15, 17},
			expected: []uint64{1, 0, 0, 0, 0, 1, 1},
			mapFunc: func(n uint64) uint64 { return uint64(n%2) },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			UInt64Slice(test.slice).Map(test.mapFunc)
			if !reflect.DeepEqual(test.expected, test.slice) {
				t.Errorf("expected %v, got %v", test.expected, test.slice)
			}
		})
	}
}

// endregion

// region BENCHMARKS
// IndexOf
func BenchmarkUInt64Slice_IndexOf(b *testing.B) {
	benchmarks := []struct {
		name string
		slice []uint64
	}{
		{
			name: "10 elements",
			slice: internal.GenUInt64Slice(10),
		},
		{
			name: "100 elements",
			slice: internal.GenUInt64Slice(100),
		},
		{
			name: "1000 elements",
			slice: internal.GenUInt64Slice(1000),
		},
		{
			name: "10000 elements",
			slice: internal.GenUInt64Slice(10000),
		},
		{
			name: "100000 elements",
			slice: internal.GenUInt64Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UInt64Slice(bm.slice).IndexOf(1)
			}
		})
	}
}

// Contains
func BenchmarkUInt64Slice_Contains(b *testing.B) {
	benchmarks := []struct {
		name string
		slice []uint64
	}{
		{
			name: "10 elements",
			slice: internal.GenUInt64Slice(10),
		},
		{
			name: "100 elements",
			slice: internal.GenUInt64Slice(100),
		},
		{
			name: "1000 elements",
			slice: internal.GenUInt64Slice(1000),
		},
		{
			name: "10000 elements",
			slice: internal.GenUInt64Slice(10000),
		},
		{
			name: "100000 elements",
			slice: internal.GenUInt64Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UInt64Slice(bm.slice).Contains(1)
			}
		})
	}
}

// SortAsc
func BenchmarkUInt64Slice_SortAsc(b *testing.B) {
	benchmarks := []struct {
		name string
		slice []uint64
	}{
		{
			name: "10 elements",
			slice: internal.GenUInt64Slice(10),
		},
		{
			name: "100 elements",
			slice: internal.GenUInt64Slice(100),
		},
		{
			name: "1000 elements",
			slice: internal.GenUInt64Slice(1000),
		},
		{
			name: "10000 elements",
			slice: internal.GenUInt64Slice(10000),
		},
		{
			name: "100000 elements",
			slice: internal.GenUInt64Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UInt64Slice(bm.slice).SortAsc()
			}
		})
	}
}

// SortDesc
func BenchmarkUInt64Slice_SortDesc(b *testing.B) {
	benchmarks := []struct {
		name string
		slice []uint64
	}{
		{
			name: "10 elements",
			slice: internal.GenUInt64Slice(10),
		},
		{
			name: "100 elements",
			slice: internal.GenUInt64Slice(100),
		},
		{
			name: "1000 elements",
			slice: internal.GenUInt64Slice(1000),
		},
		{
			name: "10000 elements",
			slice: internal.GenUInt64Slice(10000),
		},
		{
			name: "100000 elements",
			slice: internal.GenUInt64Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UInt64Slice(bm.slice).SortDesc()
			}
		})
	}
}

// Reverse
func BenchmarkUInt64Slice_Reverse(b *testing.B) {
	benchmarks := []struct {
		name string
		slice []uint64
	}{
		{
			name: "10 elements",
			slice: internal.GenUInt64Slice(10),
		},
		{
			name: "100 elements",
			slice: internal.GenUInt64Slice(100),
		},
		{
			name: "1000 elements",
			slice: internal.GenUInt64Slice(1000),
		},
		{
			name: "10000 elements",
			slice: internal.GenUInt64Slice(10000),
		},
		{
			name: "100000 elements",
			slice: internal.GenUInt64Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UInt64Slice(bm.slice).Reverse()
			}
		})
	}
}

// Filter
func BenchmarkUInt64Slice_Filter(b *testing.B) {
	benchFunc := func(n uint64) bool {
		return n%2 == 0
	}
	benchmarks := []struct {
		name string
		slice []uint64
	}{
		{
			name: "10 elements",
			slice: internal.GenUInt64Slice(10),
		},
		{
			name: "100 elements",
			slice: internal.GenUInt64Slice(100),
		},
		{
			name: "1000 elements",
			slice: internal.GenUInt64Slice(1000),
		},
		{
			name: "10000 elements",
			slice: internal.GenUInt64Slice(10000),
		},
		{
			name: "100000 elements",
			slice: internal.GenUInt64Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UInt64Slice(bm.slice).Filter(benchFunc)
			}
		})
	}
}

// Each
func BenchmarkUInt64Slice_Each(b *testing.B) {
	var rabbit uint64
	benchFunc := func(n uint64) {
		rabbit = n
	}
	benchmarks := []struct {
		name string
		slice []uint64
	}{
		{
			name: "10 elements",
			slice: internal.GenUInt64Slice(10),
		},
		{
			name: "100 elements",
			slice: internal.GenUInt64Slice(100),
		},
		{
			name: "1000 elements",
			slice: internal.GenUInt64Slice(1000),
		},
		{
			name: "10000 elements",
			slice: internal.GenUInt64Slice(10000),
		},
		{
			name: "100000 elements",
			slice: internal.GenUInt64Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UInt64Slice(bm.slice).Each(benchFunc)
			}
		})
	}

	fmt.Fprintf(ioutil.Discard, "%v", rabbit)
}

// Map
func BenchmarkUInt64Slice_Map(b *testing.B) {
	benchFunc := func(n uint64) uint64 {
		n++
		return n
	}
	benchmarks := []struct {
		name string
		slice []uint64
	}{
		{
			name: "10 elements",
			slice: internal.GenUInt64Slice(10),
		},
		{
			name: "100 elements",
			slice: internal.GenUInt64Slice(100),
		},
		{
			name: "1000 elements",
			slice: internal.GenUInt64Slice(1000),
		},
		{
			name: "10000 elements",
			slice: internal.GenUInt64Slice(10000),
		},
		{
			name: "100000 elements",
			slice: internal.GenUInt64Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UInt64Slice(bm.slice).Map(benchFunc)
			}
		})
	}
}
// endregion
