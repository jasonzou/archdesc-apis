# API calls RPC and DB

Basically, api services are working as RPC clients to connect to RPC servers. So API services need to know where to find the RPC servers,
initiate RPC clients, and make RPC calls via RPC clients.

## Config api/etc/taxonomy-api.yaml

- Edit [api/etc/taxonomy-api.yaml](https://github.com/jasonzou/archdesc-apis/blob/v0.1.3/app/taxonomy/cmd/api/etc/taxonomy-api.yaml) and add the following lines (connect to rpc server directly) and add DB and CacheEdit config.go to refer to the TaxonomyRpcConf

## Add DB/Cache into the config

## Add DB model/Cache into the serviceContext

## Internal/logic

- Edit the internal/logic
