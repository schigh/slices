// generated by github.com/schigh/slices/internal/base.go.  do not edit.
package slices

import (
	"testing"
	"reflect"
	"github.com/schigh/slices/internal"
)

//region TESTS

func TestStringSlice_IndexOf(t *testing.T) {
	tests := []struct {
		name string
		slice []string
		needle string
		expected int
	}{
		{
			name: "only item",
			slice: []string{"foo"},
			needle: "foo",
			expected: 0,
		},
		{
			name: "at index 1",
			slice: []string{"foo","bar","bar"},
			needle: "bar",
			expected: 1,
		},
		{
			name: "missing",
			slice: []string{"foo","bar","fizz","buzz"},
			needle: "baz",
			expected: NotInSlice,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := StringSlice(test.slice).IndexOf(test.needle)
			if result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		})
	}
}

func TestStringSlice_Contains(t *testing.T) {
	tests := []struct {
		name string
		slice []string
		needle string
		expected bool
	}{
		{
			name: "present",
			slice: []string{"foo","bar","fizz","buzz"},
			needle: "buzz",
			expected: true,
		},
		{
			name: "not present",
			slice: []string{"foo","bar","fizz","buzz"},
			needle: "baz",
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := StringSlice(test.slice).Contains(test.needle)
			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestStringSlice_SortAsc(t *testing.T) {
	tests := []struct {
		name string
		slice []string
		expected []string
	}{
		{
			name: "empty",
			slice: []string{},
			expected: []string{},
		},
		{
			name: "already sorted",
			slice: []string{"bar","baz","buzz","fizz","foo"},
			expected: []string{"bar","baz","buzz","fizz","foo"},
		},
		{
			name: "reversed",
			slice: []string{"foo","fizz","buzz","baz","bar"},
			expected: []string{"bar","baz","buzz","fizz","foo"},
		},
		{
			name: "random",
			slice: []string{"foo","bar","fizz","buzz","baz"},
			expected: []string{"bar","baz","buzz","fizz","foo"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := StringSlice(test.slice).SortAsc().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestStringSlice_SortDesc(t *testing.T) {
	tests := []struct {
		name string
		slice []string
		expected []string
	}{
		{
			name: "empty",
			slice: []string{},
			expected: []string{},
		},
		{
			name: "already sorted",
			slice: []string{"foo","fizz","buzz","baz","bar"},
			expected: []string{"foo","fizz","buzz","baz","bar"},
		},
		{
			name: "reversed",
			slice: []string{"bar","baz","buzz","fizz","foo"},
			expected: []string{"foo","fizz","buzz","baz","bar"},
		},
		{
			name: "random",
			slice: []string{"foo","bar","fizz","buzz","baz"},
			expected: []string{"foo","fizz","buzz","baz","bar"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := StringSlice(test.slice).SortDesc().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestStringSlice_Unique(t *testing.T) {
	tests := []struct {
		name string
		slice []string
		expected []string
	}{
		{
			name: "unaffected",
			slice: []string{"foo","bar","fizz","buzz","baz"},
			expected: []string{"foo","bar","fizz","buzz","baz"},
		},
		{
			name: "one extra buzz",
			slice: []string{"foo","bar","fizz","buzz","baz","buzz"},
			expected: []string{"foo","bar","fizz","buzz","baz"},
		},
		{
			name: "extras everywhere",
			slice: []string{"foo","foo","bar","foo","bar","fizz","fizz","fizz","buzz","foo","buzz","baz","fizz","buzz","baz","baz","fizz","bar","foo"},
			expected: []string{"foo","bar","fizz","buzz","baz"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := StringSlice(test.slice).Unique().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

func TestStringSlice_Reverse(t *testing.T) {
	tests := []struct {
		name string
		slice []string
		expected []string
	}{
		{
			name: "even length",
			slice: []string{"foo","bar","baz","fizz","buzz"},
			expected: []string{"buzz","fizz","baz","bar","foo"},
		},
		{
			name: "odd length",
			slice: []string{"foo","bar","baz","fizz","buzz","herp"},
			expected: []string{"herp","buzz","fizz","baz","bar","foo"},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := StringSlice(test.slice).Reverse().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

//endregion

//region BENCHMARKS

func BenchmarkStringSlice_IndexOf(b *testing.B) {
	benchmarks := []struct {
		name string
		slice []string
	}{
		{
			name: "10 elements",
			slice: internal.GenStringSlice(10),
		},
		{
			name: "100 elements",
			slice: internal.GenStringSlice(100),
		},
		{
			name: "1000 elements",
			slice: internal.GenStringSlice(1000),
		},
		{
			name: "10000 elements",
			slice: internal.GenStringSlice(10000),
		},
		{
			name: "100000 elements",
			slice: internal.GenStringSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				StringSlice(bm.slice).IndexOf("foo")
			}
		})
	}
}

func BenchmarkStringSlice_Contains(b *testing.B) {
	benchmarks := []struct {
		name string
		slice []string
	}{
		{
			name: "10 elements",
			slice: internal.GenStringSlice(10),
		},
		{
			name: "100 elements",
			slice: internal.GenStringSlice(100),
		},
		{
			name: "1000 elements",
			slice: internal.GenStringSlice(1000),
		},
		{
			name: "10000 elements",
			slice: internal.GenStringSlice(10000),
		},
		{
			name: "100000 elements",
			slice: internal.GenStringSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				StringSlice(bm.slice).Contains("foo")
			}
		})
	}
}

func BenchmarkStringSlice_SortAsc(b *testing.B) {
	benchmarks := []struct {
		name string
		slice []string
	}{
		{
			name: "10 elements",
			slice: internal.GenStringSlice(10),
		},
		{
			name: "100 elements",
			slice: internal.GenStringSlice(100),
		},
		{
			name: "1000 elements",
			slice: internal.GenStringSlice(1000),
		},
		{
			name: "10000 elements",
			slice: internal.GenStringSlice(10000),
		},
		{
			name: "100000 elements",
			slice: internal.GenStringSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				StringSlice(bm.slice).SortAsc()
			}
		})
	}
}

func BenchmarkStringSlice_SortDesc(b *testing.B) {
	benchmarks := []struct {
		name string
		slice []string
	}{
		{
			name: "10 elements",
			slice: internal.GenStringSlice(10),
		},
		{
			name: "100 elements",
			slice: internal.GenStringSlice(100),
		},
		{
			name: "1000 elements",
			slice: internal.GenStringSlice(1000),
		},
		{
			name: "10000 elements",
			slice: internal.GenStringSlice(10000),
		},
		{
			name: "100000 elements",
			slice: internal.GenStringSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				StringSlice(bm.slice).SortDesc()
			}
		})
	}
}

func BenchmarkStringSlice_Reverse(b *testing.B) {
	benchmarks := []struct {
		name string
		slice []string
	}{
		{
			name: "10 elements",
			slice: internal.GenStringSlice(10),
		},
		{
			name: "100 elements",
			slice: internal.GenStringSlice(100),
		},
		{
			name: "1000 elements",
			slice: internal.GenStringSlice(1000),
		},
		{
			name: "10000 elements",
			slice: internal.GenStringSlice(10000),
		},
		{
			name: "100000 elements",
			slice: internal.GenStringSlice(100000),
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				StringSlice(bm.slice).Reverse()
			}
		})
	}
}
//endregion
