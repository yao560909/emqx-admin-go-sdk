package monitor

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/yao560909/emqx-admin-go-sdk/core"
)

const (
	ApiPathPrometheusConfig    = "/api/v5/monitor/prometheus"
	ApiPathOpenTelemetryConfig = "/api/v5/monitor/opentelemetry"
)

type MonitorService struct {
	config *core.Config
}

func NewService(config *core.Config) *MonitorService {
	return &MonitorService{config: config}
}

// Get Prometheus config info
func (s *MonitorService) GetPrometheusConfig(ctx context.Context, req *GetPrometheusConfigReq, options ...core.RequestOptionFunc) (*GetPrometheusConfigResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathPrometheusConfig
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetPrometheusConfig] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &GetPrometheusConfigResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetPrometheusConfig] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// Update Prometheus config
func (s *MonitorService) UpdatePrometheusConfig(ctx context.Context, req *UpdatePrometheusConfigReq, options ...core.RequestOptionFunc) (*UpdatePrometheusConfigResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathPrometheusConfig
	apiReq.HttpMethod = "PUT"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[UpdatePrometheusConfig] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &UpdatePrometheusConfigResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[UpdatePrometheusConfig] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// Get opentelmetry configuration
func (s *MonitorService) GetOpenTelemetryConfig(ctx context.Context, req *GetOpenTelemetryConfigReq, options ...core.RequestOptionFunc) (*GetOpenTelemetryConfigResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathOpenTelemetryConfig
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetOpenTelemetryConfig] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &GetOpenTelemetryConfigResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetOpenTelemetryConfig] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// Update opentelmetry configuration
func (s *MonitorService) UpdateOpenTelemetryConfig(ctx context.Context, req *UpdateOpenTelemetryConfigReq, options ...core.RequestOptionFunc) (*UpdateOpenTelemetryConfigResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathOpenTelemetryConfig
	apiReq.HttpMethod = "PUT"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[UpdateOpenTelemetryConfig] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &UpdateOpenTelemetryConfigResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[UpdateOpenTelemetryConfig] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}
