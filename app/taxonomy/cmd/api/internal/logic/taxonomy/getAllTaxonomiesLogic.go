package taxonomy

import (
	"context"

	"archdesc-apis/app/taxonomy/cmd/api/internal/svc"
	"archdesc-apis/app/taxonomy/cmd/api/internal/types"
	"archdesc-apis/app/taxonomy/cmd/rpc/pb/taxonomyservice"

	"github.com/golang-jwt/jwt"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllTaxonomiesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetAllTaxonomiesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllTaxonomiesLogic {
	return &GetAllTaxonomiesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAllTaxonomiesLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func (l *GetAllTaxonomiesLogic) GetAllTaxonomies() (resp *types.TaxonomyListResp, err error) {
	taxonomyResp, err := l.svcCtx.TaxonomyRpcClient.GetAll(l.ctx, &taxonomyservice.ReqGetAll{})
	if err != nil {
		return nil, err
	}
	var listTaxonmoy []types.Taxonomy
	for _, tempTax := range taxonomyResp.Taxonomies {
		logx.Info(tempTax)
		var t types.Taxonomy
		_ = copier.Copy(&t, tempTax)
		listTaxonmoy = append(listTaxonmoy, t)
	}

	return &types.TaxonomyListResp{
		List: listTaxonmoy,
	}, nil

}
