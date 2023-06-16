package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	Authorization = "Authorization"
	teamCtx       = "teamId"
)

func (h *Handlers) userIdentity(c *gin.Context) {
	header := c.GetHeader(Authorization)
	if header == "" {
		newErrorresponse(c, http.StatusUnauthorized, "empty  auth header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorresponse(c, http.StatusUnauthorized, "auth header wrongly ")
		return
	}
	teamId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorresponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(teamCtx, teamId)
}

func MyProtectedRouteHandler(c *gin.Context) {
	tokenString := c.Request.Header.Get("Authorization")
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		return
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// return the secret key used to sign the token
		return []byte("secret-key"), nil
	})

	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	// set the user ID from the token in the request context
	claims, _ := token.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)

	// Pass the user ID to the HTML template
	tmpl, err := template.ParseFiles("mytemplate.html")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to render template"})
		return
	}
	tmpl.Execute(c.Writer, gin.H{"loggedIn": true, "userID": userID})
}
