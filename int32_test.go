// generated by github.com/schigh/slices/internal/base.go.  do not edit.
package slices

import (
	"testing"
	"reflect"
	"github.com/schigh/slices/internal"
)

//region TESTS
func TestInt32Slice_IndexOf(t *testing.T) {
	tests := []struct {
		name string
		slice []int32
		needle int32
		expected int
	}{
		{
			name: "only item",
			slice: []int32{1},
			needle: 1,
			expected: 0,
		},
		{
			name: "at index 1",
			slice: []int32{0,1,1},
			needle: 1,
			expected: 1,
		},
		{
			name: "missing",
			slice: []int32{1,2,3,4},
			needle: 5,
			expected: NotInSlice,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Int32Slice(test.slice).IndexOf(test.needle)
			if result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		})
	}
}

func TestInt32Slice_Contains(t *testing.T) {
	tests := []struct {
		name string
		slice []int32
		needle int32
		expected bool
	}{
		{
			name: "present",
			slice: []int32{1,2,3,4},
			needle: 4,
			expected: true,
		},
		{
			name: "not present",
			slice: []int32{1,2,3,4},
			needle: 5,
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Int32Slice(test.slice).Contains(test.needle)
			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestInt32Slice_TruncateZero(t *testing.T) {
	tests := []struct {
		name string
		slice []int32
		expected []int32
	}{
		{
			name: "unaffected",
			slice: []int32{1,2,3,4,5},
			expected: []int32{1,2,3,4,5},
		},
		{
			name: "unaffected 2",
			slice: []int32{0,0,0,0,1},
			expected: []int32{0,0,0,0,1},
		},
		{
			name: "all but one",
			slice: []int32{1,0,0,0,0},
			expected: []int32{1},
		},
		{
			name: "last one only",
			slice: []int32{1,2,3,4,0},
			expected: []int32{1,2,3,4},
		},
		{
			name: "empties the slice",
			slice: []int32{0,0,0,0,0},
			expected: []int32{},
		},
		{
			name: "empty slice",
			slice: []int32{},
			expected: []int32{},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Int32Slice(test.slice).TruncateZero().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestInt32Slice_SortAsc(t *testing.T) {
	tests := []struct {
		name string
		slice []int32
		expected []int32
	}{
		{
			name: "empty",
			slice: []int32{},
			expected: []int32{},
		},
		{
			name: "already sorted",
			slice: []int32{0,1,2,3,4,5},
			expected: []int32{0,1,2,3,4,5},
		},
		{
			name: "reversed",
			slice: []int32{5,4,3,2,1,0},
			expected: []int32{0,1,2,3,4,5},
		},
		{
			name: "random",
			slice: []int32{3,1,4,5,0,2},
			expected: []int32{0,1,2,3,4,5},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Int32Slice(test.slice).SortAsc().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestInt32Slice_SortDesc(t *testing.T) {
	tests := []struct {
		name string
		slice []int32
		expected []int32
	}{
		{
			name: "empty",
			slice: []int32{},
			expected: []int32{},
		},
		{
			name: "already sorted",
			slice: []int32{5,4,3,2,1,0},
			expected: []int32{5,4,3,2,1,0},
		},
		{
			name: "reversed",
			slice: []int32{0,1,2,3,4,5},
			expected: []int32{5,4,3,2,1,0},
		},
		{
			name: "random",
			slice: []int32{3,1,4,5,0,2},
			expected: []int32{5,4,3,2,1,0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Int32Slice(test.slice).SortDesc().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestInt32Slice_Unique(t *testing.T) {
	tests := []struct {
		name string
		slice []int32
		expected []int32
	}{
		{
			name: "unaffected",
			slice: []int32{0,1,2,3,4,5},
			expected: []int32{0,1,2,3,4,5},
		},
		{
			name: "one extra five",
			slice: []int32{5,0,1,2,3,4,5},
			expected: []int32{5,0,1,2,3,4},
		},
		{
			name: "extras everywhere",
			slice: []int32{0,0,1,0,1,2,2,2,3,0,3,4,2,3,4,4,2,1,0},
			expected: []int32{0,1,2,3,4},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Int32Slice(test.slice).Unique().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestInt32Slice_Reverse(t *testing.T) {
	tests := []struct {
		name string
		slice []int32
		expected []int32
	}{
		{
			name: "even length",
			slice: []int32{0,1,2,3,4,5},
			expected: []int32{5,4,3,2,1,0},
		},
		{
			name: "odd length",
			slice: []int32{0,1,2,3,4,5,6},
			expected: []int32{6,5,4,3,2,1,0},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Int32Slice(test.slice).Reverse().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

//endregion

//region BENCHMARKS

func BenchmarkInt32Slice_IndexOf(b *testing.B) {
	benchmarks := []struct {
		name string
		slice []int32
	}{
		{
			name: "10 elements",
			slice: internal.GenInt32Slice(10),
		},
		{
			name: "100 elements",
			slice: internal.GenInt32Slice(100),
		},
		{
			name: "1000 elements",
			slice: internal.GenInt32Slice(1000),
		},
		{
			name: "10000 elements",
			slice: internal.GenInt32Slice(10000),
		},
		{
			name: "100000 elements",
			slice: internal.GenInt32Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int32Slice(bm.slice).IndexOf(1)
			}
		})
	}
}

func BenchmarkInt32Slice_Contains(b *testing.B) {
	benchmarks := []struct {
		name string
		slice []int32
	}{
		{
			name: "10 elements",
			slice: internal.GenInt32Slice(10),
		},
		{
			name: "100 elements",
			slice: internal.GenInt32Slice(100),
		},
		{
			name: "1000 elements",
			slice: internal.GenInt32Slice(1000),
		},
		{
			name: "10000 elements",
			slice: internal.GenInt32Slice(10000),
		},
		{
			name: "100000 elements",
			slice: internal.GenInt32Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int32Slice(bm.slice).Contains(1)
			}
		})
	}
}

func BenchmarkInt32Slice_SortAsc(b *testing.B) {
	benchmarks := []struct {
		name string
		slice []int32
	}{
		{
			name: "10 elements",
			slice: internal.GenInt32Slice(10),
		},
		{
			name: "100 elements",
			slice: internal.GenInt32Slice(100),
		},
		{
			name: "1000 elements",
			slice: internal.GenInt32Slice(1000),
		},
		{
			name: "10000 elements",
			slice: internal.GenInt32Slice(10000),
		},
		{
			name: "100000 elements",
			slice: internal.GenInt32Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int32Slice(bm.slice).SortAsc()
			}
		})
	}
}

func BenchmarkInt32Slice_SortDesc(b *testing.B) {
	benchmarks := []struct {
		name string
		slice []int32
	}{
		{
			name: "10 elements",
			slice: internal.GenInt32Slice(10),
		},
		{
			name: "100 elements",
			slice: internal.GenInt32Slice(100),
		},
		{
			name: "1000 elements",
			slice: internal.GenInt32Slice(1000),
		},
		{
			name: "10000 elements",
			slice: internal.GenInt32Slice(10000),
		},
		{
			name: "100000 elements",
			slice: internal.GenInt32Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int32Slice(bm.slice).SortDesc()
			}
		})
	}
}

func BenchmarkInt32Slice_Reverse(b *testing.B) {
	benchmarks := []struct {
		name string
		slice []int32
	}{
		{
			name: "10 elements",
			slice: internal.GenInt32Slice(10),
		},
		{
			name: "100 elements",
			slice: internal.GenInt32Slice(100),
		},
		{
			name: "1000 elements",
			slice: internal.GenInt32Slice(1000),
		},
		{
			name: "10000 elements",
			slice: internal.GenInt32Slice(10000),
		},
		{
			name: "100000 elements",
			slice: internal.GenInt32Slice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				Int32Slice(bm.slice).Reverse()
			}
		})
	}
}
//endregion
