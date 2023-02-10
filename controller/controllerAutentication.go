package controller

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/autenticacaoUsuario/models"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func GenerateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var postUser models.Users
		err := c.ShouldBindJSON(&postUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"erro ao decodificar postUser para JWT ": err,
			})
			log.Println("erro ao decodificar postUser para JWT ", err)
		}

		var user models.Users
		row := DB.QueryRow(`SELECT id, email_user, password_user FROM users WHERE email_user = $1`, postUser.EmailUser)

		err = row.Scan(&user.ID, &user.EmailUser, &user.PasswordUser)
		if err != nil {
			if err == sql.ErrNoRows {
				c.JSON(http.StatusUnauthorized, gin.H{
					"message ": "usuário ou senha inválido",
				})
				log.Println("usuário ou senha inválido ", err)
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "erro ao buscar usuário no banco de dados",
			})
			log.Println("erro ao buscar usuário no banco de dados ", err)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(user.PasswordUser), []byte(postUser.PasswordUser))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "usuário ou senha inválido",
			})
			log.Println("usuário ou senha inválido ", err)
			return
		}

		// Gerar token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"user_id": user.ID,
			"email":   user.EmailUser,
			"exp":     time.Now().Add(time.Hour).Unix(),
		})

		err = godotenv.Load("./.env")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "erro ao carregar arquivo .env para token ",
			})
			log.Println("erro ao carregar arquivo .env para token ", err)
			return
		}

		secret := os.Getenv("SECRET")

		tokenString, err := token.SignedString([]byte(secret))
		//tokenString, err := token.SignedString([]byte("secret"))
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
