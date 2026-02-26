package configs

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/yao560909/emqx-admin-go-sdk/core"
)

const (
	ApiPathSysTopicsConfig  = "/api/v5/configs/sys_topics"
	ApiPathSysmonConfig     = "/api/v5/configs/sysmon"
	ApiPathGlobalZoneConfig = "/api/v5/configs/global_zone"
	ApiPathAlarmConfig      = "/api/v5/configs/alarm"
	ApiPathDashboardConfig  = "/api/v5/configs/dashboard"
	ApiPathLogConfig        = "/api/v5/configs/log"
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

// Get the sub-configurations under sysmon
func (s *ConfigsService) GetSysmonConfig(ctx context.Context, req *GetSysmonConfigReq, options ...core.RequestOptionFunc) (*GetSysmonConfigResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathSysmonConfig
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetSysmonConfig] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &GetSysmonConfigResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetSysmonConfig] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

func (s *ConfigsService) UpdateSysmonConfig(ctx context.Context, req *UpdateSysmonConfigReq, options ...core.RequestOptionFunc) (*UpdateSysmonConfigResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathSysmonConfig
	apiReq.HttpMethod = "PUT"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[UpdateSysmonConfig] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &UpdateSysmonConfigResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[UpdateSysmonConfig] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// Get the sub-configurations under global_zone
func (s *ConfigsService) GetGlobalZoneConfig(ctx context.Context, req *GetGlobalZoneConfigReq, options ...core.RequestOptionFunc) (*GetGlobalZoneConfigResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathGlobalZoneConfig
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetGlobalZoneConfig] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &GetGlobalZoneConfigResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetGlobalZoneConfig] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

func (s *ConfigsService) UpdateGlobalZoneConfig(ctx context.Context, req *UpdateGlobalZoneConfigReq, options ...core.RequestOptionFunc) (*UpdateGlobalZoneConfigResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathGlobalZoneConfig
	apiReq.HttpMethod = "PUT"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[UpdateGlobalZoneConfig] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &UpdateGlobalZoneConfigResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[UpdateGlobalZoneConfig] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// Get the sub-configurations under alarm
func (s *ConfigsService) GetAlarmConfig(ctx context.Context, req *GetAlarmConfigReq, options ...core.RequestOptionFunc) (*GetAlarmConfigResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAlarmConfig
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetAlarmConfig] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &GetAlarmConfigResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetAlarmConfig] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

func (s *ConfigsService) UpdateAlarmConfig(ctx context.Context, req *UpdateAlarmConfigReq, options ...core.RequestOptionFunc) (*UpdateAlarmConfigResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAlarmConfig
	apiReq.HttpMethod = "PUT"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[UpdateAlarmConfig] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &UpdateAlarmConfigResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[UpdateAlarmConfig] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// Get the sub-configurations under dashboard
func (s *ConfigsService) GetDashboardConfig(ctx context.Context, req *GetDashboardConfigReq, options ...core.RequestOptionFunc) (*GetDashboardConfigResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathDashboardConfig
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetDashboardConfig] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &GetDashboardConfigResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetDashboardConfig] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// Get the sub-configurations under log
func (s *ConfigsService) GetLogConfig(ctx context.Context, req *GetLogConfigReq, options ...core.RequestOptionFunc) (*GetLogConfigResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathLogConfig
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetLogConfig] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &GetLogConfigResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetLogConfig] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}
