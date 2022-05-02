#!/bin/bash
goctl model mysql ddl -src="model/*.sql" -dir="model" -c
