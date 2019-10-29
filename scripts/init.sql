create database log_monitor;

\c log_monitor

create table logmessage (
  id serial not null primary key,
  codigo_filial varchar(20) not null,
  nome_filial varchar(150) not null,
  tipo_notificacao varchar(1) not null,	
  codigo_integracao varchar(20) not null,
  nome_integracao varchar(150) not null,
  descricao_erro varchar(400),
  conteudo_mensagem_erro varchar(10485760),
  data_hora timestamp not null default current_timestamp
);