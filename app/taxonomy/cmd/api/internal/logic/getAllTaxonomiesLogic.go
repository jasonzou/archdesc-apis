package logic

import (
	"context"

	"archdesc-apis/app/taxonomy/cmd/api/internal/svc"
	"archdesc-apis/app/taxonomy/cmd/api/internal/types"

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

func (l *GetAllTaxonomiesLogic) GetAllTaxonomies() (resp *types.TaxonomyListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
