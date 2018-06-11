package examples

//go:generate slices User
//go:generate slices Account

type User struct {
	Name string
	Age  int
}

type Account struct {
	Name string
}
