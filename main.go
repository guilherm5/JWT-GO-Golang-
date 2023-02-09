package main

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/autenticacaoUsuario/routes"
)

func main() {

	router := gin.New()
	router.Use(gin.Logger())

	routes.CreateUser(router)
	routes.AuthUser(router)

	router.Run(":5000")
}
