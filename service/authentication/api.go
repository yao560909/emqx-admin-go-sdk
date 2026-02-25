package authentication

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/yao560909/emqx-admin-go-sdk/core"
)

const (
	ApiPathAuthentication                 = "/api/v5/authentication"
	ApiPathAuthentication_Id_Users        = "/api/v5/authentication/{id}/users"
	ApiPathAuthentication_Id_Position     = "/api/v5/authentication/{id}/position/{position}"
	ApiPathAuthentication_Id_Users_UserId = "/api/v5/authentication/{id}/users/{user_id}"
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

// Move authenticator position in global authentication chain.
func (s *AuthenticationService) MoveAuthenticatorPosition(ctx context.Context, req *MoveAuthenticatorPositionReq, options ...core.RequestOptionFunc) (*MoveAuthenticatorPositionResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAuthentication_Id_Position
	apiReq.HttpMethod = "PUT"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[MoveAuthenticatorPosition] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &MoveAuthenticatorPositionResp{APIResp: apiResp}
	if len(apiResp.RawBody) > 0 {
		err = json.Unmarshal(apiResp.RawBody, resp)
		if err != nil {
			s.config.Logger.Error(ctx, fmt.Sprintf("[MoveAuthenticatorPosition] fail to unmarshal response body, error: %v", err.Error()))
			return nil, err
		}
	}
	return resp, nil
}

// Get user from authenticator in global authentication chain.
func (s *AuthenticationService) GetUserFromAuthenticator(ctx context.Context, req *GetUserFromAuthenticatorReq, options ...core.RequestOptionFunc) (*GetUserFromAuthenticatorResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAuthentication_Id_Users_UserId
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetUserFromAuthenticator] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &GetUserFromAuthenticatorResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetUserFromAuthenticator] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// Update user in authenticator in global authentication chain.
func (s *AuthenticationService) UpdateUserInAuthenticator(ctx context.Context, req *UpdateUserInAuthenticatorReq, options ...core.RequestOptionFunc) (*UpdateUserInAuthenticatorResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAuthentication_Id_Users_UserId
	apiReq.HttpMethod = "PUT"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[UpdateUserInAuthenticator] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &UpdateUserInAuthenticatorResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[UpdateUserInAuthenticator] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}
