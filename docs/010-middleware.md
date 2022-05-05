# Middleware

## Services level middleware

- api/desc/taxonomy.api - Add a definition for middle  L27
- Regenerate the api codes
- Edit api/internal/middleware/taxonomyMiddleware.go

## Global middleware

- Add [common/middleware/globalMiddleware.go](https://github.com/jasonzou/archdesc-apis/blob/v0.1.0/common/middleware/globalMiddleware.go) to the root of the project
- Edit [api/taxonomy.go](https://github.com/jasonzou/archdesc-apis/blob/v0.1.0/app/taxonomy.go#L8) - add the middleware to the server
