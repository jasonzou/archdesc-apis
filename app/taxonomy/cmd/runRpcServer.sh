#!/bin/bash
pushd rpc
go run taxonomy.go -f etc/taxonomy.yaml &
popd
