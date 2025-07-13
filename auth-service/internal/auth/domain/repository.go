package domain

type UserRepository interface {
	CreateUser(user *User) error
	GetByEmail(email string) (*User, error)
	GetAllUsers() ([]User, error)
}
