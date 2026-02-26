package apikeys

import "github.com/yao560909/emqx-admin-go-sdk/core"

// --------------------
type CreateAPIKeyReq struct {
	apiReq *core.APIReq
}

type CreateAPIKeyReqBody struct {
	Expired *string `json:"expired,omitempty"`
	Note    *string `json:"note,omitempty"`
	Tag     *string `json:"tag,omitempty"`
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
	}
	return builder
}

/*
*
Expiration time, format: rfc3339, for example: 2025-12-31T23:59:59+08:00
*/
func (b *CreateAPIKeyReqBuilder) Expired(expired string) *CreateAPIKeyReqBuilder {
	b.apiReq.Body.(*CreateAPIKeyReqBody).Expired = &expired
	return b
}

/*
*
Note for the API key
*/
func (b *CreateAPIKeyReqBuilder) Note(note string) *CreateAPIKeyReqBuilder {
	b.apiReq.Body.(*CreateAPIKeyReqBody).Note = &note
	return b
}

/*
*
Tag for the API key
*/
func (b *CreateAPIKeyReqBuilder) Tag(tag string) *CreateAPIKeyReqBuilder {
	b.apiReq.Body.(*CreateAPIKeyReqBody).Tag = &tag
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
	}
	return builder
}

/*
*
API Key to delete
*/
func (b *DeleteAPIKeyReqBuilder) APIKey(apiKey string) *DeleteAPIKeyReqBuilder {
	b.apiReq.PathParams.Set("api_key", apiKey)
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
	}
	return builder
}

/*
*
Default: 100
Example: limit=50
Results per page(max 1000)
*/
func (b *ListAPIKeysReqBuilder) Limit(limit string) *ListAPIKeysReqBuilder {
	b.apiReq.QueryParams.Set("limit", limit)
	return b
}

/*
*
Default: 1
Example: page=1
Page number of the results to fetch.
*/
func (b *ListAPIKeysReqBuilder) Page(page string) *ListAPIKeysReqBuilder {
	b.apiReq.QueryParams.Set("page", page)
	return b
}

func (b *ListAPIKeysReqBuilder) Build() *ListAPIKeysReq {
	req := &ListAPIKeysReq{}
	req.apiReq = b.apiReq
	return req
}

type ListAPIKeysResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*APIKey `json:"data"`
	Meta *Meta     `json:"meta"`
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
	Expired *string `json:"expired,omitempty"`
	Note    *string `json:"note,omitempty"`
	Tag     *string `json:"tag,omitempty"`
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
	}
	return builder
}

/*
*
API Key to update
*/
func (b *UpdateAPIKeyReqBuilder) APIKey(apiKey string) *UpdateAPIKeyReqBuilder {
	b.apiReq.PathParams.Set("api_key", apiKey)
	return b
}

/*
*
Expiration time, format: rfc3339, for example: 2025-12-31T23:59:59+08:00
*/
func (b *UpdateAPIKeyReqBuilder) Expired(expired string) *UpdateAPIKeyReqBuilder {
	b.apiReq.Body.(*UpdateAPIKeyReqBody).Expired = &expired
	return b
}

/*
*
Note for the API key
*/
func (b *UpdateAPIKeyReqBuilder) Note(note string) *UpdateAPIKeyReqBuilder {
	b.apiReq.Body.(*UpdateAPIKeyReqBody).Note = &note
	return b
}

/*
*
Tag for the API key
*/
func (b *UpdateAPIKeyReqBuilder) Tag(tag string) *UpdateAPIKeyReqBuilder {
	b.apiReq.Body.(*UpdateAPIKeyReqBody).Tag = &tag
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
