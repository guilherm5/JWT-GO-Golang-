package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilherm5/autenticacaoUsuario/database"
	"github.com/guilherm5/autenticacaoUsuario/models"
)

var DB = database.Init()

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var getAllUsers []models.Users

		rows, err := DB.Query(`SELECT * FROM users`)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"erro ao preparar select nos usuarios ": err,
			})
			log.Println("erro ao preparar select nos usuarios ", err)
		}

		for rows.Next() {
			var getAll models.Users

			if err := rows.Scan(&getAll.ID, &getAll.NameUser, &getAll.EmailUser, &getAll.PasswordUser); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"erro ao realizar scan no select nos usuarios ": err,
				})
				log.Println("erro ao realizar scan no select nos usuarios ", err)
			} else {
				getAllUsers = append(getAllUsers, getAll)
				c.JSON(http.StatusOK, getAll)
			}
		}
	}
}

func PostUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var postUser models.InsertUsers
		err := c.ShouldBindJSON(&postUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"erro ao decodificar postUser ": err,
			})
			log.Println("erro ao decodificar postUser ", err)
		}

		posting, err := DB.Prepare(`INSERT INTO users (name_user, email_user, password_user) VALUES ($1, $2, $3)`)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"erro ao preparar insert postUser ": err,
			})
			log.Println("erro ao preparar insert postUser ", err)
		}

		for _, add := range postUser.Records {
			_, err := posting.Exec(add.NameUser, add.EmailUser, add.PasswordUser)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"erro ao executar insert postUser ": err,
				})
				log.Println("erro ao executar insert postUser ", err)
			} else {
				c.JSON(http.StatusOK, add)
			}
		}
	}
}

func PutUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var update models.InsertUsers

		err := c.ShouldBindJSON(&update)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"erro ao decodificar update de usuario ": err,
			})
			log.Println("erro ao decodificar update de usuario ", err)
		}

		updating, err := DB.Prepare(`UPDATE users SET name_user = $1, email_user = $2, password_user = $3 WHERE id = $4`)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"erro ao preparar update de usuario ": err,
			})
			log.Println("erro ao preparar update de usuario ", err)
		}

		for _, UP := range update.Records {
			_, err := updating.Exec(UP.NameUser, UP.EmailUser, UP.PasswordUser, UP.ID)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"erro ao executar update de usuario ": err,
				})
				log.Println("erro ao executar update de usuario ", err)
			} else {
				c.JSON(http.StatusOK, update)
			}
		}
	}
}

func DeleteUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var deleteUser models.DeleteUsers

		err := c.ShouldBindJSON(&deleteUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"erro ao decodificar delete de usuario ": err,
			})
			log.Println("erro ao decodificar delete de usuario ", err)
		}

		deleting, err := DB.Prepare(`DELETE FROM users WHERE id = $1`)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"erro ao preparar delete de usuario ": err,
			})
			log.Println("erro ao preparar delete de usuario ", err)
		}

		for _, deleted := range deleteUser.ID {
			_, err := deleting.Exec(&deleted)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"erro ao executar delete de usuario ": err,
				})
				log.Println("erro ao executar delete de usuario ", err)
			} else {
				c.JSON(http.StatusOK, deleted)
			}
		}

	}
}
