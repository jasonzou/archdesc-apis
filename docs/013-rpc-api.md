# API calls RPC

Basically, api services are working as RPC clients to connect to RPC servers. So API services need to know where to find the RPC servers,
initiate RPC clients, and make RPC calls via RPC clients.

## Config api/etc/taxonomy-api.yaml

- Edit [api/etc/taxonomy-api.yaml](https://github.com/jasonzou/archdesc-apis/blob/v0.1.3/app/taxonomy/cmd/api/etc/taxonomy-api.yaml) and add the following lines (connect to rpc server directly):
```
TaxonomyRpcConf:
  Endpoints:
    - 127.0.0.1:19999
```

## Edit config.go to refer to the TaxonomyRpcConf

- Add TaxonomyRpcConf into the config
```
TaxonomyRpcConf zrpc.RpcClientConf
```

## Add rpc into the serviceContext

## Genereate Code

- genRpcCode.sh

## Add Logic

- Edit the internal/logic
