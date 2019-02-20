# logs-monitor-docker-postgre

### Ambientes
1) Ambiente Docker  
2) Instalação Manual.  

### 1- Ambiente Docker
a) Baixar e instalar a última versão do docker.  
b) Realizar o clone do repositório.  
c) Através do Power Shell ou CMD, acessar o diretório clonado.  
d) Executar: docker-compose up -d  
e) Executar comando para visualizar os logs: docker-compose logs -f -t   

Este processo gerará dois containers e através deles, poderemos inciar e parar as aplicações de maneira independente.  
Para listar os containers: docker container ps ou docker container ps -a  
Serão listados os containers da seguinte forma:  

|CONTAINER ID | IMAGE | COMMAND | CREATED | STATUS | PORTS | NAMES |
|-------------|-------|---------|---------|--------|-------|-------|
|218c3fa8c2ed|postgres:9.6|"docker-entrypoint.s…"|About an hour ago|Up About an hour| 0.0.0.0:5433->5432/tcp|logs-monitor-docker-postgres_db_1
|b0fcfa30606a|golang:1.11.5|"bash -c 'cd /usr/lo…"|About an hour ago|Up About an hour|0.0.0.0:3000->3000/tcp|logs-monitor-docker-postgres_db_1

a) Container de banco de dados: mongodb  
Para testar se o banco está online: http://localhost:5433  

b) Container de backend: golang  
Para testar se o backend está online: http://localhost:3000/logs  

### 2- Instalação Manual
a) Baixar e instalar o GOlang na versão 1.11.5  
b) Baixar e instalar o PostgreSQL na versão 9.6  
c) Realizar o clone do repositório.  
d) Mudar a configuração de acesso ao banco de dados localhost  
e) Iniciar o banco PostgreSQL.
f) Através do Power Shell ou CMD, acessar o diretório clonado.  
g) Iniciar o backend. Comando: go run app.go  

Este processo deverá inciar a API, ficando disponível na porta 3000.  

### Testes
a) GET ALL  
http://localhost:3000/logs  

b) GET BY ID  
http://localhost:3000/logs/1  

c) GET GROUPED BY INTEGRACAO
http://localhost:3000/logs-group-integracao

c) GET GROUPED BY FILIAL E INTEGRACAO
http://localhost:3000/logs-group-filial

d) POST  
http://localhost:3000/logs  
{
	"codigo_filial": "1",
	"nome_filial": "Teste",
	"tipo_notificacao": "M",
	"codigo_integracao": "120",
	"nome_integracao": "NOTA_FISCAL_ERP",
	"descricao_erro": "eita",
	"conteudo_mensagem_erro": "NPE2"
}



