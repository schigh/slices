
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
func TestUInt16Slice_IndexOf(t *testing.T) {
	tests := []struct {
		name string
		slice []uint16
		needle uint16
		expected int
	}{
		{
			name: "only item",
			slice: []uint16{1},
			needle: 1,
			expected: 0,
		},
		{
			name: "at index 1",
			slice: []uint16{0,1,1},
			needle: 1,
			expected: 1,
		},
		{
			name: "missing",
			slice: []uint16{1,2,3,4},
			needle: 5,
			expected: NotInSlice,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UInt16Slice(test.slice).IndexOf(test.needle)
			if result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		})
	}
}

// Contains
func TestUInt16Slice_Contains(t *testing.T) {
	tests := []struct {
		name string
		slice []uint16
		needle uint16
		expected bool
	}{
		{
			name: "present",
			slice: []uint16{1,2,3,4},
			needle: 4,
			expected: true,
		},
		{
			name: "not present",
			slice: []uint16{1,2,3,4},
			needle: 5,
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UInt16Slice(test.slice).Contains(test.needle)
			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// SortAsc
func TestUInt16Slice_SortAsc(t *testing.T) {
	tests := []struct {
		name string
		slice []uint16
		expected []uint16
	}{
		{
			name: "empty",
			slice: []uint16{},
			expected: []uint16{},
		},
		{
			name: "already sorted",
			slice: []uint16{0,1,2,3,4,5},
			expected: []uint16{0,1,2,3,4,5},
		},
		{
			name: "reversed",
			slice: []uint16{5,4,3,2,1,0},
			expected: []uint16{0,1,2,3,4,5},
		},
		{
			name: "random",
			slice: []uint16{3,1,4,5,0,2},
			expected: []uint16{0,1,2,3,4,5},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UInt16Slice(test.slice).SortAsc().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// SortDesc
func TestUInt16Slice_SortDesc(t *testing.T) {
	tests := []struct {
		name string
		slice []uint16
		expected []uint16
	}{
		{
			name: "empty",
			slice: []uint16{},
			expected: []uint16{},
		},
		{
			name: "already sorted",
			slice: []uint16{5,4,3,2,1,0},
			expected: []uint16{5,4,3,2,1,0},
		},
		{
			name: "reversed",
			slice: []uint16{0,1,2,3,4,5},
			expected: []uint16{5,4,3,2,1,0},
		},
		{
			name: "random",
			slice: []uint16{3,1,4,5,0,2},
			expected: []uint16{5,4,3,2,1,0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UInt16Slice(test.slice).SortDesc().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Unique
func TestUInt16Slice_Unique(t *testing.T) {
	tests := []struct {
		name string
		slice []uint16
		expected []uint16
	}{
		{
			name: "unaffected",
			slice: []uint16{0,1,2,3,4,5},
			expected: []uint16{0,1,2,3,4,5},
		},
		{
			name: "one extra five",
			slice: []uint16{5,0,1,2,3,4,5},
			expected: []uint16{5,0,1,2,3,4},
		},
		{
			name: "extras everywhere",
			slice: []uint16{0,0,1,0,1,2,2,2,3,0,3,4,2,3,4,4,2,1,0},
			expected: []uint16{0,1,2,3,4},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UInt16Slice(test.slice).Unique().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Reverse
func TestUInt16Slice_Reverse(t *testing.T) {
	tests := []struct {
		name string
		slice []uint16
		expected []uint16
	}{
		{
			name: "even length",
			slice: []uint16{0,1,2,3,4,5},
			expected: []uint16{5,4,3,2,1,0},
		},
		{
			name: "odd length",
			slice: []uint16{0,1,2,3,4,5,6},
			expected: []uint16{6,5,4,3,2,1,0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UInt16Slice(test.slice).Reverse().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Filter
func TestUInt16Slice_Filter(t *testing.T) {
	tests := []struct {
		name string
		slice []uint16
		expected []uint16
		filterFunc func(uint16) bool
	}{
		{
			name: "gt 10",
			slice: []uint16{1, 2, 5, 11, 13, 15},
			expected: []uint16{11, 13, 15},
			filterFunc: func(n uint16) bool { return n > 10 },
		},
		{
			name: "mod 3",
			slice: []uint16{1, 2, 6, 11, 12, 15, 17},
			expected: []uint16{6, 12, 15},
			filterFunc: func(n uint16) bool { return n%3 == 0 },
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UInt16Slice(test.slice).Filter(test.filterFunc).Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Each
func TestUInt16Slice_Each(t *testing.T) {

	var rabbit uint16
	tests := []struct {
		name string
		slice []uint16
		expected uint16
		eachFunc func(uint16)
	}{
		{
			name: "add n",
			slice: []uint16{1, 2, 5, 11, 13, 15},
			expected: 47,
			eachFunc: func(n uint16) { rabbit += n },
		},
		{
			name: "subtract n",
			slice: []uint16{1, 2, 6, 8, 12},
			expected: 18,
			eachFunc: func(n uint16) { rabbit -= n },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			UInt16Slice(test.slice).Each(test.eachFunc)
			if test.expected != rabbit {
				t.Errorf("expected %v, got %v", test.expected, rabbit)
			}
		})
	}
}

// Map
func TestUInt16Slice_Map(t *testing.T) {
	tests := []struct {
		name string
		slice []uint16
		expected []uint16
		mapFunc func(uint16) uint16
	}{
		{
			name: "add 3",
			slice: []uint16{1, 2, 5, 11, 13, 15},
			expected: []uint16{4, 5, 8, 14, 16, 18},
			mapFunc: func(n uint16) uint16 { return n + 3 },
		},
		{
			name: "set mod 2",
			slice: []uint16{1, 2, 6, 8, 12, 15, 17},
			expected: []uint16{1, 0, 0, 0, 0, 1, 1},
			mapFunc: func(n uint16) uint16 { return uint16(n%2) },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			UInt16Slice(test.slice).Map(test.mapFunc)
			if !reflect.DeepEqual(test.expected, test.slice) {
				t.Errorf("expected %v, got %v", test.expected, test.slice)
			}
		})
	}
}

// endregion

// region BENCHMARKS
// IndexOf
func BenchmarkUInt16Slice_IndexOf(b *testing.B) {
	benchmarks := []struct {
		name string
		slice []uint16
	}{
		{
			name: "10 elements",
			slice: internal.GenUInt16Slice(10),
		},
		{
			name: "100 elements",
			slice: internal.GenUInt16Slice(100),
		},
		{
			name: "1000 elements",
			slice: internal.GenUInt16Slice(1000),
		},
		{
			name: "10000 elements",
			slice: internal.GenUInt16Slice(10000),
		},
		{
			name: "100000 elements",
			slice: internal.GenUInt16Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UInt16Slice(bm.slice).IndexOf(1)
			}
		})
	}
}

// Contains
func BenchmarkUInt16Slice_Contains(b *testing.B) {
	benchmarks := []struct {
		name string
		slice []uint16
	}{
		{
			name: "10 elements",
			slice: internal.GenUInt16Slice(10),
		},
		{
			name: "100 elements",
			slice: internal.GenUInt16Slice(100),
		},
		{
			name: "1000 elements",
			slice: internal.GenUInt16Slice(1000),
		},
		{
			name: "10000 elements",
			slice: internal.GenUInt16Slice(10000),
		},
		{
			name: "100000 elements",
			slice: internal.GenUInt16Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UInt16Slice(bm.slice).Contains(1)
			}
		})
	}
}

// SortAsc
func BenchmarkUInt16Slice_SortAsc(b *testing.B) {
	benchmarks := []struct {
		name string
		slice []uint16
	}{
		{
			name: "10 elements",
			slice: internal.GenUInt16Slice(10),
		},
		{
			name: "100 elements",
			slice: internal.GenUInt16Slice(100),
		},
		{
			name: "1000 elements",
			slice: internal.GenUInt16Slice(1000),
		},
		{
			name: "10000 elements",
			slice: internal.GenUInt16Slice(10000),
		},
		{
			name: "100000 elements",
			slice: internal.GenUInt16Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UInt16Slice(bm.slice).SortAsc()
			}
		})
	}
}

// SortDesc
func BenchmarkUInt16Slice_SortDesc(b *testing.B) {
	benchmarks := []struct {
		name string
		slice []uint16
	}{
		{
			name: "10 elements",
			slice: internal.GenUInt16Slice(10),
		},
		{
			name: "100 elements",
			slice: internal.GenUInt16Slice(100),
		},
		{
			name: "1000 elements",
			slice: internal.GenUInt16Slice(1000),
		},
		{
			name: "10000 elements",
			slice: internal.GenUInt16Slice(10000),
		},
		{
			name: "100000 elements",
			slice: internal.GenUInt16Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UInt16Slice(bm.slice).SortDesc()
			}
		})
	}
}

// Reverse
func BenchmarkUInt16Slice_Reverse(b *testing.B) {
	benchmarks := []struct {
		name string
		slice []uint16
	}{
		{
			name: "10 elements",
			slice: internal.GenUInt16Slice(10),
		},
		{
			name: "100 elements",
			slice: internal.GenUInt16Slice(100),
		},
		{
			name: "1000 elements",
			slice: internal.GenUInt16Slice(1000),
		},
		{
			name: "10000 elements",
			slice: internal.GenUInt16Slice(10000),
		},
		{
			name: "100000 elements",
			slice: internal.GenUInt16Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UInt16Slice(bm.slice).Reverse()
			}
		})
	}
}

// Filter
func BenchmarkUInt16Slice_Filter(b *testing.B) {
	benchFunc := func(n uint16) bool {
		return n%2 == 0
	}
	benchmarks := []struct {
		name string
		slice []uint16
	}{
		{
			name: "10 elements",
			slice: internal.GenUInt16Slice(10),
		},
		{
			name: "100 elements",
			slice: internal.GenUInt16Slice(100),
		},
		{
			name: "1000 elements",
			slice: internal.GenUInt16Slice(1000),
		},
		{
			name: "10000 elements",
			slice: internal.GenUInt16Slice(10000),
		},
		{
			name: "100000 elements",
			slice: internal.GenUInt16Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UInt16Slice(bm.slice).Filter(benchFunc)
			}
		})
	}
}

// Each
func BenchmarkUInt16Slice_Each(b *testing.B) {
	var rabbit uint16
	benchFunc := func(n uint16) {
		rabbit = n
	}
	benchmarks := []struct {
		name string
		slice []uint16
	}{
		{
			name: "10 elements",
			slice: internal.GenUInt16Slice(10),
		},
		{
			name: "100 elements",
			slice: internal.GenUInt16Slice(100),
		},
		{
			name: "1000 elements",
			slice: internal.GenUInt16Slice(1000),
		},
		{
			name: "10000 elements",
			slice: internal.GenUInt16Slice(10000),
		},
		{
			name: "100000 elements",
			slice: internal.GenUInt16Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UInt16Slice(bm.slice).Each(benchFunc)
			}
		})
	}

	fmt.Fprintf(ioutil.Discard, "%v", rabbit)
}

// Map
func BenchmarkUInt16Slice_Map(b *testing.B) {
	benchFunc := func(n uint16) uint16 {
		n++
		return n
	}
	benchmarks := []struct {
		name string
		slice []uint16
	}{
		{
			name: "10 elements",
			slice: internal.GenUInt16Slice(10),
		},
		{
			name: "100 elements",
			slice: internal.GenUInt16Slice(100),
		},
		{
			name: "1000 elements",
			slice: internal.GenUInt16Slice(1000),
		},
		{
			name: "10000 elements",
			slice: internal.GenUInt16Slice(10000),
		},
		{
			name: "100000 elements",
			slice: internal.GenUInt16Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UInt16Slice(bm.slice).Map(benchFunc)
			}
		})
	}
}
// endregion
