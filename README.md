# log-monitor-backend
Golang API that's responsible for receiving logs of applications.  
PostgreSQL version.

### Ambientes
1) Ambiente Docker  
2) Instalação Manual.  

### 1- Ambiente Docker
a) Baixar e instalar a última versão do docker.  
b) Realizar o clone do repositório.  
c) Através do terminal, acessar o diretório clonado.  
d) Executar: docker-compose up -d  
e) Executar comando para visualizar os logs: docker-compose logs -f -t   
f) Ou então, tudo de uma vez: docker-compose up -d && docker-compose logs -f -t   

Este processo gerará dois containers e através deles, poderemos inciar e parar as aplicações de maneira independente.  
Para listar os containers: docker container ps ou docker container ps -a  
Serão listados os containers da seguinte forma:  

|CONTAINER ID | IMAGE | COMMAND | CREATED | STATUS | PORTS | NAMES |
|-------------|-------|---------|---------|--------|-------|-------|
|218c3fa8c2ed|postgres:9.6|"docker-entrypoint.s…"|About an hour ago|Up About an hour| 0.0.0.0:5433->5432/tcp|logs-monitor-docker-postgres_db_1
|b0fcfa30606a|golang:1.11.5|"bash -c 'cd /usr/lo…"|About an hour ago|Up About an hour|0.0.0.0:3007->3007/tcp|logs-monitor-docker-postgres_backend_1

a) Container de banco de dados: postgreSQL  
Para testar se o banco está online: http://localhost:5433  
Para verificar se o schema foi criado corretamente, executar dentro da pasta do projeto: docker-compose exec db psql -U postgres -f /scripts/check.sql  
É possível criar uma conexão através do PGAdmin, por que existe um mapeamento da porta interna do container para a porta externa (5433).  
Para conectar no banco via pgAdmin:  
- Via PGAdmin, criar um novo Server.  
- Dar um nome ao server. Pode ser: 'logs-monitor-docker-postgres-server'.  
- Na aba Connections, informar em Host: 'localhost' e a porta '5433'.  
- Pronto. Você terá acesso ao banco de dados interno do container.  

b) Container de backend: golang  
Para testar se o backend está online: http://localhost:3007/logs-group-filial

### 2- Instalação Manual  
a) Baixar e instalar o GOlang na versão 1.11.5  
b) Baixar e instalar o PostgreSQL na versão 9.6  
c) Realizar o clone do repositório.  
d) Iniciar o banco PostgreSQL e criar um database. O nome deverá ser "log_monitor".  
e) Criar schema do banco. Vide /script/init.sql  
f) Através do terminal, acessar o diretório clonado.  
g) Iniciar o backend. Comando: go run app.go  

Este processo deverá inciar a API, ficando disponível na porta 3007.  

### 3- Testes
a) GET ALL  
http://localhost:3007/logs?limit=10&offset=0 

b) GET BY ID  
http://localhost:3007/logs/1  

c) GET GROUPED BY INTEGRACAO  
http://localhost:3007/logs-group-integracao

c) GET GROUPED BY FILIAL E INTEGRACAO  
http://localhost:3007/logs-group-filial

d) POST  
http://localhost:3007/logs  
```json
{
	"codigo_filial": "001",
	"nome_filial": "FILIAL-001",
	"tipo_notificacao": "M",
	"codigo_integracao": "100",
	"nome_integracao": "CONVERSION_PROCESSED",
}
```


