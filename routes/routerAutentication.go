package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/autenticacaoUsuario/controller"
	"github.com/guilherm5/autenticacaoUsuario/middleware"
)

func AuthUser(c *gin.Engine) {
	c.POST("/postAuth", controller.GenerateJWT())
	c.GET("/getUser", middleware.ValidateJWT(), controller.GetUserTest())
}
