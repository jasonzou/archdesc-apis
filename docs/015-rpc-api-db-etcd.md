# API calls RPC and DB with ETCD

Basically, api services are working as RPC clients to connect to RPC servers. So API services need to know where to find the RPC servers,
initiate RPC clients, and make RPC calls via RPC clients.

## Config api/etc/taxonomy-api.yaml

- Edit [api/etc/taxonomy-api.yaml](https://github.com/jasonzou/archdesc-apis/blob/v0.1.3/app/taxonomy/cmd/api/etc/taxonomy-api.yaml) and add the following lines (connect to rpc server directly) and add DB and CacheEdit config.go to refer to the TaxonomyRpcConf and change to the ETCD connection
```
  TaxonomyRpcConf
  	Etcd:
  	  Hosts:
  	  - 127.0.0.1:2379
  	  Key: taxonomyservice.rpc
```


## Config rpc/etc/taxonomy-api.yaml

- Change to the ETCD connection
```
  Etcd:
    Hosts:
    - 127.0.0.1:2379
    Key: taxonomyservice.rpc
```

## Simply way to check your etcd
```
etcdctl get --prefix ""
```