package controller

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/augustoscher/logs-monitor-docker-postgres/dao"
	"bitbucket.org/augustoscher/logs-monitor-docker-postgres/model"
	"github.com/gorilla/mux"
)

var logDAO = dao.LogDAO{}

//AllLogsEndPoint return todos
func AllLogsEndPoint(w http.ResponseWriter, r *http.Request) {
	logs, err := logDAO.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, logs)
}

//FindLogsGroupIntegracao return todos agrupando por codigo integracao
func FindLogsGroupIntegracao(w http.ResponseWriter, r *http.Request) {
	logs, err := logDAO.FindGroupIntegracao()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, logs)
}

//FindLogsGroupFilial return todos agrupado por filial e tipo
func FindLogsGroupFilial(w http.ResponseWriter, r *http.Request) {
	logs, err := logDAO.FindGroupFilialTipo()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, logs)
}

//FindLogEndpoint find a filial
func FindLogEndpoint(w http.ResponseWriter, r *http.Request) {
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
	defer r.Body.Close()
	var log model.LogMessage
	if err := json.NewDecoder(r.Body).Decode(&log); err != nil {
		respondWithError(w, http.StatusBadRequest, "Payload inválido")
		return
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
