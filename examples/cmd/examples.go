package main

import (
	"fmt"
	"github.com/schigh/slices/examples"
	"strings"
)

var (
	accounts1 []examples.Account
	users1    []examples.User
)

func zero() {
	accounts1 = []examples.Account{
		{"Foo"},
		{"Bar"},
		{"Baz"},
		{"Fizz"},
		{"Buzz"},
	}

	users1 = []examples.User{
		{"Tom", 32},
		{"Jen", 18},
		{"Mike", 28},
		{"Xander", 16},
		{"Della", 45},
		{"Larry", 63},
		{"Chris", 36},
	}
}

func main() {
	filterExample1()
	filterExample2()
	mapExample1()
	mapExample2()
}

func filterExample1() {
	zero()

	accountsWithF := examples.AccountSlice(accounts1).Filter(func(a *examples.Account) bool {
		name := strings.ToLower(a.Name)
		return len(name) > 0 && name[0] == 'f'
	}).Value()

	fmt.Println("filterExample1", accountsWithF)
}

func filterExample2() {
	zero()

	usersYoungerThan40 := examples.UserSlice(users1).Filter(func(u *examples.User) bool {
		return u.Age < 40
	}).Value()

	fmt.Println("filterExample2", usersYoungerThan40)
}

func mapExample1() {
	zero()

	examples.UserSlice(users1).Map(func(u *examples.User) {
		u.Name = "Steve"
	})

	fmt.Println("mapExample1", users1)
}

func mapExample2() {
	zero()

	examples.UserSlice(users1).Map(func(u *examples.User) {
		if u.Age < 30 {
			u.Name = "Youngling"
		}
	})

	fmt.Println("mapExample2", users1)
}
