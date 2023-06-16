package Diploma

import "time"

type Tournament struct {
	TournamentId   int    `json:"-" db:"tournament_id"`
	TournamentName string `json:"tournamentName" binding:"required"`
	Url            string `json:"url" binding:"required"`
	Description    string `json:"description" binding:"required"`
	//size            int    `json:"size" binding:"required"`
	Time time.Duration `json:"time" binding:"required"`
}
