run: build
	@ ./bin/tuhuebot start
build:
	@go build -o bin/tuhuebot
dev:
	go run main.go