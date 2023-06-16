package handlers

import (
	"Diploma"
	"Diploma/pkg/cfg"
	"Diploma/pkg/libraries"
	"Diploma/pkg/models"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var validation = libraries.NewValidation()
var teamModel = models.NewTeamModel()

type signInInput struct {
	Team_name string `json:"teamName"`
	Password  string `json:"password"`
}

func (h *Handlers) Login(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		session, _ := cfg.Store.Get(c.Request, cfg.SESSION_ID)
		if len(session.Values) != 0 {
			c.Redirect(http.StatusSeeOther, "/home")
		}
		// if session.Values["loggedIn"] == true {
		// 	log.Println("already logged in")
		// 	c.Redirect(http.StatusSeeOther, "/home")

		// 	return
		// }
		c.HTML(http.StatusOK, "login.html", nil)
	} else if c.Request.Method == http.MethodPost {
		TeamInput := &signInInput{
			Team_name: c.PostForm("teamName"),
			Password:  c.PostForm("password"),
		}

		errorMessages := validation.Struct(TeamInput)

		if err := c.ShouldBindJSON(&TeamInput); errorMessages != nil {

			c.JSON(http.StatusBadRequest, gin.H{
				"validation": errorMessages,
				"error":      err,
			})
			c.Abort()
		} else {
			if TeamInput.Team_name == "admin" && TeamInput.Password == "admin" {
				c.Redirect(http.StatusSeeOther, "/admin/teams")
			}
			var team Diploma.Teams
			teamModel.Where(&team, "team_name", TeamInput.Team_name)

			var message error

			errPassword := bcrypt.CompareHashAndPassword([]byte(team.Password), []byte(TeamInput.Password))
			if errPassword != nil {
				message = errors.New("Wrong teamname or password!")
				log.Println("wrong password")
			}
			if err := c.ShouldBindJSON(&TeamInput); message != nil {

				c.JSON(http.StatusBadRequest, gin.H{
					"err1":  err.Error(),
					"error": "Wrong teamname or password!",
				})
				c.Abort()
			} else {
				// set session
				session, _ := cfg.Store.Get(c.Request, cfg.SESSION_ID)

				session.Values["loggedIn"] = true
				// session.Values["email"] = user.Email
				session.Values["teamname"] = team.Team_name
				session.Values["player1"] = team.Player1
				session.Values["player2"] = team.Player2
				session.Values["player3"] = team.Player3
				session.Values["player4"] = team.Player4
				session.Values["player5"] = team.Player5
				session.Values["password"] = team.Password
				// session.Values["full_name"] = user.Full_name
				// session.Values["role"] = user.Role
				session.Values["id"] = team.Team_id
				session.Save(c.Request, c.Writer)

				c.JSON(http.StatusOK, gin.H{
					//"session":  session,
					"token":    cfg.Store,
					"loggedIn": true,
					"user":     session.Values["teamname"],
					"player1":  session.Values["player1"],
					"player2":  session.Values["player2"],
					"player3":  session.Values["player3"],
					"player4":  session.Values["player4"],
					"player5":  session.Values["player5"],
					"team_id":  session.Values["id"],
					"user1":    "team logged",
				})
				c.Abort()

			}
		}
	}
}

func (h *Handlers) Register(c *gin.Context) {
	var team Diploma.Teams
	errorMessages := validation.Struct(team)
	if err := c.BindJSON(&team); err != nil {
		newErrorresponse(c, http.StatusBadRequest, err.Error())
		return
	}
	if (team.Team_name == "") || (team.Player1 == "") || (team.Player2 == "") || (team.Player3 == "") || (team.Player4 == "") || (team.Player5 == "") || (team.Password == "") {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Fields cannot be empty",
		})
	} else if err := c.ShouldBindJSON(&team); errorMessages != nil {

		c.JSON(http.StatusBadRequest, gin.H{
			"validation": errorMessages,
			"err":        err.Error(),
		})
		c.Abort()
	} else {
		hashPassword, _ := bcrypt.GenerateFromPassword([]byte(team.Password), bcrypt.DefaultCost)
		team.Password = string(hashPassword)
		_, err := teamModel.CreateTeam111(team)

		if err != nil {
			newErrorresponse(c, http.StatusInternalServerError, err.Error())
		} else {
			c.JSON(http.StatusOK, map[string]interface{}{
				"team1":  "Team added successfully",
				"status": "success",
				"teams":  team,
			})
		}
	}
}

func (h *Handlers) Logout(c *gin.Context) {
	session, _ := cfg.Store.Get(c.Request, cfg.SESSION_ID)
	// delete session
	session.Options.MaxAge = -1
	err := session.Save(c.Request, c.Writer)
	if err != nil {
		log.Println("session not saved")
	}
	c.Redirect(http.StatusSeeOther, "/auth/logi")
}
