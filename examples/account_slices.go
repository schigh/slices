package examples

/* generated by slices (github.com/schigh/slices).  do not edit. */
/* gen date: Wed, 05 Sep 2018 18:06:13 -0400 */

// AccountSlice aliases []Account
type AccountSlice []Account

// Value returns the wrapped Account slice
func (slice AccountSlice) Value() []Account {
	return []Account(slice)
}

// Map applies a function to every Account in the slice.  This function will mutate the slice in place
func (slice AccountSlice) Map(f func(Account) Account) {
	for i := 0; i < len(slice); i++ {
		slice[i] = f(slice[i])
	}
}

// Filter evaluates every element in the slice, and returns all Account
// instances where the eval function returns true
func (slice AccountSlice) Filter(f func(Account) bool) AccountSlice {
	out := make([]Account, 0, len(slice))
	for i := 0; i < len(slice); i++ {
		if f(slice[i]) {
			out = append(out, slice[i])
		}
	}

	return AccountSlice(out)
}

// Each applies a function to every Account in the slice.
func (slice AccountSlice) Each(f func(Account)) {
	for i := 0; i < len(slice); i++ {
		f(slice[i])
	}
}

// TryEach applies a function to every Account in the slice,
// and returns the index of the element that caused the first error, and the error itself.
// If every member of the slice returns nil, this function will return (-1, nil)
// The iteration will halt on the first error encountered and return it.
func (slice AccountSlice) TryEach(f func(Account) error) (int, error) {
	for i := 0; i < len(slice); i++ {
		if err := f(slice[i]); err != nil {
			return i, err
		}
	}

	return -1, nil
}

// IfEach applies a function to every Account in the slice,
// and returns the index of the element that caused the function to return false.
// If every member of the slice evaluates to true, this function will return (-1, true)
// The iteration will halt on the first false return from the function.
func (slice AccountSlice) IfEach(f func(Account) bool) (int, bool) {
	for i := 0; i < len(slice); i++ {
		if !f(slice[i]) {
			return i, false
		}
	}

	return -1, false
}
