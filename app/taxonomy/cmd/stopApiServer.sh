#!/bin/bash
killall go
PID=`ps -ef|grep go-bu |grep -v grep |awk '{ print $2 }'`
kill -9 $PID
