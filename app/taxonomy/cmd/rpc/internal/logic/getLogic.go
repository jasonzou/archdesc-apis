package logic

import (
	"context"
	"errors"

	"archdesc-apis/app/taxonomy/cmd/rpc/internal/svc"
	"archdesc-apis/app/taxonomy/cmd/rpc/pb/taxonomyservice"
	"archdesc-apis/app/taxonomy/model"

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
	logx.Error(in.Id)
	logx.Info("hello000000000000000000000000000000000000")

	taxonomy, err := l.svcCtx.TaxonomyI18nModel.FindOne(l.ctx, in.Id)
	l.Logger.Error("-------- 2 ----------")
	if err != nil && err != model.ErrNotFound {
		logx.Errorf(err.Error())
		return nil, errors.New("taxonomy not found")
	}

	if taxonomy == nil {
		logx.Error("Not exist!!!!!!!!!!!!!!!!!!!!!")
		return nil, errors.New("this taxonomy does not exist")
	}
	return &taxonomyservice.Taxonomy{
		Id:   in.Id,
		Term: terms,
	}, nil
}
