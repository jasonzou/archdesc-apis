package taxonomy

import (
	"context"
	"errors"

	"archdesc-apis/app/taxonomy/cmd/api/internal/svc"
	"archdesc-apis/app/taxonomy/cmd/api/internal/types"
	"archdesc-apis/app/taxonomy/model"

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
	l.Logger.Error("-------- 1 ----------")
	taxonomy, err := l.svcCtx.TaxonomyModel.FindOne(l.ctx, req.Id)
	l.Logger.Error("-------- 2 ----------")
	if err != nil && err != model.ErrNotFound {
		logx.Errorf(err.Error())
		return nil, errors.New("taxonomy not found")
	}

	if taxonomy == nil {
		return nil, errors.New("this taxonomy does not exist")
	}

	return &types.Taxonomy{
		Id:   taxonomy.Id,
		Term: taxonomy.Usage.String,
	}, nil
}
