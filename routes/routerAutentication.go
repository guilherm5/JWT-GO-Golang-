package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/autenticacaoUsuario/controller"
)

func AuthUser(c *gin.Engine) {
	c.POST("/postAuth", controller.PostAuth())
}
