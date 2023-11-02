build:
	go build -o main .

run:
	go run main.go

test:
	go test -v ./...

docker:
	docker build -t myapp .

docker-run:
	docker run -p 8080:8080 myapp
