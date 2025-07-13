package domain

type User struct {
	UserID       string
	FirstName    string
	LastName     string
	Email        string
	Phone        string
	Password     string
	Role         string
	Token        string
	RefreshToken string
}
