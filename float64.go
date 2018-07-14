/* generated by github.com/schigh/slices/internal/base.go.  do not edit. */
package slices

import "sort"

// Float64Slice is the alias of []float64
type Float64Slice []float64

// Value returns the aliased []float64
func (slice Float64Slice) Value() []float64 {
	return []float64(slice)
}

// IndexOf returns the first index of needle, or -1 if not found
func (slice Float64Slice) IndexOf(needle float64) int {
	for idx, n := range slice {
		if needle == n {
			return idx
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

	for _, v := range slice {
		if _, ok := m[v]; !ok {
			m[v] = true
			u = append(u, v)
		}
	}

	return Float64Slice(u)
}

// Filter will return all float64 values that evaluate true in the user-supplied function
func (slice Float64Slice) Filter(f func(float64) bool) Float64Slice {
	out := make([]float64, 0, len(slice))
	for _, v := range slice {
		if f(v) {
			out = append(out, v)
		}
	}

	return Float64Slice(out)
}

// Each will apply a function to each float64 in the slice.
// This should be used with data outside of the slice as it doesn't mutate it
func (slice Float64Slice) Each(f func(float64)) {
	for _, v := range slice {
		f(v)
	}
}

// Map will apply a function to each float64 in the slice and replace the previous value
func (slice Float64Slice) Map(f func(float64) float64) {
	for i, v := range slice {
		slice[i] = f(v)
	}
}
