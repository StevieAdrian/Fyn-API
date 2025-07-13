package gorm

import (
	"time"

	"github.com/StevieAdrian/Fyn-API/auth-service/internal/auth/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type UserModel struct {
	ID           uint   `gorm:"primaryKey"`
	UserID       string `gorm:"uniqueIndex"`
	FirstName    string
	LastName     string
	Email        string `gorm:"uniqueIndex"`
	Phone        string
	Password     string
	Role         string
	Token        string
	RefreshToken string
	CreatedAt    time.Time `gorm:"autoCreateTime"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

func toUserModel(u *domain.User) *UserModel {
	return &UserModel{
		UserID:       u.UserID,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		Email:        u.Email,
		Phone:        u.Phone,
		Password:     u.Password,
		Role:         u.Role,
		Token:        u.Token,
		RefreshToken: u.RefreshToken,
	}
}

func (UserModel) TableName() string {
	return "users"
}

func toDomainUser(u *UserModel) *domain.User {
	return &domain.User{
		UserID:       u.UserID,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		Email:        u.Email,
		Phone:        u.Phone,
		Password:     u.Password,
		Role:         u.Role,
		Token:        u.Token,
		RefreshToken: u.RefreshToken,
	}
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) CreateUser(user *domain.User) error {
	model := toUserModel(user)
	return r.db.Create(model).Error
}

func (r *UserRepository) GetByEmail(email string) (*domain.User, error) {
	var model UserModel
	if err := r.db.Where("email = ?", email).First(&model).Error; err != nil {
		return nil, err
	}
	return toDomainUser(&model), nil
}

func (r *UserRepository) GetAllUsers() ([]domain.User, error) {
	var models []UserModel
	if err := r.db.Find(&models).Error; err != nil {
		return nil, err
	}

	users := make([]domain.User, len(models))
	for i, m := range models {
		users[i] = *toDomainUser(&m)
	}

	return users, nil
}
