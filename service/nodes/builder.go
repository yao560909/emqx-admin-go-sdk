package nodes

import "github.com/yao560909/emqx-admin-go-sdk/core"

type ListNodesReq struct {
	apiReq *core.APIReq
}

type ListNodesResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*Node
}

type ListNodesReqBuilder struct {
	apiReq *core.APIReq
}

func NewListNodesReqBuilder() *ListNodesReqBuilder {
	builder := &ListNodesReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *ListNodesReqBuilder) Build() *ListNodesReq {
	req := &ListNodesReq{}
	req.apiReq = b.apiReq
	return req
}
