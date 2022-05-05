package taxonomy

import (
	"net/http"

	"archdesc-apis/app/taxonomy/cmd/api/internal/logic/taxonomy"
	"archdesc-apis/app/taxonomy/cmd/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetAllTaxonomiesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := taxonomy.NewGetAllTaxonomiesLogic(r.Context(), svcCtx)
		resp, err := l.GetAllTaxonomies()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
