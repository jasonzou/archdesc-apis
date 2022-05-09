package logic

import (
	"context"
	"errors"

	"archdesc-apis/app/taxonomy/cmd/rpc/internal/svc"
	"archdesc-apis/app/taxonomy/cmd/rpc/pb/taxonomyservice"
	"archdesc-apis/app/taxonomy/model"

	"github.com/jinzhu/copier"
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

func (l *GetAllLogic) GetAll(in *taxonomyservice.ReqGetAll) (*taxonomyservice.RespGetAll, error) {
	// todo: add your logic here and delete this line

	logx.Info("RPC hello get all taxonomies")
	l.Logger.Error("RPC ---------- 1 ------------")

	taxList, err := l.svcCtx.TaxonomyModel.FindAll(l.ctx)
	l.Logger.Error("---------- 2 ------------")
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err.Error())
		return nil, errors.New("taxonomy list not found")
	}

	if taxList == nil {
		return nil, errors.New("this taxonomy does not exist")
	}

	var tList []*taxonomyservice.Taxonomy

	l.Logger.Error(len(taxList))
	if len(taxList) > 0 {

		for _, tempTax := range taxList {

			logx.Error(tempTax.Id)
			var t taxonomyservice.Taxonomy
			_ = copier.Copy(&t, tempTax)

			tList = append(tList, &t)
		}
	}

	l.Logger.Error(len(tList))
	return &taxonomyservice.RespGetAll{
		Taxonomies: tList,
	}, nil
}
