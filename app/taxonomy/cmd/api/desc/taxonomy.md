
### 1. "Get a taxonomy name"

1. 路由定义

- Url: /taxonomies/:Id
- Method: GET
- Request: `reqTaxonomyId`
- Response: `taxonomy`

2. 请求定义


```golang
type ReqTaxonomyId struct {
	Id int64 `json:"id"`
}
```


3. 返回定义


```golang
type Taxonomy struct {
	Id int64 `json:"id"`
	Term string `json:"term"`
}
```
  


### 2. "Get names of all taxonomies"

1. 路由定义

- Url: /taxonomies
- Method: GET
- Request: `-`
- Response: `taxonomyListResp`

2. 请求定义


3. 返回定义


```golang
type TaxonomyListResp struct {
	List []taxonomy `json:"list"`
}
```
  

