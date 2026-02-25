package authorization

import "github.com/yao560909/emqx-admin-go-sdk/core"

type ListRulesForUsersReq struct {
	apiReq *core.APIReq
}

type ListRulesForUsersReqBuilder struct {
	apiReq *core.APIReq
}

func NewListRulesForUsersReqBuilder() *ListRulesForUsersReqBuilder {
	builder := &ListRulesForUsersReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *ListRulesForUsersReqBuilder) LikeUsername(likeUsername string) *ListRulesForUsersReqBuilder {
	b.apiReq.QueryParams.Set("like_username", likeUsername)
	return b
}

func (b *ListRulesForUsersReqBuilder) Page(page string) *ListRulesForUsersReqBuilder {
	b.apiReq.QueryParams.Set("page", page)
	return b
}

func (b *ListRulesForUsersReqBuilder) Limit(limit string) *ListRulesForUsersReqBuilder {
	b.apiReq.QueryParams.Set("limit", limit)
	return b
}

func (b *ListRulesForUsersReqBuilder) Build() *ListRulesForUsersReq {
	req := &ListRulesForUsersReq{}
	req.apiReq = b.apiReq
	return req
}

type ListRulesForUsersResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*UserRules `json:"data"`
	Meta *Meta   `json:"meta"`
}
