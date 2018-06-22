package examples

//go:generate slices User
//go:generate slices Account

// User is a user
type User struct {
	Name string
	Age  int
}

// Account is an account
type Account struct {
	Name string
}
