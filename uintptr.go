// generated by github.com/schigh/slices/internal/base.go.  do not edit.
package slices

import "sort"

// UIntPtrSlice is the alias of []uintptr
type UIntPtrSlice []uintptr

// Value returns the aliased []uintptr
func (slice UIntPtrSlice) Value() []uintptr {
	return []uintptr(slice)
}

// IndexOf returns the first index of needle, or -1 if not found
func (slice UIntPtrSlice) IndexOf(needle uintptr) int {
	for idx, n := range slice {
		if needle == n {
			return idx
		}
	}

	return NotInSlice
}

// Contains returns true if the slice contains needle
func (slice UIntPtrSlice) Contains(needle uintptr) bool {
	return slice.IndexOf(needle) != NotInSlice
}

// SortAsc will sort an []uintptr in ascending order
func (slice UIntPtrSlice) SortAsc() UIntPtrSlice {
	sort.SliceStable(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	return slice
}

// SortDesc will sort an []uintptr in descending order
func (slice UIntPtrSlice) SortDesc() UIntPtrSlice {
	sort.SliceStable(slice, func(i, j int) bool {
		return slice[j] < slice[i]
	})
	return slice
}

// Reverse will reverse the order of the slice
func (slice UIntPtrSlice) Reverse() UIntPtrSlice {
	l := len(slice)
	for i, j := 0, l-1; i < l/2; i++ {
		slice[i], slice[j] = slice[j], slice[i]
		j--
	}

	return slice
}

// TruncateZero will trim ALL the zero values off the end of the slice,
// stopping at the first non-zero value it finds
func (slice UIntPtrSlice) TruncateZero() UIntPtrSlice {
	l := len(slice)
	if l > 0 {
		for i := l-1; i >= 0; i-- {
			if slice[i] != 0 {
				return slice[:i+1]
			}
		}
		return slice[:0]
	}

	return slice
}

// Unique filters out duplicate values
func (slice UIntPtrSlice) Unique() UIntPtrSlice {
	u := make([]uintptr, 0, len(slice))
	m := make(map[uintptr]bool)

	for _, i := range slice {
		if _, ok := m[i]; !ok {
			m[i] = true
			u = append(u, i)
		}
	}

	return UIntPtrSlice(u)
}

func (slice UIntPtrSlice) Filter(f func(uintptr) bool) UIntPtrSlice {
	out := make([]uintptr, 0, len(slice))
	for _, i := range slice {
		if f(i) {
			out = append(out, i)
		}
	}

	return UIntPtrSlice(out)
}
