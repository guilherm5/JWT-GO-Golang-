package controller

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/autenticacaoUsuario/models"
)

func PostAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		var postUser models.Users
		err := c.ShouldBindJSON(&postUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"erro ao decodificar postUser ": err,
			})
			log.Println("erro ao decodificar postUser ", err)
		}
		var user models.Users
		row := DB.QueryRow(`SELECT id, email_user, password_user FROM users WHERE email_user = $1 and password_user = $2`, postUser.EmailUser, postUser.PasswordUser)

		err = row.Scan(&user.ID, &user.EmailUser, &user.PasswordUser)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message": "usuário ou senha inválido",
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "erro ao buscar usuário no banco de dados",
			})
			log.Println("erro ao buscar usuário no banco de dados ", err)
			return
		}
		// Gerar token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": user.ID,
			"email":   user.EmailUser,
			"exp":     time.Now().Add(time.Hour * 24).Unix(),
		})

		tokenString, err := token.SignedString([]byte("secret"))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "erro ao gerar token",
			})
			log.Println("erro ao gerar token ", err)
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": tokenString,
		})
	}
}
