# archdesc-apis

A step by step approach to learn [go-zero](https://github.com/zeromicro/go-zero) via watching [tutorial videos](https://space.bilibili.com/389552232/video) and studying [the code](https://github.com/Mikaelemmmm/go-zero-looklook). A small portion of an archival descriptions management system is building along the learning process. 

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

## A to do list

- terms of a taxonomy apis
- information object apis
- accession/deaccession apis
- static pages apis -

## Things to consider

- check Omeka-s (site/settings etc.)  - checked. Site is a better implementation to encapsulate a public interface (static pages, theme)
- taxonomy apis (getTaxonomy/getAllTaxonomy)
- country code/AtoM 2 taxonomy code
