### Introduction
[`tuhuebot`](https://github.com/yeungon/tuhuebot) is a Telegram bot written in Go that offers various functionalities for interacting with users on Telegram. It is designed to be easily deployable and maintainable, with support for building and running through a `Makefile` and a `systemd` service for automatic management.

### How to Deploy `tuhuebot` Using `systemd`

1. **Clone the Repository:**
   `git clone https://github.com/yeungon/tuhuebot.git`
   `cd tuhuebot`

2. **Install dependencies:**
    `go mod tidy`
3. **Build the binary:**
    `make build`
4. **Configure the service file yourselt and copy to system:**
    Change the content of tuhuebot.service if needed (location of the code, for example).
    `sudo cp tuhuebot.service /etc/systemd/system/`
5. **Enable and Start the Service:**
    `sudo systemctl enable tuhuebot`
    `sudo systemctl start tuhuebot`
6. **Check Service Status: Verify that the bot is running correctly:**
    `sudo systemctl status tuhuebot`


