package repository

import (
	"Diploma"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateTeam(team *Diploma.Teams) (int, error)
	GetUser(teamname, password string) (Diploma.Teams, error)
}
type Repository struct {
	Authorization
}

func NewConnection(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		//Tournament:    NewTournamentPostgres(db),
	}
}
