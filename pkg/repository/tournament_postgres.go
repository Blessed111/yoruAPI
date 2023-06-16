package repository

import (
	"Diploma"

	"github.com/jmoiron/sqlx"
)

type TourPostgres struct {
	db *sqlx.DB
}

func NewTournamentPostgres(db *sqlx.DB) *TourPostgres {
	return &TourPostgres{
		db: db,
	}
}

func (T *TourPostgres) CreateTournament(tournament *Diploma.Tournament) (int, error) {
	var id int
	//query := fmt.Sprintf("INSERT INTO %s (tournament_name,team_id) values ($1,$2) Returning tournament_id", tournamenttable)
	//row := T.db.QueryRow(query, tournament.Tournament_name, tournament.Team_id)
	//if err := row.Scan(&id); err != nil {
	//	return 0, err
	//}
	return id, nil
}
func (T *TourPostgres) GetTour(tournament *Diploma.Tournament) (int, error) {
	//query:=fmt.Sprintf("SELECT tournament_id from %s where tournament_name=$1 and team_id=$2",tournamenttable)
	return 0, nil
}
