package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pass-wall/passwall-server/helper"
	"github.com/pass-wall/passwall-server/login"
)

// GeneratePassword generates new password
func GeneratePassword(c *gin.Context) {
	password := helper.Password()
	response := login.LoginResponse{"Success", password}
	c.JSON(http.StatusOK, response)
}
