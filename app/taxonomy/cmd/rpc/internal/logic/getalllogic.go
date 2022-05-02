package logic

import (
	"context"

	"archdesc-apis/app/taxonomy/cmd/rpc/internal/svc"
	"archdesc-apis/app/taxonomy/cmd/rpc/pb/taxonomy"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAllLogic {
	return &GetAllLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAllLogic) GetAll(in *taxonomy.ReqGetAll) (*taxonomy.RespGetAll, error) {
	// todo: add your logic here and delete this line

	return &taxonomy.RespGetAll{}, nil
}
