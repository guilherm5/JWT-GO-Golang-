package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func ValidateJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		headerToken := c.GetHeader("Authorization")

		if headerToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token não fornecido",
			})
			c.Abort()
			return
		}

		//verificando se o token esta no formato certo
		splitted := strings.Split(headerToken, " ")
		if len(splitted) != 2 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Token inválido",
			})
			log.Println("Token invalido ")
			c.Abort()
			return
		}

		//pegando o token que foi verificado e armazenando na variavel tokenString
		tokenString := splitted[1]
		err := godotenv.Load("./.env")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "erro ao carregar arquivo .env para validar token ",
			})
			log.Println("erro ao carregar arquivo .env para validar token ", err)
			c.Abort()
			return
		}

		secret := os.Getenv("SECRET")
		//secret := "secret"

		//decodificando o token e verificando se a chave que passei para gerar o token (secret) esta correta.
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			return []byte(secret), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token inválido",
			})
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("user_id", claims["user_id"])
			c.Set("email", claims["email"])
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token inválido",
			})
			c.Abort()
			return
		}
	}
}
