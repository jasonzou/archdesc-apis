package taxonomy

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
	m := map[int64]string{
		1: "test",
		2: "test2",
	}

	nickname := "unknown"

	if name, ok := m[req.Id]; ok {
		nickname = name
	}

	return &types.Taxonomy{
		Id:   req.Id,
		Term: nickname,
	}, nil
}
