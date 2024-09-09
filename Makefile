run: build
	@ ./bin/tuhuebot start
build:
	@go build -o bin/tuhuebot
dev:
	go run main.go
start:
	sudo systemctl start tuhuebot
restart:
	sudo systemctl restart tuhuebot
status:
	sudo systemctl status tuhuebot