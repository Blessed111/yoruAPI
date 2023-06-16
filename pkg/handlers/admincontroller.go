package handlers

import (
	"Diploma"
	"Diploma/pkg/cfg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

//var model = models.NewTeamModel()

func (H *Handlers) HomePage(c *gin.Context) {
	session, _ := cfg.Store.Get(c.Request, cfg.SESSION_ID)
	//var team models.TeamModel
	if len(session.Values) != 0 {
		c.JSON(http.StatusOK, gin.H{
			"isnotAuthorized": session.Values["loggedIn"] == false,
			//"role":            session.Values["role"],
			"teamname": session.Values["teamname"],
			"Team_id":  session.Values["id"]})
		c.Abort()
	} else if len(session.Values) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"isnotAuthorized": session.Values["loggedIn"] != false,
		})
		c.Abort()
	}

}

func (h *Handlers) AddTeam(c *gin.Context) {
	if c.Request.Method == http.MethodGet {

		c.HTML(http.StatusOK, "admin_add.html", nil)

	} else if c.Request.Method == http.MethodPost {
		team := Diploma.Teams{
			Team_name: c.PostForm("teamName"),
			Player1:   c.PostForm("player1"),
			Player2:   c.PostForm("player2"),
			Player3:   c.PostForm("player3"),
			Player4:   c.PostForm("player4"),
			Player5:   c.PostForm("player5"),
			Password:  c.PostForm("password"),
		}
		if err := c.ShouldBind(&team); err != nil {
			c.HTML(http.StatusBadRequest, "admin_add.html", gin.H{"error": err.Error()})
			return
		}

		errorMessages := validation.Struct(team)

		if errorMessages != nil {
			data := gin.H{
				"validation": errorMessages,
				"team":       team,
			}
			c.HTML(http.StatusBadRequest, "admin_add.html", data)
			return

		} else {

			// Hash password
			// hash := sha1.New()
			// hash.Write([]byte(team.Password))

			hashPassword, _ := bcrypt.GenerateFromPassword([]byte(team.Password), bcrypt.DefaultCost)
			team.Password = string(hashPassword)
			// Create user in the database
			teamModel.CreateTeam111(team)

			data := gin.H{
				"team1": "Team added successfully",
			}
			c.HTML(http.StatusOK, "admin_add.html", data)

			c.JSON(http.StatusOK, gin.H{
				"team1":  "Team added successfully",
				"status": "success",
				"teams":  team,
			})
			c.Redirect(http.StatusMovedPermanently, "/admin/teams")
		}
	}

}

func (h *Handlers) getTeamsList(c *gin.Context) {

	teams, _ := teamModel.FindAll()
	c.HTML(200, "admin.html", gin.H{
		"teams": teams,
	})
	// c.JSON(http.StatusOK, gin.H{
	// 	"status": "success",
	// 	"teams":  teams,
	// })

}

func (h *Handlers) updateTeam(c *gin.Context) {
	if c.Request.Method == http.MethodGet {
		id, _ := strconv.ParseInt(c.Query("id"), 10, 64)

		var team Diploma.Teams
		teamModel.Find(id, &team)

		data := gin.H{
			"team": team,
		}

		c.HTML(http.StatusOK, "admin_update.html", data)
	} else if c.Request.Method == http.MethodPost {

		var team Diploma.Teams

		id, _ := strconv.Atoi(c.PostForm("id"))
		team.Team_id = int(id)
		team.Team_name = c.PostForm("teamName")
		team.Player1 = c.PostForm("player1")
		team.Player2 = c.PostForm("player2")
		team.Player3 = c.PostForm("player3")
		team.Player4 = c.PostForm("player4")
		team.Player5 = c.PostForm("player5")

		teamModel.Update(team)
		// c.JSON(http.StatusOK, gin.H{
		// 	"user1": "User data successfully updated",
		// })
		data := gin.H{
			"team1": "Team updated successfully",
		}
		c.HTML(http.StatusOK, "admin_update.html", data)

		//}

		c.Redirect(http.StatusMovedPermanently, "/admin/teams/update")

	}
}

func (h *Handlers) deleteTeam(c *gin.Context) {
	idStr := c.Query("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	teamModel.Delete(id)
	// c.JSON(http.StatusOK, gin.H{
	// 	"status":  "success",
	// 	"message": "User deleted successfully",
	// })

	c.Redirect(http.StatusMovedPermanently, "/admin/teams")
}

//TEST

// func ForgotPassword(c *gin.Context) {
// 	email := c.PostForm("email")

// 	// Check if email exists in the database
// 	team, err := teamModel.GetByEmail(email)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": "Email not found"})
// 		return
// 	}

// 	// Generate a new password
// 	newPassword := models.generatePassword()

// 	// Update the user's password in the database
// 	team = hashPassword(newPassword)
// 	if err := TeaModel.UpdatePassword(user); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update password"})
// 		return
// 	}

// 	// Send the new password to the user's email
// 	if err := sendNewPasswordEmail(user.Email, newPassword); err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "New password has been sent to your email"})
// }
