package apikeys

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/yao560909/emqx-admin-go-sdk/core"
)

const (
	ApiPathAPIKey        = "/api/v5/api_key"
	ApiPathAPIKey_APIKey = "/api/v5/api_key/{api_key}"
	ApiPathAPIKey_Name   = "/api/v5/api_key/{name}"
)

type APIKeysService struct {
	config *core.Config
}

func NewService(config *core.Config) *APIKeysService {
	return &APIKeysService{config: config}
}

// Create an API key.
func (s *APIKeysService) CreateAPIKey(ctx context.Context, req *CreateAPIKeyReq, options ...core.RequestOptionFunc) (*CreateAPIKeyResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAPIKey
	apiReq.HttpMethod = "POST"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[CreateAPIKey] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &CreateAPIKeyResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[CreateAPIKey] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// Delete an API key.
func (s *APIKeysService) DeleteAPIKey(ctx context.Context, req *DeleteAPIKeyReq, options ...core.RequestOptionFunc) (*DeleteAPIKeyResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAPIKey_APIKey
	apiReq.HttpMethod = "DELETE"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[DeleteAPIKey] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &DeleteAPIKeyResp{APIResp: apiResp}
	if len(apiResp.RawBody) > 0 {
		err = json.Unmarshal(apiResp.RawBody, resp)
		if err != nil {
			s.config.Logger.Error(ctx, fmt.Sprintf("[DeleteAPIKey] fail to unmarshal response body, error: %v", err.Error()))
			return nil, err
		}
	}
	return resp, nil
}

// List all API keys.
func (s *APIKeysService) ListAPIKeys(ctx context.Context, req *ListAPIKeysReq, options ...core.RequestOptionFunc) (*ListAPIKeysResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAPIKey
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListAPIKeys] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &ListAPIKeysResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListAPIKeys] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// Get an API key.
func (s *APIKeysService) GetAPIKey(ctx context.Context, req *GetAPIKeyReq, options ...core.RequestOptionFunc) (*GetAPIKeyResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAPIKey_Name
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetAPIKey] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &GetAPIKeyResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetAPIKey] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// Update an API key.
func (s *APIKeysService) UpdateAPIKey(ctx context.Context, req *UpdateAPIKeyReq, options ...core.RequestOptionFunc) (*UpdateAPIKeyResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAPIKey_APIKey
	apiReq.HttpMethod = "PUT"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[UpdateAPIKey] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &UpdateAPIKeyResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[UpdateAPIKey] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}
