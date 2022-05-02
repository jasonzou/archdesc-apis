package handler

import (
	"net/http"

	"archdesc-apis/app/taxonomy/cmd/api/internal/logic"
	"archdesc-apis/app/taxonomy/cmd/api/internal/svc"
	"archdesc-apis/app/taxonomy/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
  	"github.com/zeromicro/go-zero/core/logx"
)

func GetTaxonomyHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ReqTaxonomyId
    logx.Error("could not run server: ")
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewGetTaxonomyLogic(r.Context(), svcCtx)
		resp, err := l.GetTaxonomy(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
