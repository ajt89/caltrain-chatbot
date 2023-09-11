compile:
	- go build -v -o app

run:
	- ./app

docker-build:
	- docker build -t caltrain-chatbot:latest .

docker-run:
	- docker run --rm caltrain-chatbot:latest
