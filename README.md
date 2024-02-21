# API LOJA
Projeto criado, em go e html, para simular operações de entrada, saída, busca e listagem de produtos de uma loja.

#### Versão do Go utilizada:
> go version go1.22.0 linux/amd64

#### Versão do PostgreSQL utilizada:
> psql (PostgreSQL) 16.2 (Ubuntu 16.2-1.pgdg22.04+1)

### Reconstruir as imagens e iniciar o container
    docker compose up --build

### Iniciar o container já criado
    docker compose up

### Parar o container
    docker compose stop

### Compilar o programa Go
    go build -o api main.go
