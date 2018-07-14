// generated by slices (github.com/schigh/slices).  do not edit.
// gen date: Sat, 14 Jul 2018 00:30:20 -0400

package examples

// AccountSlice aliases []Account
type AccountSlice []Account

// Value returns the wrapped Account slice
func (slice AccountSlice) Value() []Account {
	return []Account(slice)
}

// Map applies a function to every Account in the slice.  This function will mutate the slice in place
func (slice AccountSlice) Map(f func(Account) Account) {
	for i, v := range slice {
		slice[i] = f(v)
	}
}

// Filter evaluates every element in the slice, and returns all Account
// instances where the eval function returns true
func (slice AccountSlice) Filter(f func(Account) bool) AccountSlice {
	out := make([]Account, 0, len(slice))
	for _, v := range slice {
		if f(v) {
			out = append(out, v)
		}
	}

	return AccountSlice(out)
}

// Each applies a function to every Account in the slice.
func (slice AccountSlice) Each(f func(Account)) {
	for _, v := range slice {
		f(v)
	}
}
