package logic

import (
	"context"

	"archdesc-apis/app/taxonomy/cmd/rpc/internal/svc"
	"archdesc-apis/app/taxonomy/cmd/rpc/pb/taxonomyservice"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetLogic) Get(in *taxonomyservice.ReqTaxonomyId) (*taxonomyservice.Taxonomy, error) {
	// todo: add your logic here and delete this line
	logx.Info("hello")
	terms := []string{"test", "fjdsdsl"}

	return &taxonomyservice.Taxonomy{
		Id:   in.Id,
		Term: terms,
	}, nil
}
