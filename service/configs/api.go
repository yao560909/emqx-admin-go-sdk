package configs

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/yao560909/emqx-admin-go-sdk/core"
)

const (
	ApiPathSysTopicsConfig = "/api/v5/configs/sys_topics"
)

type ConfigsService struct {
	config *core.Config
}

func NewService(config *core.Config) *ConfigsService {
	return &ConfigsService{config: config}
}

// Get the sub-configurations under sys_topics
func (s *ConfigsService) GetSysTopicsConfig(ctx context.Context, req *GetSysTopicsConfigReq, options ...core.RequestOptionFunc) (*GetSysTopicsConfigResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathSysTopicsConfig
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetSysTopicsConfig] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &GetSysTopicsConfigResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetSysTopicsConfig] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

func (s *ConfigsService) UpdateSysTopicsConfig(ctx context.Context, req *UpdateSysTopicsConfigReq, options ...core.RequestOptionFunc) (*UpdateSysTopicsConfigResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathSysTopicsConfig
	apiReq.HttpMethod = "PUT"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[UpdateSysTopicsConfig] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &UpdateSysTopicsConfigResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[UpdateSysTopicsConfig] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}
