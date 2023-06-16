package Diploma

type Teams struct {
	Team_id   int    `json:"-" db:"team_id"`
	Team_name string `json:"teamName"`
	Player1   string `json:"player1"`
	Player2   string `json:"player2"`
	Player3   string `json:"player3"`
	Player4   string `json:"player4"`
	Player5   string `json:"player5"`
	Password  string `json:"password"`
	//Cpassword string `json:"cpassword" validate:"required,eqfield=Password" label:"Confirm Password"`
}
