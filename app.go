package main

import (
	"fmt"
	"log"
	"net/http"

	"bitbucket.org/augustoscher/logs-monitor-docker-postgres/dao"
	"bitbucket.org/augustoscher/logs-monitor-docker-postgres/routes"
)

func main() {
	fmt.Println("Inciando server na porta 3000...")
	dao.Connect()
	fmt.Println("Conexão com banco de dados realizada com sucesso.")
	router := routes.InitRoutes()
	if err := http.ListenAndServe(":3000", router); err != nil {
		log.Fatal(err)
	}
}
