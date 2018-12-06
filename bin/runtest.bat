REM start api server
cd demoapi
start ./demoapi.exe

REM start tcp server
cd ../server
start ./tcpservertest 

REM start 10 bot, delay 3s
cd ../botclient
timeout 3
start ./simplebot startmultibot 127.0.0.1:6666 sockettest 10