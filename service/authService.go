package service

import (
	"example.com/hexagonal-auth/domain"
	"example.com/hexagonal-auth/dto"
	"example.com/hexagonal-auth/errs"
)

type AuthService interface {
	Login(dto.LoginRequest) (*dto.LoginResponse, *errs.AppError)
}

type DefaultAuthService struct {
	repo            domain.AuthRepository
	rolePermissions domain.RolePermissions
}

func (s DefaultAuthService) Login(req dto.LoginRequest) (*dto.LoginResponse, *errs.AppError) {
	var appErr *errs.AppError
	var login *domain.Login
	var token *string

	if login, appErr = s.repo.FindBy(req.Username, req.Password); appErr != nil {
		return nil, appErr
	}

	if token, appErr = login.GenerateToken(); appErr != nil {
		return nil, appErr
	}

	return &dto.LoginResponse{AccessToken: *token}, nil
}

func NewLoginService(repo domain.AuthRepository, permission domain.RolePermissions) DefaultAuthService {
	return DefaultAuthService{repo: repo, rolePermissions: permission}
}
