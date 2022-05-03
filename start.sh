# start etcd
etcd &

. ./env.sh

# start taxonomy rpc server
pushd $project/app/taxonomy/cmd/rpc
go run taxonomy.go -f etc/taxonomy.yaml &
popd

# start taxonomy api server
pushd $project/app/taxonomy/cmd/api
go run taxonomy.go -f etc/taxonomy-api.yaml &
popd
