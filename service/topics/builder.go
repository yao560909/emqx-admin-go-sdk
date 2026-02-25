package topics

import (
	"encoding/json"
	"fmt"

	"github.com/yao560909/emqx-admin-go-sdk/core"
)

// -------------------------
type ListTopicsReq struct {
	apiReq *core.APIReq
}

type ListTopicsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*Topic `json:"data"`
	Meta *Meta    `json:"meta"`
}

type ListTopicsReqBuilder struct {
	apiReq *core.APIReq
}

func NewListTopicsReqBuilder() *ListTopicsReqBuilder {
	builder := &ListTopicsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *ListTopicsReqBuilder) Build() *ListTopicsReq {
	req := &ListTopicsReq{}
	req.apiReq = b.apiReq
	return req
}

func (b *ListTopicsReqBuilder) Limit(limit int64) *ListTopicsReqBuilder {
	b.apiReq.QueryParams.Set("limit", fmt.Sprint(limit))
	return b
}

func (b *ListTopicsReqBuilder) Page(page int64) *ListTopicsReqBuilder {
	b.apiReq.QueryParams.Set("page", fmt.Sprint(page))
	return b
}

func (b *ListTopicsReqBuilder) Node(node string) *ListTopicsReqBuilder {
	b.apiReq.QueryParams.Set("node", node)
	return b
}

func (b *ListTopicsReqBuilder) Topic(topic string) *ListTopicsReqBuilder {
	b.apiReq.QueryParams.Set("topic", topic)
	return b
}

// -------------------------
type GetTopicReq struct {
	apiReq *core.APIReq
}

type GetTopicResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*Topic `json:"-"`
}

func (resp *GetTopicResp) UnmarshalJSON(b []byte) error {
	var topics []*Topic
	if err := json.Unmarshal(b, &topics); err == nil {
		resp.Data = topics
		return nil
	}
	type alias GetTopicResp
	return json.Unmarshal(b, (*alias)(resp))
}

type GetTopicReqBuilder struct {
	apiReq *core.APIReq
}

func NewGetTopicReqBuilder() *GetTopicReqBuilder {
	builder := &GetTopicReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *GetTopicReqBuilder) Topic(topic string) *GetTopicReqBuilder {
	b.apiReq.PathParams.Set("topic", topic)
	return b
}

func (b *GetTopicReqBuilder) Build() *GetTopicReq {
	req := &GetTopicReq{}
	req.apiReq = b.apiReq
	return req
}
