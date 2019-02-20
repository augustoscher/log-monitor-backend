package model

//LogMessage representa uma log
type LogMessage struct {
	tableName            struct{} `sql:"logmessage"`
	ID                   string   `sql:"id" json:"id"`
	CodigoFilial         string   `sql:"codigo_filial" json:"codigo_filial"`
	NomeFilial           string   `sql:"nome_filial" json:"nome_filial"`
	TipoNotificacao      string   `sql:"tipo_notificacao" json:"tipo_notificacao"`
	CodigoIntegracao     string   `sql:"codigo_integracao" json:"codigo_integracao"`
	NomeIntegracao       string   `sql:"nome_integracao" json:"nome_integracao"`
	DescricaoErro        string   `sql:"descricao_erro" json:"descricao_erro"`
	ConteudoMensagemErro string   `sql:"conteudo_mensagem_erro" json:"conteudo_mensagem_erro"`
	DataHora             string   `sql:"data_hora" json:"data_hora"`
}
