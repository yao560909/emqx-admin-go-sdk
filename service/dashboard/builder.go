package dashboard

import (
	"encoding/json"

	"github.com/yao560909/emqx-admin-go-sdk/core"
)

type ListUsersReq struct {
	apiReq *core.APIReq
}

// ListUsersReqBuilder represents the builder for ListUsersReq.
type ListUsersReqBuilder struct {
	apiReq *core.APIReq
}

// NewListUsersReqBuilder creates a new ListUsersReqBuilder.
func NewListUsersReqBuilder() *ListUsersReqBuilder {
	builder := &ListUsersReqBuilder{}
	builder.apiReq = &core.APIReq{
		SkipAuth:    true,
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

// Build builds the ListUsersReq.
func (b *ListUsersReqBuilder) Build() *ListUsersReq {
	req := &ListUsersReq{}
	req.apiReq = b.apiReq
	return req
}

// ListUsersResp represents the response for listing users.
type ListUsersResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*User `json:"-"`
}

func (resp *ListUsersResp) UnmarshalJSON(b []byte) error {
	var users []*User
	if err := json.Unmarshal(b, &users); err == nil {
		resp.Data = users
		return nil
	}
	type alias ListUsersResp
	return json.Unmarshal(b, (*alias)(resp))
}

type LoginReq struct {
	apiReq *core.APIReq
}

type LoginReqBody struct {
	Username *string `json:"username,omitempty"`
	Password *string `json:"password,omitempty"`
}

type LoginReqBuilder struct {
	apiReq *core.APIReq
}

func NewLoginReqBuilder() *LoginReqBuilder {
	builder := &LoginReqBuilder{}
	builder.apiReq = &core.APIReq{
		SkipAuth:    true,
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &LoginReqBody{},
	}
	return builder
}

func (b *LoginReqBuilder) Username(username string) *LoginReqBuilder {
	b.apiReq.Body.(*LoginReqBody).Username = &username
	return b
}

func (b *LoginReqBuilder) Password(password string) *LoginReqBuilder {
	b.apiReq.Body.(*LoginReqBody).Password = &password
	return b
}

func (b *LoginReqBuilder) Builder() *LoginReq {
	req := &LoginReq{}
	req.apiReq = b.apiReq
	return req
}

type LoginResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	LoginResponse
}

type CreateUserReq struct {
	apiReq *core.APIReq
}

type CreateUserReqBody struct {
	Username    *string `json:"username,omitempty"`
	Password    *string `json:"password,omitempty"`
	Description *string `json:"description,omitempty"`
}

type CreateUserReqBuilder struct {
	apiReq *core.APIReq
}

func NewCreateUserReqBuilder() *CreateUserReqBuilder {
	builder := &CreateUserReqBuilder{}
	builder.apiReq = &core.APIReq{
		SkipAuth:    true,
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &CreateUserReqBody{},
	}
	return builder
}

func (b *CreateUserReqBuilder) Username(username string) *CreateUserReqBuilder {
	b.apiReq.Body.(*CreateUserReqBody).Username = &username
	return b
}

func (b *CreateUserReqBuilder) Password(password string) *CreateUserReqBuilder {
	b.apiReq.Body.(*CreateUserReqBody).Password = &password
	return b
}

func (b *CreateUserReqBuilder) Description(description string) *CreateUserReqBuilder {
	b.apiReq.Body.(*CreateUserReqBody).Description = &description
	return b
}

func (b *CreateUserReqBuilder) Build() *CreateUserReq {
	req := &CreateUserReq{}
	req.apiReq = b.apiReq
	return req
}

type CreateUserResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	User
}

type UpdateUserReq struct {
	apiReq *core.APIReq
}

type UpdateUserReqBuilder struct {
	apiReq *core.APIReq
}

func NewUpdateUserReqBuilder() *UpdateUserReqBuilder {
	builder := &UpdateUserReqBuilder{}
	builder.apiReq = &core.APIReq{
		SkipAuth:    true,
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &UpdateUserReqBody{},
	}
	return builder
}

type UpdateUserReqBody struct {
	Description *string `json:"description,omitempty"`
}

func (b *UpdateUserReqBuilder) Build() *UpdateUserReq {
	req := &UpdateUserReq{}
	req.apiReq = b.apiReq
	return req
}

func (b *UpdateUserReqBuilder) Description(description string) *UpdateUserReqBuilder {
	b.apiReq.Body.(*UpdateUserReqBody).Description = &description
	return b
}

type UpdateUserResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	User
}

type DeleteUserReq struct {
	apiReq *core.APIReq
}

type DeleteUserReqBuilder struct {
	apiReq *core.APIReq
}

func NewDeleteUserReqBuilder() *DeleteUserReqBuilder {
	builder := &DeleteUserReqBuilder{}
	builder.apiReq = &core.APIReq{
		SkipAuth:    true,
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *DeleteUserReqBuilder) Username(username string) *DeleteUserReqBuilder {
	b.apiReq.PathParams["username"] = username
	return b
}

func (b *DeleteUserReqBuilder) Build() *DeleteUserReq {
	req := &DeleteUserReq{}
	req.apiReq = b.apiReq
	return req
}

type DeleteUserResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type ChangePasswordReq struct {
	apiReq *core.APIReq
}

type ChangePasswordReqBody struct {
	OldPwd *string `json:"old_pwd,omitempty"`
	NewPwd *string `json:"new_pwd,omitempty"`
}

type ChangePasswordReqBuilder struct {
	apiReq *core.APIReq
}

func NewChangePasswordReqBuilder() *ChangePasswordReqBuilder {
	builder := &ChangePasswordReqBuilder{}
	builder.apiReq = &core.APIReq{
		SkipAuth:    true,
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &ChangePasswordReqBody{},
	}
	return builder
}

func (b *ChangePasswordReqBuilder) Username(username string) *ChangePasswordReqBuilder {
	b.apiReq.PathParams["username"] = username
	return b
}

func (b *ChangePasswordReqBuilder) OldPwd(oldPwd string) *ChangePasswordReqBuilder {
	b.apiReq.Body.(*ChangePasswordReqBody).OldPwd = &oldPwd
	return b
}

func (b *ChangePasswordReqBuilder) NewPwd(newPwd string) *ChangePasswordReqBuilder {
	b.apiReq.Body.(*ChangePasswordReqBody).NewPwd = &newPwd
	return b
}

func (b *ChangePasswordReqBuilder) Build() *ChangePasswordReq {
	req := &ChangePasswordReq{}
	req.apiReq = b.apiReq
	return req
}

type ChangePasswordResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type LogoutReq struct {
	apiReq *core.APIReq
}

type LogoutReqBody struct {
	Username *string `json:"username,omitempty"`
}

type LogoutReqBuilder struct {
	apiReq *core.APIReq
}

func NewLogoutReqBuilder() *LogoutReqBuilder {
	builder := &LogoutReqBuilder{}
	builder.apiReq = &core.APIReq{
		SkipAuth:    true,
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &LogoutReqBody{},
	}
	return builder
}

func (b *LogoutReqBuilder) username(username string) *LogoutReqBuilder {
	b.apiReq.Body.(*LogoutReqBody).Username = &username
	return b
}

func (b *LogoutReqBuilder) Build() *LogoutReq {
	req := &LogoutReq{}
	req.apiReq = b.apiReq
	return req
}

type LogoutResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}
