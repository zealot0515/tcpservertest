#!/bin/sh
#kill demoapi
kill `ps xu | grep demoapi | grep -v grep | awk '{print \$2}'`
#kill tcpservertest
kill `ps xu | grep tcpservertest | grep -v grep | awk '{print \$2}'`