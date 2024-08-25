build:
	docker build -t go-books:latest .

local_run:
	docker run --name go-books -d -p 8080:8080 go-books:latest

local_inside:
	docker exec -it go-books /bin/sh

local_stop:
	docker stop go-books
	docker rm go-books