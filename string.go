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
	for idx, n := range slice {
		if needle == n {
			return idx
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

	for _, v := range slice {
		if _, ok := m[v]; !ok {
			m[v] = true
			u = append(u, v)
		}
	}

	return StringSlice(u)
}

// Filter will return all string values that evaluate true in the user-supplied function
func (slice StringSlice) Filter(f func(string) bool) StringSlice {
	out := make([]string, 0, len(slice))
	for _, v := range slice {
		if f(v) {
			out = append(out, v)
		}
	}

	return StringSlice(out)
}

// Each will apply a function to each string in the slice.
// This should be used with data outside of the slice as it doesn't mutate it
func (slice StringSlice) Each(f func(string)) {
	for _, v := range slice {
		f(v)
	}
}

// Map will apply a function to each string in the slice and replace the previous value
func (slice StringSlice) Map(f func(string) string) {
	for i, v := range slice {
		slice[i] = f(v)
	}
}
