package mapper

import (
	"github.com/StevieAdrian/Fyn-API/auth-service/internal/auth/domain"
	"github.com/StevieAdrian/Fyn-API/auth-service/internal/auth/dto"
	"github.com/StevieAdrian/Fyn-API/auth-service/internal/auth/repository/gorm"
	// "github.com/StevieAdrian/Fyn-API/auth-service/models"
)

func FromSignupDTOToDomain(dto dto.SignupRequest) domain.User {
	return domain.User{
		FirstName: dto.FirstName,
		LastName:  dto.LastName,
		Email:     dto.Email,
		Phone:     dto.Phone,
		Password:  dto.Password,
		Role:      dto.Role,
	}
}

func FromDomainToModel(user domain.User) gorm.UserModel {
	return gorm.UserModel{
		UserID:       user.UserID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		Phone:        user.Phone,
		Password:     user.Password,
		Role:         user.Role,
		Token:        user.Token,
		RefreshToken: user.RefreshToken,
	}
}

func FromModelToDomain(m gorm.UserModel) domain.User {
	return domain.User{
		UserID:       m.UserID,
		FirstName:    m.FirstName,
		LastName:     m.LastName,
		Email:        m.Email,
		Phone:        m.Phone,
		Password:     m.Password,
		Role:         m.Role,
		Token:        m.Token,
		RefreshToken: m.RefreshToken,
	}
}

func getStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
