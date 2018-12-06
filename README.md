# tcpservertest
this is demo project for take home test

i create a server with socket, receive client's cmd and send request to other server via http request

server status can be check at serverhost:8080/info

there are three binarys in bin folder, can use it to make a quick run at linux x64 system
# QuickRun
```
linux x64
exec runtest.sh in bin folder
```
```
windows
exec runtest.bat in bin folder
```
it will create local demoapi, tcpserver, and 10 botclient to test
check http://localhost:8080/info 
can see it's status

and i create a server on gcp, 
http://pagitcptest.ddns.net:8080/info 
can see it's status

and also can change config let botclient connect to or let server send httprequest to that server

## ServerConfig
see bin/server/conf.json
```
{
    "tcpport":6666,					server listen port
    "apihost":"127.0.0.1:7777"		apihost
}
```

## DemoApiConfig
see bin/demoapi/conf/app.conf
```
appname = demo
httpport = 7777			apihost listen port
runmode = release
```

## BotClient
it's just a commandline bot, so it can use without shell
see bin/botclient/start_10_bot.sh
```
#!/bin/sh
cd "$(dirname "$0")"
./simplebot startmultibot 127.0.0.1:6666 sockettest 10
#						  serverhost:port  bot_use_state  bot_create_count
```
