package service

import (
	"example.com/hexagonal-auth/domain"
	"example.com/hexagonal-auth/dto"
)

type AuthService interface {
	Login(dto.LoginRequest) (*dto.LoginResponse, error)
}

type DefaultAuthService struct {
	repo            domain.AuthRepository
	rolePermissions domain.RolePermissions
}

func (s DefaultAuthService) Login(req dto.LoginRequest) (*dto.LoginResponse, error) {
	login, err := s.repo.FindBy(req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	token, err := login.GenerateToken()
	if err != nil {
		return nil, err
	}
	return &dto.LoginResponse{AccessToken: *token}, nil
}

func NewLoginService(repo domain.AuthRepository, permission domain.RolePermissions) DefaultAuthService {
	return DefaultAuthService{repo: repo, rolePermissions: permission}
}
