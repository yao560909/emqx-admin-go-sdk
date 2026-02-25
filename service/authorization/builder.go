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
	Meta *Meta        `json:"meta"`
}

type AddRuleReq struct {
	apiReq *core.APIReq
}

type AddRuleReqBody struct {
	Username string  `json:"username"`
	Rules    []*Rule `json:"rules"`
}

type AddRuleReqBuilder struct {
	apiReq *core.APIReq
}

func NewAddRuleReqBuilder() *AddRuleReqBuilder {
	builder := &AddRuleReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *AddRuleReqBuilder) Username(username string) *AddRuleReqBuilder {
	b.apiReq.Body.(*AddRuleReqBody).Username = username
	return b
}

func (b *AddRuleReqBuilder) Rules(rules []*Rule) *AddRuleReqBuilder {
	b.apiReq.Body.(*AddRuleReqBody).Rules = rules
	return b
}

func (b *AddRuleReqBuilder) Build() *AddRuleReq {
	req := &AddRuleReq{}
	req.apiReq = b.apiReq
	return req
}

type AddRuleResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type ListRulesForClientsReq struct {
	apiReq *core.APIReq
}

type ListRulesForClientsReqBuilder struct {
	apiReq *core.APIReq
}

func NewListRulesForClientsReqBuilder() *ListRulesForClientsReqBuilder {
	builder := &ListRulesForClientsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *ListRulesForClientsReqBuilder) LikeClientId(likeClientId string) *ListRulesForClientsReqBuilder {
	b.apiReq.QueryParams.Set("like_clientid", likeClientId)
	return b
}

func (b *ListRulesForClientsReqBuilder) Page(page string) *ListRulesForClientsReqBuilder {
	b.apiReq.QueryParams.Set("page", page)
	return b
}

func (b *ListRulesForClientsReqBuilder) Limit(limit string) *ListRulesForClientsReqBuilder {
	b.apiReq.QueryParams.Set("limit", limit)
	return b
}

func (b *ListRulesForClientsReqBuilder) Build() *ListRulesForClientsReq {
	req := &ListRulesForClientsReq{}
	req.apiReq = b.apiReq
	return req
}

type ListRulesForClientsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*ClientRules `json:"data"`
	Meta *Meta          `json:"meta"`
}

type AddRuleForClientsReq struct {
	apiReq *core.APIReq
}

type AddRuleForClientsReqBody struct {
	ClientId string  `json:"clientid"`
	Rules    []*Rule `json:"rules"`
}

type AddRuleForClientsReqBuilder struct {
	apiReq *core.APIReq
}

func NewAddRuleForClientsReqBuilder() *AddRuleForClientsReqBuilder {
	builder := &AddRuleForClientsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *AddRuleForClientsReqBuilder) ClientId(clientId string) *AddRuleForClientsReqBuilder {
	b.apiReq.Body.(*AddRuleForClientsReqBody).ClientId = clientId
	return b
}

func (b *AddRuleForClientsReqBuilder) Rules(rules []*Rule) *AddRuleForClientsReqBuilder {
	b.apiReq.Body.(*AddRuleForClientsReqBody).Rules = rules
	return b
}

func (b *AddRuleForClientsReqBuilder) Build() *AddRuleForClientsReq {
	req := &AddRuleForClientsReq{}
	req.apiReq = b.apiReq
	return req
}

type AddRuleForClientsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type GetRuleForClientReq struct {
	apiReq *core.APIReq
}

type GetRuleForClientReqBuilder struct {
	apiReq *core.APIReq
}

func NewGetRuleForClientReqBuilder() *GetRuleForClientReqBuilder {
	builder := &GetRuleForClientReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *GetRuleForClientReqBuilder) ClientId(clientId string) *GetRuleForClientReqBuilder {
	b.apiReq.PathParams.Set("clientid", clientId)
	return b
}

func (b *GetRuleForClientReqBuilder) Build() *GetRuleForClientReq {
	req := &GetRuleForClientReq{}
	req.apiReq = b.apiReq
	return req
}

type GetRuleForClientResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	ClientRules
}
