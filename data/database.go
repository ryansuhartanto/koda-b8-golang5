package data

import "slices"

type Database struct {
	Current *User
	Users   []User
}

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

	this.Users = append(this.Users, *user)
}

func (this *Database) Login(
	email string,
	password string,
) {
	index := slices.IndexFunc(this.Users, func(user User) bool {
		return user.Auth(email, password)
	})

	if index == -1 {
		panic("")
	}

	this.Current = &this.Users[index]
}

func (this *Database) ForgotPassword(
	email string,
	auth bool,
) *User {
	index := slices.IndexFunc(this.Users, func(user User) bool {
		return user.Email == email
	})

	if index == -1 {
		panic("")
	}

	user := this.Users[index]

	if !auth {
		panic("")
	}

	return &user
}

func (this *Database) Logout() {
	this.Current = nil
}
