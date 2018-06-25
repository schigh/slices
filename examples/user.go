package examples

//go:generate slices User filter(byvalue)
//go:generate slices Account all

// User is a user
type User struct {
	Name string
	Age  int
}

// Account is an account
type Account struct {
	Name string
}
