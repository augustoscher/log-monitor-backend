package model

//LogGroupIntegracao representa log agrupado por integracao
type LogGroupIntegracao struct {
	CodigoIntegracao string `sql:"codigo_integracao" json:"codigo_integracao"`
	NomeIntegracao   string `sql:"nome_integracao" json:"nome_integracao"`
	Qtd              string `sql:"qtd" json:"qtd"`
}
