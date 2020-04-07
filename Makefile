run:
	go run main.go

run-compose:
	docker-compose build
	docker-compose up

image:
	docker build . -t docker.io/mtinside/prom-counter:latest

image-push: image
	docker push docker.io/mtinside/prom-counter:latest
