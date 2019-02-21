package dao

import (
	"os"

	"github.com/go-pg/pg"
)

var db *pg.DB

//Connect realiza conexão com banco de dados
func Connect() {

	dbHost := os.Getenv("DB_HOST")
	if len(dbHost) <= 0 {
		dbHost = "localhost"
	}
	dbPort := os.Getenv("DB_PORT")
	if len(dbPort) <= 0 {
		dbPort = "5432"
	}
	dbUser := os.Getenv("DB_USER")
	if len(dbUser) <= 0 {
		dbUser = "postgres"
	}
	dbPasswd := os.Getenv("DB_PASS")
	if len(dbPasswd) <= 0 {
		dbPasswd = "postgres"
	}
	dbName := os.Getenv("DB_NAME")
	if len(dbName) <= 0 {
		dbName = "log_monitor"
	}
	address := string(dbHost + ":" + dbPort)

	db = pg.Connect(&pg.Options{
		Addr:     address,
		User:     dbUser,
		Password: dbPasswd,
		Database: dbName,
	})
}
