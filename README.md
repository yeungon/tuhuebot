
## Introduction

[`tuhuebot`](https://github.com/yeungon/tuhuebot) is a Telegram bot written in Go that offers various functionalities for interacting with users on Telegram. It is designed to be easily deployable and maintainable, with support for building and running through a `Makefile` and a `systemd` service for automatic management.

## How to Deploy `tuhuebot` Using `systemd`
 
1.  **Clone the Repository:**

`git clone https://github.com/yeungon/tuhuebot.git`

`cd tuhuebot`

  2.  **Install dependencies:**

`go mod tidy`

3.  **Build the binary:**

`make build`

4.  **Configure the service file yourselt and copy to system:**

Change the content of tuhuebot.service if needed (location of the code, for example).

`sudo cp tuhuebot.service /etc/systemd/system/`

5.  **Enable and Start the Service:**

`sudo systemctl enable tuhuebot`

`sudo systemctl start tuhuebot`

6.  **Check Service Status: Verify that the bot is running correctly:**

`sudo systemctl status tuhuebot`
 
## Database

Currently `@tuhuebot` is using sqlite and postgresql database. Cache implemented via bigCache.

## Update the code and re-deploy

Here’s a short guide for updating and redeploying your `tuhuebot` project:

1. **Pull Latest Changes**  
Login to your server, go to the folder where the code is put. First, we update your local repository with the latest code from GitHub:
	`sudo git pull`

2. **Build the Project**  
   Compile the code and prepare it for deployment:
	`make build`

3. **Restart the Service**  
   Use `make` to restart the systemd service for the changes to take effect:
	`make restart`

This process will update your repository, build the project, and restart the service to apply any new changes.