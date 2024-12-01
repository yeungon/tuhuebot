# Variables
SERVICE_FILE := tuhuebot.service
INSTALL_PATH := /etc/systemd/system/
# Install the service file to /etc/systemd/system/
copy:
	sudo cp $(SERVICE_FILE) $(INSTALL_PATH)
create:
	sudo mkdir -p bin
	sudo chmod 7777 bin
install: copy reload create build start enable status log
run: build
	@ ./bin/tuhuebot
build: 
	@go build -o bin/tuhuebot -buildvcs=false
dev:
	go run main.go
start:
	sudo systemctl start tuhuebot
enable:
	sudo systemctl enable tuhuebot
stop:
	sudo systemctl stop tuhuebot
restart:
	sudo systemctl restart tuhuebot
status:
	sudo systemctl status tuhuebot
pull:
	sudo git pull
update: pull build restart status log
#Print out the log after running with systemd
log:
	journalctl -u tuhuebot -f