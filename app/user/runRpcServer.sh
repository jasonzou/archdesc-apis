#!/bin/bash
pushd rpc
go run taxonomyservice.go -f etc/taxonomyservice.yaml &
popd
