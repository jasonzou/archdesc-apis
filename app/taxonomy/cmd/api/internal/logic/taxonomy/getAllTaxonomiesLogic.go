package taxonomy

import (
	"context"
	"errors"

	"archdesc-apis/app/taxonomy/cmd/api/internal/svc"
	"archdesc-apis/app/taxonomy/cmd/api/internal/types"
	"archdesc-apis/app/taxonomy/model"

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

func (l *GetAllTaxonomiesLogic) GetAllTaxonomies() (resp *types.TaxonomyListResp, err error) {
	logx.Info("hello get all taxonomies")
	l.Logger.Error("---------- 1 ------------")
	taxList, err := l.svcCtx.TaxonomyModel.FindAll(l.ctx)
	l.Logger.Error("---------- 2 ------------")
	if err != nil && err != model.ErrNotFound {
		l.Logger.Error(err.Error())
		return nil, errors.New("taxonomy list not found")
	}

	if taxList == nil {
		return nil, errors.New("this taxonomy does not exist")
	}
	var tList []types.Taxonomy

	l.Logger.Error(len(taxList))
	if len(taxList) > 0 {

		for _, taxonomy := range taxList {

			var t types.Taxonomy
			_ = copier.Copy(&t, taxonomy)

			tList = append(tList, t)
		}
	}
	return &types.TaxonomyListResp{
		List: tList,
	}, nil
}
