package logic

import (
	"context"

	"archdesc-apis/app/taxonomy/cmd/api/internal/svc"
	"archdesc-apis/app/taxonomy/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTaxonomyLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTaxonomyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaxonomyLogic {
	return &GetTaxonomyLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTaxonomyLogic) GetTaxonomy(req *types.ReqTaxonomyId) (resp *types.Taxonomy, err error) {
	// todo: add your logic here and delete this line

	return
}
