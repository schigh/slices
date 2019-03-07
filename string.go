package slices

import "sort"

/* generated by github.com/schigh/slices/base/base.go.  do not edit. */

// StringSlice is the alias of []string
type StringSlice []string

// Value returns the aliased []string
func (slice StringSlice) Value() []string {
	return []string(slice)
}

// IndexOf returns the first index of needle, or -1 if not found
func (slice StringSlice) IndexOf(needle string) int {
	for i := 0; i < len(slice); i++ {
		if needle == slice[i] {
			return i
		}
	}

	return NotInSlice
}

// Contains returns true if the slice contains needle
func (slice StringSlice) Contains(needle string) bool {
	return slice.IndexOf(needle) != NotInSlice
}

// SortAsc will sort the slice in ascending order
func (slice StringSlice) SortAsc() StringSlice {
	sort.SliceStable(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	return slice
}

// SortDesc will sort the slice in descending order
func (slice StringSlice) SortDesc() StringSlice {
	sort.SliceStable(slice, func(i, j int) bool {
		return slice[j] < slice[i]
	})
	return slice
}

// Reverse will reverse the order of the slice
func (slice StringSlice) Reverse() StringSlice {
	l := len(slice)
	for i, j := 0, l-1; i < l/2; i++ {
		slice[i], slice[j] = slice[j], slice[i]
		j--
	}

	return slice
}

// Unique filters out duplicate string values
func (slice StringSlice) Unique() StringSlice {
	u := make([]string, 0, len(slice))
	m := make(map[string]bool)

	for i := 0; i < len(slice); i++ {
		if _, ok := m[slice[i]]; !ok {
			m[slice[i]] = true
			u = append(u, slice[i])
		}
	}

	return StringSlice(u)
}

// Filter will return all string values that evaluate true in the user-supplied function
func (slice StringSlice) Filter(f func(string) bool) StringSlice {
	out := make([]string, 0, len(slice))
	for i := 0; i < len(slice); i++ {
		if f(slice[i]) {
			out = append(out, slice[i])
		}
	}

	return StringSlice(out)
}

// Each will apply a function to each string in the slice.
// This function will iterate over the slice completely.  No
// items in the slice should be mutated by this operation.
func (slice StringSlice) Each(f func(string)) {
	for i := 0; i < len(slice); i++ {
		f(slice[i])
	}
}

// TryEach will apply a function to each string in the slice.
// If the function returns an error, the iteration will stop and return
// the index of the element that caused the function to return the error.
// The second returned value will be first error returned from the
// supplied function, and nil otherwise.
// No items in the slice should be mutated by this operation.
func (slice StringSlice) TryEach(f func(string) error) (int, error) {
	for i := 0; i < len(slice); i++ {
		if err := f(slice[i]); err != nil {
			return i, err
		}
	}

	return NotInSlice, nil
}

// IfEach will apply a function to each string in the slice.
// If the function returns false, the iteration will stop and return
// the index of the element that caused the function to return false.
// The second returned value will be true if all members of the slice
// cause the provided function to return true, and false otherwise.
// No items in the slice should be mutated by this operation.
func (slice StringSlice) IfEach(f func(string) bool) (int, bool) {
	for i := 0; i < len(slice); i++ {
		if !f(slice[i]) {
			return i, false
		}
	}

	return NotInSlice, true
}

// Map will apply a function to each string in the slice and replace the previous value
func (slice StringSlice) Map(f func(string) string) {
	for i := 0; i < len(slice); i++ {
		slice[i] = f(slice[i])
	}
}

// Chunk will divide the slice of string into smaller slices defined by chunk length
func (slice StringSlice) Chunk(size int) [][]string {
	l := len(slice)
	if l == 0 || size <= 0 {
		return make([][]string, 0)
	}

	floor := l / size
	out := make([][]string, 0, floor+1)
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
