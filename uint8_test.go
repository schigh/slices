// generated by github.com/schigh/slices/internal/base.go.  do not edit.
package slices

import (
	"testing"
	"reflect"
	"github.com/schigh/slices/internal"
)

//region TESTS
func TestUInt8Slice_IndexOf(t *testing.T) {
	tests := []struct {
		name string
		slice []uint8
		needle uint8
		expected int
	}{
		{
			name: "only item",
			slice: []uint8{1},
			needle: 1,
			expected: 0,
		},
		{
			name: "at index 1",
			slice: []uint8{0,1,1},
			needle: 1,
			expected: 1,
		},
		{
			name: "missing",
			slice: []uint8{1,2,3,4},
			needle: 5,
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

func TestUInt8Slice_Contains(t *testing.T) {
	tests := []struct {
		name string
		slice []uint8
		needle uint8
		expected bool
	}{
		{
			name: "present",
			slice: []uint8{1,2,3,4},
			needle: 4,
			expected: true,
		},
		{
			name: "not present",
			slice: []uint8{1,2,3,4},
			needle: 5,
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

func TestUInt8Slice_SortAsc(t *testing.T) {
	tests := []struct {
		name string
		slice []uint8
		expected []uint8
	}{
		{
			name: "empty",
			slice: []uint8{},
			expected: []uint8{},
		},
		{
			name: "already sorted",
			slice: []uint8{0,1,2,3,4,5},
			expected: []uint8{0,1,2,3,4,5},
		},
		{
			name: "reversed",
			slice: []uint8{5,4,3,2,1,0},
			expected: []uint8{0,1,2,3,4,5},
		},
		{
			name: "random",
			slice: []uint8{3,1,4,5,0,2},
			expected: []uint8{0,1,2,3,4,5},
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

func TestUInt8Slice_SortDesc(t *testing.T) {
	tests := []struct {
		name string
		slice []uint8
		expected []uint8
	}{
		{
			name: "empty",
			slice: []uint8{},
			expected: []uint8{},
		},
		{
			name: "already sorted",
			slice: []uint8{5,4,3,2,1,0},
			expected: []uint8{5,4,3,2,1,0},
		},
		{
			name: "reversed",
			slice: []uint8{0,1,2,3,4,5},
			expected: []uint8{5,4,3,2,1,0},
		},
		{
			name: "random",
			slice: []uint8{3,1,4,5,0,2},
			expected: []uint8{5,4,3,2,1,0},
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

func TestUInt8Slice_Unique(t *testing.T) {
	tests := []struct {
		name string
		slice []uint8
		expected []uint8
	}{
		{
			name: "unaffected",
			slice: []uint8{0,1,2,3,4,5},
			expected: []uint8{0,1,2,3,4,5},
		},
		{
			name: "one extra five",
			slice: []uint8{5,0,1,2,3,4,5},
			expected: []uint8{5,0,1,2,3,4},
		},
		{
			name: "extras everywhere",
			slice: []uint8{0,0,1,0,1,2,2,2,3,0,3,4,2,3,4,4,2,1,0},
			expected: []uint8{0,1,2,3,4},
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

func TestUInt8Slice_Reverse(t *testing.T) {
	tests := []struct {
		name string
		slice []uint8
		expected []uint8
	}{
		{
			name: "even length",
			slice: []uint8{0,1,2,3,4,5},
			expected: []uint8{5,4,3,2,1,0},
		},
		{
			name: "odd length",
			slice: []uint8{0,1,2,3,4,5,6},
			expected: []uint8{6,5,4,3,2,1,0},
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

func TestUInt8Slice_Filter(t *testing.T) {
	tests := []struct {
		name string
		slice []uint8
		expected []uint8
		filterFunc func(uint8) bool
	}{
		{
			name: "gt 10",
			slice: []uint8{1, 2, 5, 11, 13, 15},
			expected: []uint8{11, 13, 15},
			filterFunc: func(n uint8) bool { return n > 10 },
		},
		{
			name: "mod 3",
			slice: []uint8{1, 2, 6, 11, 12, 15, 17},
			expected: []uint8{6, 12, 15},
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

//endregion

//region BENCHMARKS

func BenchmarkUInt8Slice_IndexOf(b *testing.B) {
	benchmarks := []struct {
		name string
		slice []uint8
	}{
		{
			name: "10 elements",
			slice: internal.GenUInt8Slice(10),
		},
		{
			name: "100 elements",
			slice: internal.GenUInt8Slice(100),
		},
		{
			name: "1000 elements",
			slice: internal.GenUInt8Slice(1000),
		},
		{
			name: "10000 elements",
			slice: internal.GenUInt8Slice(10000),
		},
		{
			name: "100000 elements",
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

func BenchmarkUInt8Slice_Contains(b *testing.B) {
	benchmarks := []struct {
		name string
		slice []uint8
	}{
		{
			name: "10 elements",
			slice: internal.GenUInt8Slice(10),
		},
		{
			name: "100 elements",
			slice: internal.GenUInt8Slice(100),
		},
		{
			name: "1000 elements",
			slice: internal.GenUInt8Slice(1000),
		},
		{
			name: "10000 elements",
			slice: internal.GenUInt8Slice(10000),
		},
		{
			name: "100000 elements",
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

func BenchmarkUInt8Slice_SortAsc(b *testing.B) {
	benchmarks := []struct {
		name string
		slice []uint8
	}{
		{
			name: "10 elements",
			slice: internal.GenUInt8Slice(10),
		},
		{
			name: "100 elements",
			slice: internal.GenUInt8Slice(100),
		},
		{
			name: "1000 elements",
			slice: internal.GenUInt8Slice(1000),
		},
		{
			name: "10000 elements",
			slice: internal.GenUInt8Slice(10000),
		},
		{
			name: "100000 elements",
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

func BenchmarkUInt8Slice_SortDesc(b *testing.B) {
	benchmarks := []struct {
		name string
		slice []uint8
	}{
		{
			name: "10 elements",
			slice: internal.GenUInt8Slice(10),
		},
		{
			name: "100 elements",
			slice: internal.GenUInt8Slice(100),
		},
		{
			name: "1000 elements",
			slice: internal.GenUInt8Slice(1000),
		},
		{
			name: "10000 elements",
			slice: internal.GenUInt8Slice(10000),
		},
		{
			name: "100000 elements",
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

func BenchmarkUInt8Slice_Reverse(b *testing.B) {
	benchmarks := []struct {
		name string
		slice []uint8
	}{
		{
			name: "10 elements",
			slice: internal.GenUInt8Slice(10),
		},
		{
			name: "100 elements",
			slice: internal.GenUInt8Slice(100),
		},
		{
			name: "1000 elements",
			slice: internal.GenUInt8Slice(1000),
		},
		{
			name: "10000 elements",
			slice: internal.GenUInt8Slice(10000),
		},
		{
			name: "100000 elements",
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

func BenchmarkUInt8Slice_Filter(b *testing.B) {
	benchFunc := func(n uint8) bool {
		return n%2 == 0
	}
	benchmarks := []struct {
		name string
		slice []uint8
	}{
		{
			name: "10 elements",
			slice: internal.GenUInt8Slice(10),
		},
		{
			name: "100 elements",
			slice: internal.GenUInt8Slice(100),
		},
		{
			name: "1000 elements",
			slice: internal.GenUInt8Slice(1000),
		},
		{
			name: "10000 elements",
			slice: internal.GenUInt8Slice(10000),
		},
		{
			name: "100000 elements",
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
//endregion
