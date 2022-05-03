#!/bin/bash
pushd api
go run taxonomy.go -f etc/taxonomy-api.yaml &
popd
