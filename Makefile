#run-db:
#	docker run -d -p 27017:27017 --name=go-mongo mongo:latest
run-db:
	docker-compose up -d