#!/bin/sh
cd "$(dirname "$0")"

#start api server
cd demoapi
./demoapi > /dev/null 2>&1 &

#start tcp server
cd ../server
./tcpservertest > tcpservertestLog 2>&1 &

#start 10 bot, delay 3s
cd ../botclient
sleep 3
./start_10_bot.sh
