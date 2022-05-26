// Code generated by goctl. DO NOT EDIT!
// Source: taxonomyservice.proto

package taxonomyservice

import (
	"context"

	"archdesc-apis/app/taxonomy/cmd/rpc/pb/taxonomyservice"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ReqGetAll     = taxonomyservice.ReqGetAll
	ReqTaxonomyId = taxonomyservice.ReqTaxonomyId
	RespGetAll    = taxonomyservice.RespGetAll
	Taxonomy      = taxonomyservice.Taxonomy

	TaxonomyService interface {
		Get(ctx context.Context, in *ReqTaxonomyId, opts ...grpc.CallOption) (*Taxonomy, error)
		GetAll(ctx context.Context, in *ReqGetAll, opts ...grpc.CallOption) (*RespGetAll, error)
	}

	defaultTaxonomyService struct {
		cli zrpc.Client
	}
)

func NewTaxonomyService(cli zrpc.Client) TaxonomyService {
	return &defaultTaxonomyService{
		cli: cli,
	}
}

func (m *defaultTaxonomyService) Get(ctx context.Context, in *ReqTaxonomyId, opts ...grpc.CallOption) (*Taxonomy, error) {
	client := taxonomyservice.NewTaxonomyServiceClient(m.cli.Conn())
	return client.Get(ctx, in, opts...)
}

func (m *defaultTaxonomyService) GetAll(ctx context.Context, in *ReqGetAll, opts ...grpc.CallOption) (*RespGetAll, error) {
	client := taxonomyservice.NewTaxonomyServiceClient(m.cli.Conn())
	return client.GetAll(ctx, in, opts...)
}