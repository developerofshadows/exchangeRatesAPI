.PHONY: dc run test lint

dc:
	docker-compose up  --remove-orphans --build

run:
	go build -o app.exe cmd/rate/main.go && .\app.exe

test:
	go test -race ./...

lint:
	golangci-lint run