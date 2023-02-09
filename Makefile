# создает контейнер с БД
run-db:
	docker-compose up -d

# удаляет контейнер с БД
rm-db:
	docker stop go-mongo & docker rm go-mogno

# запускает приложение в режиме разработки
run:
	go run cmd/main.go

# собирает в исполняемый файл
build:
	go build cmd/main.go