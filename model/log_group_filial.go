package model

//LogGroupFilialTipo representa os logs agrupados por filial, integracao e tipo
type LogGroupFilialTipo struct {
	CodigoIntegracao string `sql:"codigo_integracao" json:"codigo_integracao"`
	NomeIntegracao   string `sql:"nome_integracao" json:"nome_integracao"`
	CodigoFilial     string `sql:"codigo_filial" json:"codigo_filial"`
	NomeFilial       string `sql:"nome_filial" json:"nome_filial"`
	TipoNotificacao  string `sql:"tipo_notificacao" json:"tipo_notificacao"`
	Qtd              string `sql:"qtd" json:"qtd"`
}
