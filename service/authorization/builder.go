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

type SetRuleForClientReq struct {
	apiReq *core.APIReq
}

type SetRuleForClientReqBody struct {
	Rules []*Rule `json:"rules"`
}

type SetRuleForClientReqBuilder struct {
	apiReq *core.APIReq
}

func NewSetRuleForClientReqBuilder() *SetRuleForClientReqBuilder {
	builder := &SetRuleForClientReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &SetRuleForClientReqBody{},
	}
	return builder
}

func (b *SetRuleForClientReqBuilder) ClientId(clientId string) *SetRuleForClientReqBuilder {
	b.apiReq.PathParams.Set("clientid", clientId)
	return b
}

func (b *SetRuleForClientReqBuilder) Rules(rules []*Rule) *SetRuleForClientReqBuilder {
	b.apiReq.Body.(*SetRuleForClientReqBody).Rules = rules
	return b
}

func (b *SetRuleForClientReqBuilder) Build() *SetRuleForClientReq {
	req := &SetRuleForClientReq{}
	req.apiReq = b.apiReq
	return req
}

type SetRuleForClientResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	ClientRules
}

type DeleteRuleForClientReq struct {
	apiReq *core.APIReq
}

type DeleteRuleForClientReqBuilder struct {
	apiReq *core.APIReq
}

func NewDeleteRuleForClientReqBuilder() *DeleteRuleForClientReqBuilder {
	builder := &DeleteRuleForClientReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *DeleteRuleForClientReqBuilder) ClientId(clientId string) *DeleteRuleForClientReqBuilder {
	b.apiReq.PathParams.Set("clientid", clientId)
	return b
}

func (b *DeleteRuleForClientReqBuilder) Build() *DeleteRuleForClientReq {
	req := &DeleteRuleForClientReq{}
	req.apiReq = b.apiReq
	return req
}

type DeleteRuleForClientResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type GetAuthorizationSourceStatusReq struct {
	apiReq *core.APIReq
}

type GetAuthorizationSourceStatusReqBuilder struct {
	apiReq *core.APIReq
}

func NewGetAuthorizationSourceStatusReqBuilder() *GetAuthorizationSourceStatusReqBuilder {
	builder := &GetAuthorizationSourceStatusReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *GetAuthorizationSourceStatusReqBuilder) Type(sourceType string) *GetAuthorizationSourceStatusReqBuilder {
	b.apiReq.PathParams.Set("type", sourceType)
	return b
}

func (b *GetAuthorizationSourceStatusReqBuilder) Build() *GetAuthorizationSourceStatusReq {
	req := &GetAuthorizationSourceStatusReq{}
	req.apiReq = b.apiReq
	return req
}

type GetAuthorizationSourceStatusResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	AuthorizationSourceStatus
}

type ListAuthorizationSourcesReq struct {
	apiReq *core.APIReq
}

type ListAuthorizationSourcesReqBuilder struct {
	apiReq *core.APIReq
}

func NewListAuthorizationSourcesReqBuilder() *ListAuthorizationSourcesReqBuilder {
	builder := &ListAuthorizationSourcesReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *ListAuthorizationSourcesReqBuilder) Build() *ListAuthorizationSourcesReq {
	req := &ListAuthorizationSourcesReq{}
	req.apiReq = b.apiReq
	return req
}

type ListAuthorizationSourcesResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*AuthorizationSource `json:"sources"`
}

type GetRuleForUserReq struct {
	apiReq *core.APIReq
}

type GetRuleForUserReqBuilder struct {
	apiReq *core.APIReq
}

func NewGetRuleForUserReqBuilder() *GetRuleForUserReqBuilder {
	builder := &GetRuleForUserReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *GetRuleForUserReqBuilder) Username(username string) *GetRuleForUserReqBuilder {
	b.apiReq.PathParams.Set("username", username)
	return b
}

func (b *GetRuleForUserReqBuilder) Build() *GetRuleForUserReq {
	req := &GetRuleForUserReq{}
	req.apiReq = b.apiReq
	return req
}

type GetRuleForUserResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	UserRules
}

type SetRuleForUserReq struct {
	apiReq *core.APIReq
}

type SetRuleForUserReqBody struct {
	Username string  `json:"username"`
	Rules    []*Rule `json:"rules"`
}

type SetRuleForUserReqBuilder struct {
	apiReq *core.APIReq
}

func NewSetRuleForUserReqBuilder() *SetRuleForUserReqBuilder {
	builder := &SetRuleForUserReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *SetRuleForUserReqBuilder) Username(username string) *SetRuleForUserReqBuilder {
	b.apiReq.PathParams.Set("username", username)
	b.apiReq.Body.(*SetRuleForUserReqBody).Username = username
	return b
}

func (b *SetRuleForUserReqBuilder) Rules(rules []*Rule) *SetRuleForUserReqBuilder {
	b.apiReq.Body.(*SetRuleForUserReqBody).Rules = rules
	return b
}

func (b *SetRuleForUserReqBuilder) Build() *SetRuleForUserReq {
	req := &SetRuleForUserReq{}
	req.apiReq = b.apiReq
	return req
}

type SetRuleForUserResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type DeleteRuleForUserReq struct {
	apiReq *core.APIReq
}

type DeleteRuleForUserReqBuilder struct {
	apiReq *core.APIReq
}

func NewDeleteRuleForUserReqBuilder() *DeleteRuleForUserReqBuilder {
	builder := &DeleteRuleForUserReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *DeleteRuleForUserReqBuilder) Username(username string) *DeleteRuleForUserReqBuilder {
	b.apiReq.PathParams.Set("username", username)
	return b
}

func (b *DeleteRuleForUserReqBuilder) Build() *DeleteRuleForUserReq {
	req := &DeleteRuleForUserReq{}
	req.apiReq = b.apiReq
	return req
}

type DeleteRuleForUserResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type DeleteAllRulesReq struct {
	apiReq *core.APIReq
}

type DeleteAllRulesReqBuilder struct {
	apiReq *core.APIReq
}

func NewDeleteAllRulesReqBuilder() *DeleteAllRulesReqBuilder {
	builder := &DeleteAllRulesReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *DeleteAllRulesReqBuilder) Build() *DeleteAllRulesReq {
	req := &DeleteAllRulesReq{}
	req.apiReq = b.apiReq
	return req
}

type DeleteAllRulesResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type MoveAuthorizationSourceReq struct {
	apiReq *core.APIReq
}

type MoveAuthorizationSourceReqBody struct {
	Position string `json:"position"`
}

type MoveAuthorizationSourceReqBuilder struct {
	apiReq *core.APIReq
}

func NewMoveAuthorizationSourceReqBuilder() *MoveAuthorizationSourceReqBuilder {
	builder := &MoveAuthorizationSourceReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

/*
*
Enum: "file" "built_in_database" "http" "redis" "mysql" "postgresql" "mongodb" "ldap"
Authorization type
*/
func (b *MoveAuthorizationSourceReqBuilder) Type(sourceType string) *MoveAuthorizationSourceReqBuilder {
	b.apiReq.PathParams.Set("type", sourceType)
	return b
}

func (b *MoveAuthorizationSourceReqBuilder) Position(position string) *MoveAuthorizationSourceReqBuilder {
	b.apiReq.Body.(*MoveAuthorizationSourceReqBody).Position = position
	return b
}

func (b *MoveAuthorizationSourceReqBuilder) Build() *MoveAuthorizationSourceReq {
	req := &MoveAuthorizationSourceReq{}
	req.apiReq = b.apiReq
	return req
}

type MoveAuthorizationSourceResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type GetAuthorizationSourceReq struct {
	apiReq *core.APIReq
}

type GetAuthorizationSourceReqBuilder struct {
	apiReq *core.APIReq
}

func NewGetAuthorizationSourceReqBuilder() *GetAuthorizationSourceReqBuilder {
	builder := &GetAuthorizationSourceReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

/*
*
Enum: "file" "built_in_database" "http" "redis" "mysql" "postgresql" "mongodb" "ldap"
Authorization type
*/
func (b *GetAuthorizationSourceReqBuilder) Type(sourceType string) *GetAuthorizationSourceReqBuilder {
	b.apiReq.PathParams.Set("type", sourceType)
	return b
}

func (b *GetAuthorizationSourceReqBuilder) Build() *GetAuthorizationSourceReq {
	req := &GetAuthorizationSourceReq{}
	req.apiReq = b.apiReq
	return req
}

type GetAuthorizationSourceResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	sources []*AuthorizationSource `json:"sources"`
}

type DeleteAuthorizationSourceReq struct {
	apiReq *core.APIReq
}

type DeleteAuthorizationSourceResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type DeleteAuthorizationSourceReqBuilder struct {
	apiReq *core.APIReq
}

func NewDeleteAuthorizationSourceReqBuilder() *DeleteAuthorizationSourceReqBuilder {
	builder := &DeleteAuthorizationSourceReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

/*
*
Enum: "file" "built_in_database" "http" "redis" "mysql" "postgresql" "mongodb" "ldap"
Authorization type
*/
func (b *DeleteAuthorizationSourceReqBuilder) Type(sourceType string) *DeleteAuthorizationSourceReqBuilder {
	b.apiReq.PathParams.Set("type", sourceType)
	return b
}

func (b *DeleteAuthorizationSourceReqBuilder) Build() *DeleteAuthorizationSourceReq {
	req := &DeleteAuthorizationSourceReq{}
	req.apiReq = b.apiReq
	return req
}

type CleanAuthorizationCacheReq struct {
	apiReq *core.APIReq
}

type CleanAuthorizationCacheResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type CleanAuthorizationCacheReqBuilder struct {
	apiReq *core.APIReq
}

func NewCleanAuthorizationCacheReqBuilder() *CleanAuthorizationCacheReqBuilder {
	builder := &CleanAuthorizationCacheReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *CleanAuthorizationCacheReqBuilder) Build() *CleanAuthorizationCacheReq {
	req := &CleanAuthorizationCacheReq{}
	req.apiReq = b.apiReq
	return req
}
