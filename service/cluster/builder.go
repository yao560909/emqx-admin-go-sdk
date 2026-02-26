package cluster

import (
	"github.com/yao560909/emqx-admin-go-sdk/core"
)

type ListClusterReq struct {
	apiReq *core.APIReq
}

type ListClusterReqBuilder struct {
	apiReq *core.APIReq
}

// NewListClusterReqBuilder creates a new ListClusterReqBuilder
func NewListClusterReqBuilder() *ListClusterReqBuilder {
	builder := &ListClusterReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

// Build builds the ListClusterReq
func (b *ListClusterReqBuilder) Build() *ListClusterReq {
	req := &ListClusterReq{}
	req.apiReq = b.apiReq
	return req
}

type ListClusterResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Cluster
}
