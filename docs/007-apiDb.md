# DB

## Two ways to generate model codes

- From \*.sql files - genModel.sh
- From databases directly - genModel1.sh

### Generate Model Files

- a sql file is associated two model files
  - xxxnameModel.go
    - define xxxnameModel interface which can be extended
    - provide func NewxxxnameModel
  - xxxnameModel_gen.go
    - provide default implentation of xxxnameModel
    - Insert
    - FindOne
    - Update
    - Delete

## Steps to add DB support

- generate model codes
- etc/*.yaml
  - [add DB config info](https://github.com/jasonzou/archdesc-apis/blob/v0.0.7/app/taxonomy/cmd/api/etc/taxonomy-api.yaml#L5-L6)
  - pay attention to the **(localhost:3306)** where the DB is available

    ```
      DataSource: atom:hewlett1@tcp(localhost:3306)/atom?charset=utf8mb4
    ```
- api/internal/config/config.go
  - [add DB config info](https://github.com/jasonzou/archdesc-apis/blob/v0.0.7/app/taxonomy/cmd/api/internal/config/config.go#L9-L10)
- api/internal/svc/servicecontext.go
  - [change type ServiceContext](https://github.com/jasonzou/archdesc-apis/blob/v0.0.7/app/taxonomy/cmd/api/internal/svc/servicecontext.go#L12)
  - [change func NewServiceContext](https://github.com/jasonzou/archdesc-apis/blob/v0.0.7/app/taxonomy/cmd/api/internal/svc/servicecontext.go#L18)
