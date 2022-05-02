#!/bin/bash
goctl model mysql ddl -c -src $1.sql -dir .
