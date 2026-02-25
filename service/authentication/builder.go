package authentication

import (
	"encoding/json"

	"github.com/yao560909/emqx-admin-go-sdk/core"
)

// --------------------------------
type ListAuthenticatorsReq struct {
	apiReq *core.APIReq
}

type ListAuthenticatorsReqBuilder struct {
	apiReq *core.APIReq
}

func NewListAuthenticatorsReqBuilder() *ListAuthenticatorsReqBuilder {
	builder := &ListAuthenticatorsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *ListAuthenticatorsReqBuilder) Build() *ListAuthenticatorsReq {
	req := &ListAuthenticatorsReq{}
	req.apiReq = b.apiReq
	return req
}

type ListAuthenticatorsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*Authentication `json:"-"`
}

func (resp *ListAuthenticatorsResp) UnmarshalJSON(b []byte) error {
	var authentication []*Authentication
	if err := json.Unmarshal(b, &authentication); err == nil {
		resp.Data = authentication
		return nil
	}
	type alias ListAuthenticatorsResp
	return json.Unmarshal(b, (*alias)(resp))
}

// --------------------------------------
type ListUsersInAuthenticatorReq struct {
	apiReq *core.APIReq
}

type ListUsersInAuthenticatorReqBuilder struct {
	apiReq *core.APIReq
}

func NewListUsersInAuthenticatorReqBuilder() *ListUsersInAuthenticatorReqBuilder {
	builder := &ListUsersInAuthenticatorReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

/*
*
Authenticator ID.
*/
func (b *ListUsersInAuthenticatorReqBuilder) Id(id string) *ListUsersInAuthenticatorReqBuilder {
	b.apiReq.PathParams.Set("id", id)
	return b
}

/*
*
Default: 1
Example: page=1
Page number of the results to fetch.
*/
func (b *ListUsersInAuthenticatorReqBuilder) Page(page string) *ListUsersInAuthenticatorReqBuilder {
	b.apiReq.QueryParams.Set("page", page)
	return b
}

/*
*
Default: 100
Example: limit=50
Results per page(max 1000)
*/
func (b *ListUsersInAuthenticatorReqBuilder) Limit(limit string) *ListUsersInAuthenticatorReqBuilder {
	b.apiReq.QueryParams.Set("limit", limit)
	return b
}

/*
*
Fuzzy search user_id (username or clientid).
*/
func (b *ListUsersInAuthenticatorReqBuilder) LikeUserId(likeUserId string) *ListUsersInAuthenticatorReqBuilder {
	b.apiReq.QueryParams.Set("like_user_id", likeUserId)
	return b
}

/*
*
Is superuser
*/
func (b *ListUsersInAuthenticatorReqBuilder) IsSuperuser(isSuperuser string) *ListUsersInAuthenticatorReqBuilder {
	b.apiReq.QueryParams.Set("is_superuser", isSuperuser)
	return b
}

func (b *ListUsersInAuthenticatorReqBuilder) Build() *ListUsersInAuthenticatorReq {
	req := &ListUsersInAuthenticatorReq{}
	req.apiReq = b.apiReq
	return req
}

type ListUsersInAuthenticatorResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*User `json:"data"`
	Meta *Meta   `json:"meta"`
}

// ------------------------------------------
type CreateUsersForAuthenticatorReq struct {
	apiReq *core.APIReq
}

type CreateUsersForAuthenticatorReqBuilder struct {
	apiReq *core.APIReq
}

func NewCreateUsersForAuthenticatorReqBuilder() *CreateUsersForAuthenticatorReqBuilder {
	builder := &CreateUsersForAuthenticatorReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

/*
*
Authenticator ID.
*/
func (b *CreateUsersForAuthenticatorReqBuilder) Id(id string) *CreateUsersForAuthenticatorReqBuilder {
	b.apiReq.PathParams.Set("id", id)
	return b
}

type UserBody struct {
	UserId      *string `json:"user_id,omitempty"`
	IsSuperuser *string `json:"is_superuser,omitempty"`
	Password    *string `json:"password,omitempty"`
}

func (b *CreateUsersForAuthenticatorReqBuilder) UserId(userId string) *CreateUsersForAuthenticatorReqBuilder {
	b.apiReq.Body.(*UserBody).UserId = &userId
	return b
}

func (b *CreateUsersForAuthenticatorReqBuilder) Password(password string) *CreateUsersForAuthenticatorReqBuilder {
	b.apiReq.Body.(*UserBody).Password = &password
	return b
}

/*
*
Default: false
*/
func (b *CreateUsersForAuthenticatorReqBuilder) IsSuperuser(isSuperuser string) *CreateUsersForAuthenticatorReqBuilder {
	b.apiReq.Body.(*UserBody).IsSuperuser = &isSuperuser
	return b
}

func (b *CreateUsersForAuthenticatorReqBuilder) Build() *CreateUsersForAuthenticatorReq {
	req := &CreateUsersForAuthenticatorReq{}
	req.apiReq = b.apiReq
	return req
}

type CreateUsersForAuthenticatorResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	User
}

// ------------------------------------------
type MoveAuthenticatorPositionReq struct {
	apiReq *core.APIReq
}

type MoveAuthenticatorPositionReqBuilder struct {
	apiReq *core.APIReq
}

func NewMoveAuthenticatorPositionReqBuilder() *MoveAuthenticatorPositionReqBuilder {
	builder := &MoveAuthenticatorPositionReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

/*
*
Authenticator ID.
*/
func (b *MoveAuthenticatorPositionReqBuilder) Id(id string) *MoveAuthenticatorPositionReqBuilder {
	b.apiReq.PathParams.Set("id", id)
	return b
}

/*
*
Position of the authenticator in the chain.
*/
func (b *MoveAuthenticatorPositionReqBuilder) Position(position string) *MoveAuthenticatorPositionReqBuilder {
	b.apiReq.PathParams.Set("position", position)
	return b
}

func (b *MoveAuthenticatorPositionReqBuilder) Build() *MoveAuthenticatorPositionReq {
	req := &MoveAuthenticatorPositionReq{}
	req.apiReq = b.apiReq
	return req
}

type MoveAuthenticatorPositionResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

// ------------------------------------------
type GetUserFromAuthenticatorReq struct {
	apiReq *core.APIReq
}

type GetUserFromAuthenticatorReqBuilder struct {
	apiReq *core.APIReq
}

func NewGetUserFromAuthenticatorReqBuilder() *GetUserFromAuthenticatorReqBuilder {
	builder := &GetUserFromAuthenticatorReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

/*
*
Authenticator ID.
*/
func (b *GetUserFromAuthenticatorReqBuilder) Id(id string) *GetUserFromAuthenticatorReqBuilder {
	b.apiReq.PathParams.Set("id", id)
	return b
}

/*
*
User ID.
*/
func (b *GetUserFromAuthenticatorReqBuilder) UserId(userId string) *GetUserFromAuthenticatorReqBuilder {
	b.apiReq.PathParams.Set("user_id", userId)
	return b
}

func (b *GetUserFromAuthenticatorReqBuilder) Build() *GetUserFromAuthenticatorReq {
	req := &GetUserFromAuthenticatorReq{}
	req.apiReq = b.apiReq
	return req
}

type GetUserFromAuthenticatorResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	User
}

// ------------------------------------------
type UpdateUserInAuthenticatorReq struct {
	apiReq *core.APIReq
}

type UpdateUserInAuthenticatorReqBuilder struct {
	apiReq *core.APIReq
}

func NewUpdateUserInAuthenticatorReqBuilder() *UpdateUserInAuthenticatorReqBuilder {
	builder := &UpdateUserInAuthenticatorReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &UserBody{},
	}
	return builder
}

/*
*
Authenticator ID.
*/
func (b *UpdateUserInAuthenticatorReqBuilder) Id(id string) *UpdateUserInAuthenticatorReqBuilder {
	b.apiReq.PathParams.Set("id", id)
	return b
}

/*
*
User ID.
*/
func (b *UpdateUserInAuthenticatorReqBuilder) UserId(userId string) *UpdateUserInAuthenticatorReqBuilder {
	b.apiReq.PathParams.Set("user_id", userId)
	return b
}

func (b *UpdateUserInAuthenticatorReqBuilder) Password(password string) *UpdateUserInAuthenticatorReqBuilder {
	b.apiReq.Body.(*UserBody).Password = &password
	return b
}

/*
*
Default: false
*/
func (b *UpdateUserInAuthenticatorReqBuilder) IsSuperuser(isSuperuser string) *UpdateUserInAuthenticatorReqBuilder {
	b.apiReq.Body.(*UserBody).IsSuperuser = &isSuperuser
	return b
}

func (b *UpdateUserInAuthenticatorReqBuilder) Build() *UpdateUserInAuthenticatorReq {
	req := &UpdateUserInAuthenticatorReq{}
	req.apiReq = b.apiReq
	return req
}

type UpdateUserInAuthenticatorResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	User
}

// ------------------------------------------
type DeleteUserFromAuthenticatorReq struct {
	apiReq *core.APIReq
}

type DeleteUserFromAuthenticatorReqBuilder struct {
	apiReq *core.APIReq
}

func NewDeleteUserFromAuthenticatorReqBuilder() *DeleteUserFromAuthenticatorReqBuilder {
	builder := &DeleteUserFromAuthenticatorReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

/*
*
Authenticator ID.
*/
func (b *DeleteUserFromAuthenticatorReqBuilder) Id(id string) *DeleteUserFromAuthenticatorReqBuilder {
	b.apiReq.PathParams.Set("id", id)
	return b
}

/*
*
User ID.
*/
func (b *DeleteUserFromAuthenticatorReqBuilder) UserId(userId string) *DeleteUserFromAuthenticatorReqBuilder {
	b.apiReq.PathParams.Set("user_id", userId)
	return b
}

func (b *DeleteUserFromAuthenticatorReqBuilder) Build() *DeleteUserFromAuthenticatorReq {
	req := &DeleteUserFromAuthenticatorReq{}
	req.apiReq = b.apiReq
	return req
}

type DeleteUserFromAuthenticatorResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}
