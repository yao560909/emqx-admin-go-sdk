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

type ListenerStatusReq struct {
	apiReq *core.APIReq
}

type ListenerStatusReqBuilder struct {
	apiReq *core.APIReq
}

func NewListenerStatusReqBuilder() *ListenerStatusReqBuilder {
	builder := &ListenerStatusReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *ListenerStatusReqBuilder) Build() *ListenerStatusReq {
	req := &ListenerStatusReq{}
	req.apiReq = b.apiReq
	return req
}

type ListenerStatusResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*ListenerStatus `json:"-"`
}

func (resp *ListenerStatusResp) UnmarshalJSON(b []byte) error {
	var listenerStatus []*ListenerStatus
	if err := json.Unmarshal(b, &listenerStatus); err == nil {
		resp.Data = listenerStatus
		return nil
	}
	type alias ListenerStatusResp
	return json.Unmarshal(b, (*alias)(resp))
}
