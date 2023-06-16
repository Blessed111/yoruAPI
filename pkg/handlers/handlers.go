package handlers

import (
	"Diploma/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	services *service.Service
}

func NewHandler(service *service.Service) *Handlers {
	return &Handlers{services: service}
}

func (h *Handlers) Init() *gin.Engine {
	router := gin.Default()
	//MAIN
	router.GET("/home", h.HomePage)
	//AUTH
	auth := router.Group("/auth")
	{
		auth.POST("/register", h.Register)
		auth.GET("/register", h.Register)
		auth.POST("/login", h.Login)
		auth.GET("/login", h.Login)
	}
	//USER INTERFACE
	team := router.Group("/teams")
	{
		team.GET("/team/profile/:id", h.teamProfile)
		team.POST("/team/update", h.teamProfile)
	}

	//TOURNAMENTS
	tournament := router.Group("/tournaments")
	{
		tournament.GET("/tournament", h.TournamentInfo)
	}
	//LOGOUT
	router.GET("/lout", h.Logout)

	//ADMIN INTERFACE
	admin := router.Group("/admin")
	{
		admin.GET("/teams", h.getTeamsList)
		admin.GET("/teams/delete", h.deleteTeam)
		admin.POST("/teams/create", h.AddTeam)
		admin.GET("/teams/add", h.AddTeam)
		admin.POST("/teams/update", h.updateTeam)
		admin.GET("/teams/update", h.updateTeam)
	}
	return router
}
