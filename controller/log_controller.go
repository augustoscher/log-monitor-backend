package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"bitbucket.org/augustoscher/logs-monitor-docker-postgres/dao"
	"bitbucket.org/augustoscher/logs-monitor-docker-postgres/model"
	"github.com/gorilla/mux"
)

var logDAO = dao.LogDAO{}

//AllLogsPageableEndPoint return todos paginado
func AllLogsPageableEndPoint(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)
	limit, err := strconv.Atoi(params["limit"])
	offset, err := strconv.Atoi(params["offset"])

	logs, err := logDAO.FindAllPageable(limit, offset)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, logs)
}

//FindLogsFilialIntegracaoEndPoint return todos filtrando por filial e codigo de integracao
func FindLogsFilialIntegracaoEndPoint(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)
	logs, err := logDAO.FindByIntegracaoFilial(params["integracao"], params["codigo"])
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, logs)
}

//FindLogsGroupIntegracao return todos agrupando por codigo integracao
func FindLogsGroupIntegracao(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	logs, err := logDAO.FindGroupIntegracao()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, logs)
}

//FindLogsGroupFilial return todos agrupado por filial e tipo
func FindLogsGroupFilial(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	logs, err := logDAO.FindGroupFilialTipo()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, logs)
}

//FindLogEndpoint find a filial
func FindLogEndpoint(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)
	filial, err := logDAO.FindByID(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "ID invalido")
		return
	}
	respondWithJSON(w, http.StatusOK, filial)
}

//CreateLogEndPoint cria novo registro
func CreateLogEndPoint(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	defer r.Body.Close()
	var log model.LogMessage
	if err := json.NewDecoder(r.Body).Decode(&log); err != nil {
		respondWithError(w, http.StatusBadRequest, "Payload inválido")
		return
	}
	if len(log.DescricaoErro) > 400 {
		log.DescricaoErro = string(log.DescricaoErro[0:400])
	}
	if len(log.ConteudoMensagemErro) > 10000000 {
		log.ConteudoMensagemErro = string(log.ConteudoMensagemErro[0:10000000])
	}
	if err := logDAO.Insert(log); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, log)
}

//UpdateLogEndPoint update existing movie
func UpdateLogEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var log model.LogMessage
	if err := json.NewDecoder(r.Body).Decode(&log); err != nil {
		respondWithError(w, http.StatusBadRequest, "Payload inválido")
		return
	}
	if err := logDAO.Update(log); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

//DeleteLogEndPoint exclui registro
func DeleteLogEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var log model.LogMessage
	if err := json.NewDecoder(r.Body).Decode(&log); err != nil {
		respondWithError(w, http.StatusBadRequest, "Payload inválido")
		return
	}
	if err := logDAO.Delete(log); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"error": msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// func setupResponse(w *http.ResponseWriter, req *http.Request) {
// 	(*w).Header().Set("Access-Control-Allow-Origin", "*")
// 	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
// 	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
// }

// func indexHandler(w http.ResponseWriter, req *http.Request) {
// 	setupResponse(&w, req)
// 	if (*req).Method == "OPTIONS" {
// 		return
// 	}
// }
