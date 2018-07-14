package examples

import (
	"math/rand"
	"testing"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func makeName() string {
	caps := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lower := "abcdefghijklmnopqrstuvwxyz"
	size := 1 + rand.Intn(19)

	nb := make([]byte, size)
	nb[0] = caps[rand.Intn(25)]
	for i := 1; i < size-1; i++ {
		nb[i] = lower[rand.Intn(25)]
	}
	return string(nb)
}

func makeUsers(num int) []User {
	users := make([]User, num)
	for i := 0; i < num; i++ {
		users[i] = User{
			Name: makeName(),
			Age:  1 + rand.Intn(99),
		}
	}

	return users
}

func BenchmarkUserSlice_Map(b *testing.B) {
	benchmarks := []struct {
		name    string
		slice   []User
		mapFunc func(User) User
	}{
		{
			name:  "10 users",
			slice: makeUsers(10),
			mapFunc: func(user User) User {
				user.Name = "woot"
				user.Age = 36

				return user
			},
		},
		{
			name:  "100 users",
			slice: makeUsers(100),
			mapFunc: func(user User) User {
				user.Name = "woot"
				user.Age = 36

				return user
			},
		},
		{
			name:  "1000 users",
			slice: makeUsers(1000),
			mapFunc: func(user User) User {
				user.Name = "woot"
				user.Age = 36

				return user
			},
		},
		{
			name:  "10000 users",
			slice: makeUsers(10000),
			mapFunc: func(user User) User {
				user.Name = "woot"
				user.Age = 36

				return user
			},
		},
		{
			name:  "100000 users",
			slice: makeUsers(100000),
			mapFunc: func(user User) User {
				user.Name = "woot"
				user.Age = 36

				return user
			},
		},
	}
	for _, bm := range benchmarks {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				UserSlice(bm.slice).Map(bm.mapFunc)
			}
		})
	}
}
