package service

import (
	"Diploma"
	"Diploma/pkg/repository"
)

type Authorization interface {
	CreateTeam(teams *Diploma.Teams) (int, error)
	GenerateToken(teamname string, password string) (string, error)
	ParseToken(token string) (int, error)
	//CreateTeam111(team Diploma.Teams) (int64, error)
}
type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
