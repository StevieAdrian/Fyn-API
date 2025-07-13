package usecase

import (
	"context"
	"errors"

	"github.com/StevieAdrian/Fyn-API/auth-service/internal/auth/domain"
	"github.com/StevieAdrian/Fyn-API/auth-service/internal/auth/dto"
	"github.com/StevieAdrian/Fyn-API/auth-service/internal/auth/mapper"

	// "github.com/StevieAdrian/Fyn-API/auth-service/models"
	"github.com/StevieAdrian/Fyn-API/auth-service/pkg/hash"
	"github.com/StevieAdrian/Fyn-API/auth-service/pkg/token"
	"github.com/google/uuid"
)

type UserUsecase interface {
	Signup(ctx context.Context, input dto.SignupRequest) error
	Login(ctx context.Context, email, password string) (*domain.User, string, string, error)
	GetAll(ctx context.Context) ([]domain.User, error)
}

type userUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(repo domain.UserRepository) UserUsecase {
	return &userUsecase{
		repo: repo,
	}
}

func (u *userUsecase) Signup(ctx context.Context, input dto.SignupRequest) error {
	existing, _ := u.repo.GetByEmail(input.Email)
	if existing != nil {
		return errors.New("email or phone number already used")
	}

	user := mapper.FromSignupDTOToDomain(input)
	user.UserID = uuid.NewString()
	hashed := hash.HashPassword(&user.Password)
	user.Password = *hashed

	accessToken, refreshToken := token.GenerateToken(user.Email, user.UserID, user.Role)
	user.Token = accessToken
	user.RefreshToken = refreshToken

	return u.repo.CreateUser(&user)
}

func (u *userUsecase) Login(ctx context.Context, email, password string) (*domain.User, string, string, error) {
	domainUser, err := u.repo.GetByEmail(email)
	if err != nil {
		return nil, "", "", errors.New("email not found")
	}

	ok, msg := hash.VerifyPassword(domainUser.Password, password)
	if !ok {
		return nil, "", "", msg
	}

	accessToken, refreshToken := token.GenerateToken(domainUser.Email, domainUser.UserID, domainUser.Role)
	_ = token.UpdateAllToken(accessToken, refreshToken, domainUser.UserID)

	domainUser.Token = accessToken
	domainUser.RefreshToken = refreshToken

	return domainUser, accessToken, refreshToken, nil
}

func (u *userUsecase) GetAll(ctx context.Context) ([]domain.User, error) {
	return u.repo.GetAllUsers()
}
