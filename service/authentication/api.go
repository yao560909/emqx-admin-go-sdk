package authentication

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/yao560909/emqx-admin-go-sdk/core"
)

const (
	ApiPathAuthentication          = "/api/v5/authentication"
	ApiPathAuthentication_Id_Users = "/api/v5/authentication/{id}/users"
)

type AuthenticationService struct {
	config *core.Config
}

func NewService(config *core.Config) *AuthenticationService {
	return &AuthenticationService{config: config}
}

// List authenticators for global authentication.
func (s *AuthenticationService) ListAuthenticators(ctx context.Context, req *ListAuthenticatorsReq, options ...core.RequestOptionFunc) (*ListAuthenticatorsResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAuthentication
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListAuthenticators] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &ListAuthenticatorsResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListAuthenticators] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// List users in authenticator in global authentication chain.
func (s *AuthenticationService) ListUsersInAuthenticator(ctx context.Context, req *ListUsersInAuthenticatorReq, options ...core.RequestOptionFunc) (*ListUsersInAuthenticatorResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAuthentication_Id_Users
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListUsersInAuthenticator] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &ListUsersInAuthenticatorResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListUsersInAuthenticator] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// Create users for authenticator in global authentication chain.
func (s *AuthenticationService) CreateUsersForAuthenticator(ctx context.Context, req *CreateUsersForAuthenticatorReq, options ...core.RequestOptionFunc) (*CreateUsersForAuthenticatorResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAuthentication_Id_Users
	apiReq.HttpMethod = "POST"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[CreateUsersForAuthenticator] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &CreateUsersForAuthenticatorResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[CreateUsersForAuthenticator] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}
