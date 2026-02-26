package listeners

import (
	"encoding/json"

	"github.com/yao560909/emqx-admin-go-sdk/core"
)

// --------------------------------
type ListListenersReq struct {
	apiReq *core.APIReq
}

type ListListenersReqBuilder struct {
	apiReq *core.APIReq
}

func NewListListenersReqBuilder() *ListListenersReqBuilder {
	builder := &ListListenersReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *ListListenersReqBuilder) Build() *ListListenersReq {
	req := &ListListenersReq{}
	req.apiReq = b.apiReq
	return req
}

func (b *ListListenersReqBuilder) Type(sourceType string) *ListListenersReqBuilder {
	b.apiReq.QueryParams.Set("type", sourceType)
	return b
}

type ListListenersResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*Listener `json:"-"`
}

func (resp *ListListenersResp) UnmarshalJSON(b []byte) error {
	var listeners []*Listener
	if err := json.Unmarshal(b, &listeners); err == nil {
		resp.Data = listeners
		return nil
	}
	type alias ListListenersResp
	return json.Unmarshal(b, (*alias)(resp))
}
