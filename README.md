# archdesc-apis
Archival descriptions Management System - APIs (go-zero)

## go-zero APIs

### Template 

- https://go-zero.dev/en/api-grammar.html

### API directory structure

- (source: https://go-zero.dev/en/api-dir.html )
```
.
├── etc
│   └── greet-api.yaml              // yaml configuration file
├── go.mod                          // go module file
├── greet.api                       // api interface description language file
├── greet.go                        // main function entry
└── internal                        
    ├── config  
    │   └── config.go               // configuration declaration type
    ├── handler                     // routing and handler forwarding
    │   ├── greethandler.go
    │   └── routes.go
    ├── logic                       // business logic
    │   └── greetlogic.go
    ├── middleware                  // middleware file
    │   └── greetmiddleware.go
    ├── svc                         // the resource pool that logic depends on
    │   └── servicecontext.go
    └── types                       // The struct of request and response is automatically generated according to the api, and editing is not recommended
        └── types.go
```
