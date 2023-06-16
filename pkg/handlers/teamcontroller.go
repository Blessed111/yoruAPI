package handlers

import (
	"Diploma"
	"Diploma/pkg/cfg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (H *Handlers) TournamentInfo(c *gin.Context) {
	session, _ := cfg.Store.Get(c.Request, cfg.SESSION_ID)
	//var team models.TeamModel
	if c.Request.Method == http.MethodGet {
		data := gin.H{
			"isnotAuthorized": session.Values["loggedIn"] != true,
			//"role":            session.Values["role"],
			"teamname": session.Values["teamname"],
			"Team_id":  session.Values["id"],
		}
		c.HTML(http.StatusOK, "teampage.html", data)

	}
}

func (h *Handlers) teamProfile(c *gin.Context) {
	session, _ := cfg.Store.Get(c.Request, cfg.SESSION_ID)
	//var team models.TeamModel
	if c.Request.Method == http.MethodGet {
		var message error
		var team Diploma.Teams
		idStr := c.Param("id")
		id, _ := strconv.ParseInt(idStr, 10, 64)
		teamModel.Find(id, &team)
		if err := c.ShouldBindJSON(&team); message != nil {
			newErrorresponse(c, http.StatusBadRequest, err.Error())
			return
		} else {
			c.JSON(http.StatusOK, map[string]interface{}{
				"status":  "success",
				"team":    session.Values["teamname"],
				"player1": session.Values["player1"],
				"player2": session.Values["player2"],
				"player3": session.Values["player3"],
				"player4": session.Values["player4"],
				"player5": session.Values["player5"],
				"team_id": session.Values["id"],

				"teams": team,
			})
		}
	}
}
