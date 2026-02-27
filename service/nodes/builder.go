package nodes

import (
	"encoding/json"

	"github.com/yao560909/emqx-admin-go-sdk/core"
)

// ------------------------
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

type GetNodeReq struct {
	apiReq *core.APIReq
}
type GetNodeReqBuilder struct {
	apiReq *core.APIReq
}

func NewGetNodeReqBuilder() *GetNodeReqBuilder {
	builder := &GetNodeReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *GetNodeReqBuilder) Node(node string) *GetNodeReqBuilder {
	b.apiReq.PathParams.Set("node", node)
	return b
}

func (b *GetNodeReqBuilder) Build() *GetNodeReq {
	req := &GetNodeReq{}
	req.apiReq = b.apiReq
	return req
}

type GetNodeResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Node
}

type GetNodeMetricsReq struct {
	apiReq *core.APIReq
}

type GetNodeMetricsReqBuilder struct {
	apiReq *core.APIReq
}

func NewGetNodeMetricsReqBuilder() *GetNodeMetricsReqBuilder {
	builder := &GetNodeMetricsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *GetNodeMetricsReqBuilder) Node(node string) *GetNodeMetricsReqBuilder {
	b.apiReq.PathParams.Set("node", node)
	return b
}

func (b *GetNodeMetricsReqBuilder) Build() *GetNodeMetricsReq {
	req := &GetNodeMetricsReq{}
	req.apiReq = b.apiReq
	return req
}

type GetNodeMetricsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	NodeMetrics
}

type GetNodeStatsReq struct {
	apiReq *core.APIReq
}

type GetNodeStatsReqBuilder struct {
	apiReq *core.APIReq
}

func NewGetNodeStatsReqBuilder() *GetNodeStatsReqBuilder {
	builder := &GetNodeStatsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *GetNodeStatsReqBuilder) Node(node string) *GetNodeStatsReqBuilder {
	b.apiReq.PathParams.Set("node", node)
	return b
}

func (b *GetNodeStatsReqBuilder) Build() *GetNodeStatsReq {
	req := &GetNodeStatsReq{}
	req.apiReq = b.apiReq
	return req
}

type GetNodeStatsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	NodeStats
}
