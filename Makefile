include .env

compile:
	- go mod tidy; \
	go build -v -o app

format:
	- gofmt -w .

run:
	- export DISCORD_TOKEN=$(DISCORD_TOKEN); \
	./app

docker-build:
	- docker build -t caltrain-chatbot:latest .

docker-run:
	- docker run --rm -e DISCORD_TOKEN=$(DISCORD_TOKEN) caltrain-chatbot:latest
