// Code generated by goctl. DO NOT EDIT.
package types

type Taxonomy struct {
	Id   int64  `json:"id"`
	Term string `json:"term"`
}

type TaxonomyListResp struct {
	List []Taxonomy `json:"list"`
}

type ReqTaxonomyId struct {
	Id int64 `json:"id"`
}
