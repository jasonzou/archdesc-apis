# Redis Cache

## Create a docker-compose.yml for a redis server

- Use the official redis:6.2-alpine docker image in the docker-compose.yml

## Generate the model codes with cache support

- genModel.sh
- Edit the etc/taxonomy-api.yaml and add the following lines
```
CacheRedis:
  - Host: 127.0.0.1:36379
    Pass: G62m50oigInC30sf
    Type: node
```
- Edit the api/internal/config/config.go and add the following line to refer to the CacheRedis
```	
CacheRedis cache.CacheConf
```
- Modify the internal/svc/serviceContext.go and change the codes relate to the Models
  