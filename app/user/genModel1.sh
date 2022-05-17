#!/bin/bash
tables==$2
modeldir=./genModel

host=127.0.0.1
port=3306
dbname=$1
username=root
passwd=hewlett1


echo "$dbname DB - Table $2"
goctl model mysql datasource -url="${username}:${passwd}@tcp(${host}:${port})/${dbname}" -table="${tables}"  -dir="${modeldir}" -cache=true --style=goZero
