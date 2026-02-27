package banned

import "github.com/yao560909/emqx-admin-go-sdk/core"

// --------------------
type BannedReq struct {
	apiReq *core.APIReq
}

type BannedReqBody struct {
	As     *string `json:"as,omitempty"`
	Who    *string `json:"who,omitempty"`
	By     *string `json:"by,omitempty"`
	Reason *string `json:"reason,omitempty"`
	At     *string `json:"at,omitempty"`
	Until  *string `json:"until,omitempty"`
}

type BannedReqBuilder struct {
	apiReq *core.APIReq
}

func NewBannedReqBuilder() *BannedReqBuilder {
	builder := &BannedReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &BannedReqBody{},
	}
	return builder
}

/*
*
Enum: "clientid" "username" "peerhost"
Ban method, which can be client ID, username or IP address.
*/
func (b *BannedReqBuilder) As(as string) *BannedReqBuilder {
	b.apiReq.Body.(*BannedReqBody).As = &as
	return b
}

/*
*
Ban object, specific client ID, username or IP address.
*/
func (b *BannedReqBuilder) Who(who string) *BannedReqBuilder {
	b.apiReq.Body.(*BannedReqBody).Who = &who
	return b
}

/*
*
Ban reason, record the reason why the current object was banned.
*/
func (b *BannedReqBuilder) Reason(reason string) *BannedReqBuilder {
	b.apiReq.Body.(*BannedReqBody).Reason = &reason
	return b
}

/*
*
The end time of the ban, the format is rfc3339, the default is the time when the operation was initiated + 1 year.
*/
func (b *BannedReqBuilder) Until(until string) *BannedReqBuilder {
	b.apiReq.Body.(*BannedReqBody).Until = &until
	return b
}

/*
*
The start time of the ban, the format is rfc3339, the default is the time when the operation was initiated.
*/
func (b *BannedReqBuilder) At(at string) *BannedReqBuilder {
	b.apiReq.Body.(*BannedReqBody).At = &at
	return b
}

/*
*
Initiator of the ban.
*/
func (b *BannedReqBuilder) By(by string) *BannedReqBuilder {
	b.apiReq.Body.(*BannedReqBody).By = &by
	return b
}

func (b BannedReqBuilder) Build() *BannedReq {
	req := &BannedReq{}
	req.apiReq = b.apiReq
	return req
}

type BannedResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Banned
}

// ---------------------------
type RemoveBannedReq struct {
	apiReq *core.APIReq
}

type RemoveBannedReqBuilder struct {
	apiReq *core.APIReq
}

func NewRemoveBannedReqBuilder() *RemoveBannedReqBuilder {
	builder := &RemoveBannedReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

/*
*
Enum: "clientid" "username" "peerhost"
Example: username
Ban method, which can be client ID, username or IP address.
*/
func (b *RemoveBannedReqBuilder) As(as string) *RemoveBannedReqBuilder {
	b.apiReq.PathParams.Set("as", as)
	return b
}

/*
*
Example: Badass
Ban object, specific client ID, username or IP address.
*/
func (b *RemoveBannedReqBuilder) Who(who string) *RemoveBannedReqBuilder {
	b.apiReq.PathParams.Set("who", who)
	return b
}

func (b *RemoveBannedReqBuilder) Build() *RemoveBannedReq {
	req := &RemoveBannedReq{}
	req.apiReq = b.apiReq
	return req
}

type RemoveBannedResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

// ----------------------------
type ClearAllBannedReq struct {
	apiReq *core.APIReq
}

type ClearAllBannedReqBuilder struct {
	apiReq *core.APIReq
}

func NewClearAllBannedReqBuilder() *ClearAllBannedReqBuilder {
	builder := &ClearAllBannedReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *ClearAllBannedReqBuilder) Build() *ClearAllBannedReq {
	req := &ClearAllBannedReq{}
	req.apiReq = b.apiReq
	return req
}

type ClearAllBannedResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

// ---------------------------------
type ListCurrentlyBannedReq struct {
	apiReq *core.APIReq
}

type ListCurrentlyBannedReqBuilder struct {
	apiReq *core.APIReq
}

func NewListCurrentlyBannedReqBuilder() *ListCurrentlyBannedReqBuilder {
	builder := &ListCurrentlyBannedReqBuilder{}
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
func (b *ListCurrentlyBannedReqBuilder) Limit(limit string) *ListCurrentlyBannedReqBuilder {
	b.apiReq.QueryParams.Set("limit", limit)
	return b
}

/*
*
Default: 1
Example: page=1
Page number of the results to fetch.
*/
func (b *ListCurrentlyBannedReqBuilder) Page(page string) *ListCurrentlyBannedReqBuilder {
	b.apiReq.QueryParams.Set("page", page)
	return b
}

func (b *ListCurrentlyBannedReqBuilder) Build() *ListCurrentlyBannedReq {
	req := &ListCurrentlyBannedReq{}
	req.apiReq = b.apiReq
	return req
}

type ListCurrentlyBannedResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*Banned `json:"data"`
	Meta *Meta     `json:"meta"`
}
