package metrics

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/yao560909/emqx-admin-go-sdk/core"
)

const (
	ApiPathMetricsStats                     = "/api/v5/stats"
	ApiPathMetrics                          = "/api/v5/metrics"
	ApiPathMetricsMonitor                   = "/api/v5/monitor"
	ApiPathMetricsMonitorNodes_Node         = "/api/v5/monitor/nodes/{node}"
	ApiPathMetricsMonitor_current           = "/api/v5/monitor_current"
	ApiPathMetricsMonitor_currentNodes_Node = "/api/v5/monitor_current/nodes/{node}"
)

type MetricsService struct {
	config *core.Config
}

func NewService(config *core.Config) *MetricsService {
	return &MetricsService{config: config}
}

// EMQX stats
func (s *MetricsService) Stats(ctx context.Context, req *StatsReq, options ...core.RequestOptionFunc) (*StatsResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathMetricsStats
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[Stats] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &StatsResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[Stats] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// EMQX metrics
func (s *MetricsService) Metrics(ctx context.Context, req *MetricsReq, options ...core.RequestOptionFunc) (*MetricsResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathMetrics
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[Metrics] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &MetricsResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[Metrics] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// List monitor (statistics) data for the whole cluster.
func (s *MetricsService) ListMonitorForTheCluster(ctx context.Context, req *ListMonitorForTheClusterReq, options ...core.RequestOptionFunc) (*ListMonitorForTheClusterResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathMetricsMonitor
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListMonitorForTheCluster] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &ListMonitorForTheClusterResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListMonitorForTheCluster] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// List the monitor (statistics) data on the specified node.
func (s *MetricsService) ListMonitorForNode(ctx context.Context, req *ListMonitorForNodeReq, options ...core.RequestOptionFunc) (*ListMonitorForNodeResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathMetricsMonitorNodes_Node
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListMonitorForNode] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &ListMonitorForNodeResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListMonitorForNode] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// Current monitor (statistics) data, e.g. number of connections and connection rate in the whole cluster.
func (s *MetricsService) CurrentMonitorForTheCluster(ctx context.Context, req *CurrentMonitorForTheClusterReq, options ...core.RequestOptionFunc) (*CurrentMonitorForTheClusterResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathMetricsMonitor_current
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[CurrentMonitorForTheCluster] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &CurrentMonitorForTheClusterResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[CurrentMonitorForTheCluster] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// Node monitor (statistics) data, e.g. number of connections and connection rate on the specified node.
func (s *MetricsService) CurrentMonitorForNode(ctx context.Context, req *CurrentMonitorForNodeRep, options ...core.RequestOptionFunc) (*CurrentMonitorForNodeResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathMetricsMonitor_currentNodes_Node
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[CurrentMonitorForNode] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &CurrentMonitorForNodeResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[CurrentMonitorForNode] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}
