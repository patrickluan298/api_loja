CREATE DATABASE lojaDB;

CREATE TABLE produtos(
	id serial primary key,
	nome varchar(50),
	descricao varchar(50),
	preco real,
	quantidade integer
);