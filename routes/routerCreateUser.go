package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/autenticacaoUsuario/controller"
)

func CreateUser(c *gin.Engine) {
	c.GET("/user", controller.GetUser())
	c.POST("/user", controller.PostUser())
	c.DELETE("/user", controller.DeleteUser())
	c.PUT("/user", controller.PutUser())
}
