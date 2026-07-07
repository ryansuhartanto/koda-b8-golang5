package data

import "slices"

type Database []User

func (this *Database) Register(
	first string,
	last string,
	email string,
	password string,
) {
	user := NewUser(
		first,
		last,
		email,
		password,
	)

	*this = append(*this, *user)
}

func (this *Database) Login(
	email string,
	password string,
) *User {
	index := slices.IndexFunc(*this, func(user User) bool {
		return user.Auth(email, password)
	})

	if index == -1 {
		panic("")
	}

	return &(*this)[index]
}
