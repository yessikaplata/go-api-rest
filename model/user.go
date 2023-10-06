package model

type User struct {
	Id       int
	UserName string
	Password string
	Email    string
}

func NewUser(userName, password, email string) User {
	return User{
		UserName: userName,
		Password: password,
		Email:    email,
	}
}
