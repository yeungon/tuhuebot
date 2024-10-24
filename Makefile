run: build
	@ ./bin/tuhuebot
build:
	@go build -o bin/tuhuebot -buildvcs=false
dev:
	go run main.go
start:
	sudo systemctl start tuhuebot
stop:
	sudo systemctl stop tuhuebot
restart:
	sudo systemctl restart tuhuebot
status:
	sudo systemctl status tuhuebot
#Print out the log after running with systemd
log:
	journalctl -u tuhuebot -f