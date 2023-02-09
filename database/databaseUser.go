package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Init() *sql.DB {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Println("erro ao carregar variaveis de amviente ", err)
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	stringConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", stringConnection)
	if err != nil {
		log.Println("erro ao iniciar conexão com banco de dados ", err)
	} else {
		fmt.Println("sucesso ao realizar conexão com o banco de dados")
	}

	return db
}
