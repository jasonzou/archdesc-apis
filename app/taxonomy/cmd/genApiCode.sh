#!/bin/bash
# without cache
#goctl api go -api api/desc/$1.api -dir api/. 
# with cache
goctl api go -api api/desc/$1.api -dir api/. -style goZero -c
