package routes

import (
	"bitbucket.org/augustoscher/logs-monitor-docker-postgres/controller"
	"github.com/gorilla/mux"
)

//InitRoutes inicializa rotas
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = setRoutes(router)
	// router = SetAuthenticationRoutes(router)
	return router
}

func setRoutes(routes *mux.Router) *mux.Router {
	routes.HandleFunc("/logs-group-integracao", controller.FindLogsGroupIntegracao).Methods("GET")
	routes.HandleFunc("/logs-group-filial", controller.FindLogsGroupFilial).Methods("GET")
	routes.HandleFunc("/logs/{integracao}/filial", controller.FindLogsFilialIntegracaoEndPoint).Queries("codigo", "{codigo}").Methods("GET")
	routes.HandleFunc("/logs", controller.AllLogsPageableEndPoint).Queries("limit", "{limit}", "offset", "{offset}").Methods("GET")
	routes.HandleFunc("/logs", controller.AllLogsPageableEndPoint).Queries("limit", "{limit}", "offset", "{offset}").Methods("OPTIONS")
	// routes.HandleFunc("/logs", controller.AllLogsEndPoint).Methods("GET")
	routes.HandleFunc("/logs", controller.CreateLogEndPoint).Methods("POST")
	routes.HandleFunc("/logs", controller.UpdateLogEndPoint).Methods("PUT")
	routes.HandleFunc("/logs", controller.DeleteLogEndPoint).Methods("DELETE")
	routes.HandleFunc("/logs/{id}", controller.FindLogEndpoint).Methods("GET")
	return routes
}
