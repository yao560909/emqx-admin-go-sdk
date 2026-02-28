package apikeys

import (
	"encoding/json"

	"github.com/yao560909/emqx-admin-go-sdk/core"
)

// --------------------
type CreateAPIKeyReq struct {
	apiReq *core.APIReq
}

type CreateAPIKeyReqBody struct {
	Name      *string `json:"name,omitempty"`
	Expired   *bool   `json:"expired,omitempty"`
	ExpiredAt *string `json:"expired_at,omitempty"`
	Desc      *string `json:"desc,omitempty"`
	Enable    *bool   `json:"enable,omitempty"`
}

type CreateAPIKeyReqBuilder struct {
	apiReq *core.APIReq
}

func NewCreateAPIKeyReqBuilder() *CreateAPIKeyReqBuilder {
	builder := &CreateAPIKeyReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &CreateAPIKeyReqBody{},
		SkipAuth:    true,
	}
	return builder
}

func (b *CreateAPIKeyReqBuilder) Expired(expired bool) *CreateAPIKeyReqBuilder {
	b.apiReq.Body.(*CreateAPIKeyReqBody).Expired = &expired
	return b
}
func (b *CreateAPIKeyReqBuilder) ExpiredAt(expiredAt string) *CreateAPIKeyReqBuilder {
	b.apiReq.Body.(*CreateAPIKeyReqBody).ExpiredAt = &expiredAt
	return b
}

func (b *CreateAPIKeyReqBuilder) Desc(desc string) *CreateAPIKeyReqBuilder {
	b.apiReq.Body.(*CreateAPIKeyReqBody).Desc = &desc
	return b
}

func (b *CreateAPIKeyReqBuilder) Name(name string) *CreateAPIKeyReqBuilder {
	b.apiReq.Body.(*CreateAPIKeyReqBody).Name = &name
	return b
}

func (b *CreateAPIKeyReqBuilder) Enable(enable bool) *CreateAPIKeyReqBuilder {
	b.apiReq.Body.(*CreateAPIKeyReqBody).Enable = &enable
	return b
}

func (b CreateAPIKeyReqBuilder) Build() *CreateAPIKeyReq {
	req := &CreateAPIKeyReq{}
	req.apiReq = b.apiReq
	return req
}

type CreateAPIKeyResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	APIKey
}

// ---------------------------
type DeleteAPIKeyReq struct {
	apiReq *core.APIReq
}

type DeleteAPIKeyReqBuilder struct {
	apiReq *core.APIReq
}

func NewDeleteAPIKeyReqBuilder() *DeleteAPIKeyReqBuilder {
	builder := &DeleteAPIKeyReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		SkipAuth:    true,
	}
	return builder
}

func (b *DeleteAPIKeyReqBuilder) Name(name string) *DeleteAPIKeyReqBuilder {
	b.apiReq.PathParams.Set("name", name)
	return b
}

func (b *DeleteAPIKeyReqBuilder) Build() *DeleteAPIKeyReq {
	req := &DeleteAPIKeyReq{}
	req.apiReq = b.apiReq
	return req
}

type DeleteAPIKeyResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

// ---------------------------------
type ListAPIKeysReq struct {
	apiReq *core.APIReq
}

type ListAPIKeysReqBuilder struct {
	apiReq *core.APIReq
}

func NewListAPIKeysReqBuilder() *ListAPIKeysReqBuilder {
	builder := &ListAPIKeysReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		SkipAuth:    true,
	}
	return builder
}

func (b *ListAPIKeysReqBuilder) Build() *ListAPIKeysReq {
	req := &ListAPIKeysReq{}
	req.apiReq = b.apiReq
	return req
}

type ListAPIKeysResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*APIKey `json:"-"`
}

func (resp *ListAPIKeysResp) UnmarshalJSON(b []byte) error {
	var apikeys []*APIKey
	if err := json.Unmarshal(b, &apikeys); err == nil {
		resp.Data = apikeys
		return nil
	}
	type alias ListAPIKeysResp
	return json.Unmarshal(b, (*alias)(resp))
}

// ---------------------------------
type GetAPIKeyReq struct {
	apiReq *core.APIReq
}

type GetAPIKeyReqBuilder struct {
	apiReq *core.APIReq
}

func NewGetAPIKeyReqBuilder() *GetAPIKeyReqBuilder {
	builder := &GetAPIKeyReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		SkipAuth:    true,
	}
	return builder
}

/*
*
API Key name
*/
func (b *GetAPIKeyReqBuilder) Name(name string) *GetAPIKeyReqBuilder {
	b.apiReq.PathParams.Set("name", name)
	return b
}

func (b *GetAPIKeyReqBuilder) Build() *GetAPIKeyReq {
	req := &GetAPIKeyReq{}
	req.apiReq = b.apiReq
	return req
}

type GetAPIKeyResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	APIKey
}

// ---------------------------
type UpdateAPIKeyReq struct {
	apiReq *core.APIReq
}

type UpdateAPIKeyReqBody struct {
	Expired   *bool   `json:"expired,omitempty"`
	ExpiredAt *string `json:"expired_at,omitempty"`
	Desc      *string `json:"desc,omitempty"`
	Enable    *bool   `json:"enable,omitempty"`
}

type UpdateAPIKeyReqBuilder struct {
	apiReq *core.APIReq
}

func NewUpdateAPIKeyReqBuilder() *UpdateAPIKeyReqBuilder {
	builder := &UpdateAPIKeyReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &UpdateAPIKeyReqBody{},
		SkipAuth:    true,
	}
	return builder
}

func (b *UpdateAPIKeyReqBuilder) Name(name string) *UpdateAPIKeyReqBuilder {
	b.apiReq.PathParams.Set("name", name)
	return b
}

func (b *UpdateAPIKeyReqBuilder) Expired(expired bool) *UpdateAPIKeyReqBuilder {
	b.apiReq.Body.(*UpdateAPIKeyReqBody).Expired = &expired
	return b
}

func (b *UpdateAPIKeyReqBuilder) ExpiredAt(expiredAt string) *UpdateAPIKeyReqBuilder {
	b.apiReq.Body.(*UpdateAPIKeyReqBody).ExpiredAt = &expiredAt
	return b
}

func (b *UpdateAPIKeyReqBuilder) Desc(desc string) *UpdateAPIKeyReqBuilder {
	b.apiReq.Body.(*UpdateAPIKeyReqBody).Desc = &desc
	return b
}

func (b *UpdateAPIKeyReqBuilder) Enable(enable bool) *UpdateAPIKeyReqBuilder {
	b.apiReq.Body.(*UpdateAPIKeyReqBody).Enable = &enable
	return b
}

func (b *UpdateAPIKeyReqBuilder) Build() *UpdateAPIKeyReq {
	req := &UpdateAPIKeyReq{}
	req.apiReq = b.apiReq
	return req
}

type UpdateAPIKeyResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	APIKey
}
