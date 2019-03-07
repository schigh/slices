package slices

import "sort"

/* generated by github.com/schigh/slices/base/base.go.  do not edit. */

// Float64Slice is the alias of []float64
type Float64Slice []float64

// Value returns the aliased []float64
func (slice Float64Slice) Value() []float64 {
	return []float64(slice)
}

// IndexOf returns the first index of needle, or -1 if not found
func (slice Float64Slice) IndexOf(needle float64) int {
	for i := 0; i < len(slice); i++ {
		if needle == slice[i] {
			return i
		}
	}

	return NotInSlice
}

// Contains returns true if the slice contains needle
func (slice Float64Slice) Contains(needle float64) bool {
	return slice.IndexOf(needle) != NotInSlice
}

// SortAsc will sort the slice in ascending order
func (slice Float64Slice) SortAsc() Float64Slice {
	sort.SliceStable(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	return slice
}

// SortDesc will sort the slice in descending order
func (slice Float64Slice) SortDesc() Float64Slice {
	sort.SliceStable(slice, func(i, j int) bool {
		return slice[j] < slice[i]
	})
	return slice
}

// Reverse will reverse the order of the slice
func (slice Float64Slice) Reverse() Float64Slice {
	l := len(slice)
	for i, j := 0, l-1; i < l/2; i++ {
		slice[i], slice[j] = slice[j], slice[i]
		j--
	}

	return slice
}

// Unique filters out duplicate float64 values
func (slice Float64Slice) Unique() Float64Slice {
	u := make([]float64, 0, len(slice))
	m := make(map[float64]bool)

	for i := 0; i < len(slice); i++ {
		if _, ok := m[slice[i]]; !ok {
			m[slice[i]] = true
			u = append(u, slice[i])
		}
	}

	return Float64Slice(u)
}

// Filter will return all float64 values that evaluate true in the user-supplied function
func (slice Float64Slice) Filter(f func(float64) bool) Float64Slice {
	out := make([]float64, 0, len(slice))
	for i := 0; i < len(slice); i++ {
		if f(slice[i]) {
			out = append(out, slice[i])
		}
	}

	return Float64Slice(out)
}

// Each will apply a function to each float64 in the slice.
// This function will iterate over the slice completely.  No
// items in the slice should be mutated by this operation.
func (slice Float64Slice) Each(f func(float64)) {
	for i := 0; i < len(slice); i++ {
		f(slice[i])
	}
}

// TryEach will apply a function to each float64 in the slice.
// If the function returns an error, the iteration will stop and return
// the index of the element that caused the function to return the error.
// The second returned value will be first error returned from the
// supplied function, and nil otherwise.
// No items in the slice should be mutated by this operation.
func (slice Float64Slice) TryEach(f func(float64) error) (int, error) {
	for i := 0; i < len(slice); i++ {
		if err := f(slice[i]); err != nil {
			return i, err
		}
	}

	return NotInSlice, nil
}

// IfEach will apply a function to each float64 in the slice.
// If the function returns false, the iteration will stop and return
// the index of the element that caused the function to return false.
// The second returned value will be true if all members of the slice
// cause the provided function to return true, and false otherwise.
// No items in the slice should be mutated by this operation.
func (slice Float64Slice) IfEach(f func(float64) bool) (int, bool) {
	for i := 0; i < len(slice); i++ {
		if !f(slice[i]) {
			return i, false
		}
	}

	return NotInSlice, true
}

// Map will apply a function to each float64 in the slice and replace the previous value
func (slice Float64Slice) Map(f func(float64) float64) {
	for i := 0; i < len(slice); i++ {
		slice[i] = f(slice[i])
	}
}

// Chunk will divide the slice of float64 into smaller slices defined by chunk length
func (slice Float64Slice) Chunk(size int) [][]float64 {
	l := len(slice)
	if l == 0 || size <= 0 {
		return make([][]float64, 0)
	}

	floor := l / size
	out := make([][]float64, 0, floor+1)
	var k int

	for i := 0; i < floor; i++ {
		k = i*size + size
		out = append(out, slice[i*size:k])
	}
	if l > k {
		out = append(out, slice[k:])
	}

	return out
}
