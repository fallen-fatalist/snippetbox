APP_NAME=snippetbox
PORT=8080

run:
	go run . --port $(PORT)

build: 
	go build -o ./bin/$(APP_NAME) .

clean:
	rm bin/*

up:
	docker-compose -f deploy/docker-compose.yaml up -d

down:
	docker-compose -f deploy/docker-compose.yaml down --volumes

logs:
	docker-compose -f deploy/docker-compose.yaml logs
