#!/bin/bash
#// with Redis cache
#//goctl model mysql ddl -src="model/*.sql" -dir="model" -c
#// without Redis cache
goctl model mysql ddl -src="model/*.sql" -dir="model" -style goZero 
