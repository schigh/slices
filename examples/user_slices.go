// generated by slices (github.com/schigh/slices).  do not edit.
// gen date: Sat, 14 Jul 2018 00:30:20 -0400

package examples

// UserSlice aliases []User
type UserSlice []User

// Value returns the wrapped User slice
func (slice UserSlice) Value() []User {
	return []User(slice)
}

// Filter evaluates every element in the slice, and returns all User
// instances where the eval function returns true
func (slice UserSlice) Filter(f func(User) bool) UserSlice {
	out := make([]User, 0, len(slice))
	for _, v := range slice {
		if f(v) {
			out = append(out, v)
		}
	}

	return UserSlice(out)
}

// Map applies a function to every User in the slice.  This function will mutate the slice in place
func (slice UserSlice) Map(f func(User) User) {
	for i, v := range slice {
		slice[i] = f(v)
	}
}
