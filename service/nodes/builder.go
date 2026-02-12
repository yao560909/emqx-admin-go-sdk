package nodes

import (
	"encoding/json"

	"github.com/yao560909/emqx-admin-go-sdk/core"
)

type ListNodesReq struct {
	apiReq *core.APIReq
}

type ListNodesResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*Node `json:"-"`
}

func (resp *ListNodesResp) UnmarshalJSON(b []byte) error {
	var nodes []*Node
	if err := json.Unmarshal(b, &nodes); err == nil {
		resp.Data = nodes
		return nil
	}
	type alias ListNodesResp
	return json.Unmarshal(b, (*alias)(resp))
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
