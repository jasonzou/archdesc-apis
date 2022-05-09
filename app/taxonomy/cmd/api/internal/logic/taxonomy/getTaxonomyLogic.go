package taxonomy

import (
	"context"
	"errors"

	"archdesc-apis/app/taxonomy/cmd/api/internal/svc"
	"archdesc-apis/app/taxonomy/cmd/api/internal/types"
	"archdesc-apis/app/taxonomy/cmd/rpc/pb/taxonomyservice"

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
	taxonomyResp, err := l.svcCtx.TaxonomyRpcClient.Get(l.ctx, &taxonomyservice.ReqTaxonomyId{Id: req.Id})
	if err != nil {
		logx.Error(err.Error())
		return nil, errors.New("errorrrrrrrrrrrrrrrrrrrr")
	}

	return &types.Taxonomy{
		Id:   taxonomyResp.Id,
		Term: "hello",
	}, nil
	/*
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
	*/
}
