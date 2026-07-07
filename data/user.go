package data

import "crypto/md5"

type User struct {
	First    string
	Last     string
	Email    string
	Password [16]byte
}

func NewUser(
	first string,
	last string,
	email string,
	password string,
) *User {
	passwordHash := md5.Sum([]byte(password))

	return &User{
		First:    first,
		Last:     last,
		Email:    email,
		Password: passwordHash,
	}
}

func (this *User) Auth(
	email string,
	password string,
) bool {
	passwordHash := md5.Sum([]byte(password))

	return this.Email == email && this.Password == passwordHash
}

func (this *User) ChangePassword(password string) {
	passwordHash := md5.Sum([]byte(password))

	this.Password = passwordHash
}
