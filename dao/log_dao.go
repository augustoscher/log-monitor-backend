package dao

import (
	"bitbucket.org/augustoscher/logs-monitor-docker-postgres/model"
)

//LogDAO representa a conexao do server e banco
type LogDAO struct {
}

//FindAll retorna todos
func (m *LogDAO) FindAll() ([]model.LogMessage, error) {
	var logsMessage []model.LogMessage
	err := db.Model(&logsMessage).Order("id ASC").Select()
	return logsMessage, err
}

//FindByID busca por id
func (m *LogDAO) FindByID(id string) (model.LogMessage, error) {
	logMessage := model.LogMessage{
		ID: id,
	}
	err := db.Select(&logMessage)
	return logMessage, err
}

// Insert adiciona novo
func (m *LogDAO) Insert(logMessage model.LogMessage) error {
	_, err := db.Model(&logMessage).Insert()
	return err
}

// Delete deleta um registro
func (m *LogDAO) Delete(logMessage model.LogMessage) error {
	_, err := db.Model(logMessage).WherePK().Delete()
	return err
}

// Update atualiza um registro existente
func (m *LogDAO) Update(logMessage model.LogMessage) error {
	err := db.Update(logMessage)
	return err
}
