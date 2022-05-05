#!/bin/bash
#// with cache
goctl model mysql ddl -src="model/*.sql" -dir="model" -c -style goZero
#// without Redis cache
#goctl model mysql ddl -src="model/*.sql" -dir="model" -style goZero 
