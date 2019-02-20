package dao

import (
	"os"

	"github.com/go-pg/pg"
)

var db *pg.DB

//Connect realiza conexão com banco de dados
func Connect() {

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPasswd := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	address := string(dbHost + ":" + dbPort)

	db = pg.Connect(&pg.Options{
		Addr:     address,
		User:     dbUser,
		Password: dbPasswd,
		Database: dbName,
	})
}
