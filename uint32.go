/* generated by github.com/schigh/slices/internal/base.go.  do not edit. */
package slices

import "sort"

// UInt32Slice is the alias of []uint32
type UInt32Slice []uint32

// Value returns the aliased []uint32
func (slice UInt32Slice) Value() []uint32 {
	return []uint32(slice)
}

// IndexOf returns the first index of needle, or -1 if not found
func (slice UInt32Slice) IndexOf(needle uint32) int {
	for idx, n := range slice {
		if needle == n {
			return idx
		}
	}

	return NotInSlice
}

// Contains returns true if the slice contains needle
func (slice UInt32Slice) Contains(needle uint32) bool {
	return slice.IndexOf(needle) != NotInSlice
}

// SortAsc will sort the slice in ascending order
func (slice UInt32Slice) SortAsc() UInt32Slice {
	sort.SliceStable(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	return slice
}

// SortDesc will sort the slice in descending order
func (slice UInt32Slice) SortDesc() UInt32Slice {
	sort.SliceStable(slice, func(i, j int) bool {
		return slice[j] < slice[i]
	})
	return slice
}

// Reverse will reverse the order of the slice
func (slice UInt32Slice) Reverse() UInt32Slice {
	l := len(slice)
	for i, j := 0, l-1; i < l/2; i++ {
		slice[i], slice[j] = slice[j], slice[i]
		j--
	}

	return slice
}

// Unique filters out duplicate uint32 values
func (slice UInt32Slice) Unique() UInt32Slice {
	u := make([]uint32, 0, len(slice))
	m := make(map[uint32]bool)

	for _, v := range slice {
		if _, ok := m[v]; !ok {
			m[v] = true
			u = append(u, v)
		}
	}

	return UInt32Slice(u)
}

// Filter will return all uint32 values that evaluate true in the user-supplied function
func (slice UInt32Slice) Filter(f func(uint32) bool) UInt32Slice {
	out := make([]uint32, 0, len(slice))
	for _, v := range slice {
		if f(v) {
			out = append(out, v)
		}
	}

	return UInt32Slice(out)
}

// Each will apply a function to each uint32 in the slice.
// This should be used with data outside of the slice as it doesn't mutate it
func (slice UInt32Slice) Each(f func(uint32)) {
	for _, v := range slice {
		f(v)
	}
}

// Map will apply a function to each uint32 in the slice and replace the previous value
func (slice UInt32Slice) Map(f func(uint32) uint32) {
	for i, v := range slice {
		slice[i] = f(v)
	}
}
