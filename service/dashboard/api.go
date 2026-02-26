package dashboard

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/yao560909/emqx-admin-go-sdk/core"
)

const (
	ApiPathUsers          = "/api/v5/users"
	ApiPathUsers_Username = "/api/v5/users/{username}"
	ApiPathLogin          = "/api/v5/login"
)

type DashboardService struct {
	config *core.Config
}

func NewService(config *core.Config) *DashboardService {
	return &DashboardService{config: config}
}

// Get Dashboard Auth Token.
func (s *DashboardService) Login(ctx context.Context, req *LoginReq, options ...core.RequestOptionFunc) (*LoginResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathLogin
	apiReq.HttpMethod = "POST"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[Login] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &LoginResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[Login] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil

}

// 需要先调用login获取token
// Dashboard list users
func (s *DashboardService) ListUsers(ctx context.Context, req *ListUsersReq, options ...core.RequestOptionFunc) (*ListUsersResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathUsers
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListUsers] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &ListUsersResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListUsers] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// 需要先调用login获取token
// Create dashboard user
func (s *DashboardService) CreateUser(ctx context.Context, req *CreateUserReq, options ...core.RequestOptionFunc) (*CreateUserResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathUsers
	apiReq.HttpMethod = "POST"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[CreateUser] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &CreateUserResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[CreateUser] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// 需要先调用login获取token
// Update dashboard user description
func (s *DashboardService) UpdateUser(ctx context.Context, req *UpdateUserReq, options ...core.RequestOptionFunc) (*UpdateUserResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathUsers_Username
	apiReq.HttpMethod = "PUT"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[UpdateUser] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &UpdateUserResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[UpdateUser] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// 需要先调用login获取token
// Delete dashboard user
func (s *DashboardService) DeleteUser(ctx context.Context, req *DeleteUserReq, options ...core.RequestOptionFunc) (*DeleteUserResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathUsers_Username
	apiReq.HttpMethod = "DELETE"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[DeleteUser] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &DeleteUserResp{APIResp: apiResp}
	if len(apiResp.RawBody) > 0 {
		err = json.Unmarshal(apiResp.RawBody, resp)
		if err != nil {
			s.config.Logger.Error(ctx, fmt.Sprintf("[DeleteUser] fail to unmarshal response body, error: %v", err.Error()))
			return nil, err
		}
	}
	return resp, nil
}
