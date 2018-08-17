package slices

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"
)

/* generated by github.com/schigh/slices/base/base.go.  do not edit. */

// region TESTS
// IndexOf
func TestFloat64Slice_IndexOf(t *testing.T) {
	tests := []struct {
		name     string
		slice    []float64
		needle   float64
		expected int
	}{
		{
			name:     "only item",
			slice:    []float64{1},
			needle:   1,
			expected: 0,
		},
		{
			name:     "at index 1",
			slice:    []float64{0, 1.1, 1.1},
			needle:   1.1,
			expected: 1,
		},
		{
			name:     "missing",
			slice:    []float64{1.1, 2.1, 3.1, 4},
			needle:   5.1,
			expected: NotInSlice,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Float64Slice(test.slice).IndexOf(test.needle)
			if result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		})
	}
}

// Contains
func TestFloat64Slice_Contains(t *testing.T) {
	tests := []struct {
		name     string
		slice    []float64
		needle   float64
		expected bool
	}{
		{
			name:     "present",
			slice:    []float64{1.1, 2.0, 3.2, 4.5},
			needle:   4.5,
			expected: true,
		},
		{
			name:     "not present",
			slice:    []float64{1.1, 2.0, 3.2, 4.5},
			needle:   5.6,
			expected: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Float64Slice(test.slice).Contains(test.needle)
			if result != test.expected {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// SortAsc
func TestFloat64Slice_SortAsc(t *testing.T) {
	tests := []struct {
		name     string
		slice    []float64
		expected []float64
	}{
		{
			name:     "empty",
			slice:    []float64{},
			expected: []float64{},
		},
		{
			name:     "already sorted",
			slice:    []float64{0, 0.1, 0.11, 0.111, 0.1111, 0.11111},
			expected: []float64{0, 0.1, 0.11, 0.111, 0.1111, 0.11111},
		},
		{
			name:     "reversed",
			slice:    []float64{0.11111, 0.1111, 0.111, 0.11, 0.1, 0},
			expected: []float64{0, 0.1, 0.11, 0.111, 0.1111, 0.11111},
		},
		{
			name:     "random",
			slice:    []float64{3.0, 1.1, 4.32, 5.4, 0.123, 2.99},
			expected: []float64{0.123, 1.1, 2.99, 3.0, 4.32, 5.4},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Float64Slice(test.slice).SortAsc().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// SortDesc
func TestFloat64Slice_SortDesc(t *testing.T) {
	tests := []struct {
		name     string
		slice    []float64
		expected []float64
	}{
		{
			name:     "empty",
			slice:    []float64{},
			expected: []float64{},
		},
		{
			name:     "already sorted",
			slice:    []float64{0.11111, 0.1111, 0.111, 0.11, 0.1, 0},
			expected: []float64{0.11111, 0.1111, 0.111, 0.11, 0.1, 0},
		},
		{
			name:     "reversed",
			slice:    []float64{0, 0.1, 0.11, 0.111, 0.1111, 0.11111},
			expected: []float64{0.11111, 0.1111, 0.111, 0.11, 0.1, 0},
		},
		{
			name:     "random",
			slice:    []float64{3.0, 1.1, 4.32, 5.4, 0.123, 2.99},
			expected: []float64{5.4, 4.32, 3.0, 2.99, 1.1, 0.123},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Float64Slice(test.slice).SortDesc().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Unique
func TestFloat64Slice_Unique(t *testing.T) {
	tests := []struct {
		name     string
		slice    []float64
		expected []float64
	}{
		{
			name:     "unaffected",
			slice:    []float64{0.123, 0.1, 0.3, 0.99, 0.432, 0.101},
			expected: []float64{0.123, 0.1, 0.3, 0.99, 0.432, 0.101},
		},
		{
			name:     "one extra 0.101",
			slice:    []float64{0.123, 0.1, 0.101, 0.3, 0.99, 0.432, 0.101},
			expected: []float64{0.123, 0.1, 0.101, 0.3, 0.99, 0.432},
		},
		{
			name:     "extras everywhere",
			slice:    []float64{0.1, 0.1, 1.2, 0.1, 1.2, 2.3, 2.3, 2.3, 3.4, 0.1, 3.4, 4.5, 2.3, 3.4, 4.5, 4.5, 2.3, 1.2, 0.1},
			expected: []float64{0.1, 1.2, 2.3, 3.4, 4.5},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Float64Slice(test.slice).Unique().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Reverse
func TestFloat64Slice_Reverse(t *testing.T) {
	tests := []struct {
		name     string
		slice    []float64
		expected []float64
	}{
		{
			name:     "even length",
			slice:    []float64{0.123, 1.234, 2.345, 3.456, 4.567, 5.678},
			expected: []float64{5.678, 4.567, 3.456, 2.345, 1.234, 0.123},
		},
		{
			name:     "odd length",
			slice:    []float64{0.123, 1.234, 2.345, 3.456, 4.567, 5.678, 6.789},
			expected: []float64{6.789, 5.678, 4.567, 3.456, 2.345, 1.234, 0.123},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Float64Slice(test.slice).Reverse().Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Filter
func TestFloat64Slice_Filter(t *testing.T) {
	tests := []struct {
		name       string
		slice      []float64
		expected   []float64
		filterFunc func(float64) bool
	}{
		{
			name:       "gt 10.5",
			slice:      []float64{1.2, 2.3, 5.6, 11.12, 13.14, 15.16},
			expected:   []float64{11.12, 13.14, 15.16},
			filterFunc: func(n float64) bool { return n > 10.5 },
		},
		{
			name:       "gt 0",
			slice:      []float64{0, -0.000001, 0.000001, -0.0000001, 0.0000001},
			expected:   []float64{0.000001, 0.0000001},
			filterFunc: func(n float64) bool { return n > 0 },
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := Float64Slice(test.slice).Filter(test.filterFunc).Value()
			if !reflect.DeepEqual(test.expected, result) {
				t.Errorf("expected %v, got %v", test.expected, result)
			}
		})
	}
}

// Each
func TestFloat64Slice_Each(t *testing.T) {

	var rabbit float64
	tests := []struct {
		name     string
		slice    []float64
		expected float64
		eachFunc func(float64)
	}{
		{
			name:     "add n",
			slice:    []float64{1.25, 2.5, 5.75, 11.25, 13.5, 15.25},
			expected: 49.5,
			eachFunc: func(n float64) { rabbit += n },
		},
		{
			name:     "subtract n",
			slice:    []float64{1.5, 2.25, 6.75, 8.5, 12.25},
			expected: 18.25,
			eachFunc: func(n float64) { rabbit -= n },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			Float64Slice(test.slice).Each(test.eachFunc)
			if test.expected != rabbit {
				t.Errorf("expected %v, got %v", test.expected, rabbit)
			}
		})
	}

	fmt.Fprintf(ioutil.Discard, "%v", rabbit)
}

// CheckEach
func TestFloat64Slice_CheckEach(t *testing.T) {

	var rabbit float64
	myErr := errors.New("i am an error")
	tests := []struct {
		name     string
		slice    []float64
		expected float64
		before   func()
		err      error
		eachFunc func(float64) error
	}{
		{
			name:     "add n",
			slice:    []float64{1.25, 2.5, 5.75, 11.25, 13.5, 15.25},
			expected: 49.5,
			eachFunc: func(n float64) error {
				rabbit += n
				return nil
			},
		},
		{
			name:     "subtract n",
			slice:    []float64{1.5, 2.25, 6.75, 8.5, 12.25},
			expected: 18.25,
			eachFunc: func(n float64) error {
				rabbit -= n
				return nil
			},
		},
		{
			name:     "errors",
			slice:    []float64{1.25, 2.5, 5.75, 11.25, 13.5, 15.25},
			expected: 9.5,
			err:      myErr,
			before:   func() { rabbit = 0 },
			eachFunc: func(n float64) error {
				if rabbit > 9 {
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
			checkErr := Float64Slice(test.slice).CheckEach(test.eachFunc)
			if test.expected != rabbit {
				t.Errorf("expected %v, got %v", test.expected, rabbit)
			}
			if test.err != checkErr {
				t.Errorf("expected %v, got %v", myErr, test.err)
			}
		})
	}

	fmt.Fprintf(ioutil.Discard, "%v", rabbit)
}

// Map
func TestFloat64Slice_Map(t *testing.T) {
	tests := []struct {
		name     string
		slice    []float64
		expected []float64
		mapFunc  func(float64) float64
	}{
		{
			name:     "add 3.5",
			slice:    []float64{1.23, 2.34, 5.67, 11.5, 13.25, 15.25},
			expected: []float64{4.73, 5.84, 9.17, 15, 16.75, 18.75},
			mapFunc:  func(n float64) float64 { return n + 3.5 },
		},
		{
			name:     "multiply by 2",
			slice:    []float64{1.5, 2.5, 6.5, 8.5, 12.5, 15.5, 17.5},
			expected: []float64{3.0, 5.0, 13.0, 17.0, 25.0, 31.0, 35.0},
			mapFunc:  func(n float64) float64 { return n * 2 },
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			Float64Slice(test.slice).Map(test.mapFunc)
			if !reflect.DeepEqual(test.expected, test.slice) {
				t.Errorf("expected %v, got %v", test.expected, test.slice)
			}
		})
	}
}

//endregion
