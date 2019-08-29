# Peach
 ##### A Discord server management and chat bot written in Python.
[![GitHub version](https://badge.fury.io/gh/juliscrazy%2FCuddler.svg)](https://github.com/juliscrazy/Cuddler/issues)
[![Open Source Love](https://badges.frapsoft.com/os/mit/mit.svg?v=102)](https://github.com/juliscrazy/Cuddler)
 
## Features
 - Webinterface
   - Bot Statistics
   - Enable and Disable Commands
   - Access logs
 - Easy Command Expansion
 - Easy Feature Integration

## Setup
##### How to host the bot on a raspberrypi / Linux machine
Disclaimer: I usually work on Linux. This is the way I did it. There might be better ways to do it. That's fine, nice for you. I however don't know any better and in this case I'm fine with that.

#### Downloading the code
Clone the Repository:
```
cd ~/Documents
mkdir Apps
cd Apps
git clone https://github.com/juliscrazy/Peach
```
Install the dependencies:
```
cd Peach
python3.7 -m pip install --user -r requirements.txt
```
Now we need to get ourselves a discord bot app and a token. Go to https://discordapp.com/developers/, create a new application, give it a name and a picture. Then in the app settings go to "Bot", create a bot and copy the token.

Back on the pi in `~/Documents/Apps/Peach` create a json file `nano auth.json` and enter:
```
{"TOKEN": "<token>"}
```
Save and exit (replace the `<token>` with the token you copied (duh...) )

#### Setting up the services

To set up the bot we'll need to create 3 system services:
   - one for the tcp server
   - one for the interface
   - and one for the bot itself
   
on your pi / Linux go to `~/.config/systemd/user`
create a new service file (`nano peach_interface.service`)

Then enter this into the service file:
```
[Unit]
# Human readable name of the unit

Description=Peach Bot Interface Service

[Service]
WorkingDirectory=/home/<user>/Documents/Apps/Peach

ExecStart=/usr/bin/python3.7 /home/pi/Documents/Apps/Peach/peach_interface/main.py

Restart=on-abort

[Install]
WantedBy=default.target
```
Save and Exit.....

Repeat this for all 3 services (change the `ExecStart` filepath and the `Description` respectably)
you should end up with:

   - peach_interface.service
   - peach_bot.service
   - peach_tcp_server.service

Next step is to start the services

First enable all 3 services:
`systemctl --user enable peach_tcp_server` (repeat for the other 2)

then reload the daemons:
`systemctl --user daemon`

and then start the services (I recommend starting the tcp server first):
`systemctl --user start peach_tcp_server` (To check the status afterwards: `systemctl --user status peach_tcp_server`)

Your Bot should now be up and running

#### Accessing the interface

You can access the interface through a computer in the same network with a browser (duh)

As URL enter `<ip>:5000` (again, replace the ip with the host machines ip)

## Technologies
 - Flask
   - HTML, CSS, Bootstrap
 - RabbitMQ
