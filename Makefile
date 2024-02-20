up: # Sobe o container
	docker-compose up --build

stop: # Para o container
	docker-compose -f stop

build: # Faz o build
	go build -o api main.go

run: # Inicia o código
	go run main.go