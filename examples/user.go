package examples

//go:generate slices User filter(byvalue) map(byvalue)
//go:generate slices Account all(byvalue)

// User is a user
type User struct {
	Name string
	Age  int
}

// Account is an account
type Account struct {
	Name  string
	Value float64
}
