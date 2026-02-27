package nodes

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/yao560909/emqx-admin-go-sdk/core"
)

const (
	ApiPathNodes              = "/api/v5/nodes"
	ApiPathNodes_Node         = "/api/v5/nodes/{node}"
	ApiPathNodes_Node_Metrics = "/api/v5/nodes/{node}/metrics"
	ApiPathNodes_Node_Stats   = "/api/v5/nodes/{node}/stats"
)

type NodesService struct {
	config *core.Config
}

func NewService(config *core.Config) *NodesService {
	return &NodesService{config: config}
}

// List EMQX nodes
func (s *NodesService) ListNodes(ctx context.Context, req *ListNodesReq, options ...core.RequestOptionFunc) (*ListNodesResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathNodes
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListNodes] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &ListNodesResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListNodes] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// Get node info
func (s *NodesService) GetNode(ctx context.Context, req *GetNodeReq, options ...core.RequestOptionFunc) (*GetNodeResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathNodes_Node
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetNode] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &GetNodeResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetNode] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

func (s *NodesService) GetNodeMetrics(ctx context.Context, req *GetNodeMetricsReq, options ...core.RequestOptionFunc) (*GetNodeMetricsResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathNodes_Node_Metrics
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetNodeMetrics] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &GetNodeMetricsResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetNodeMetrics] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

func (s *NodesService) GetNodeStats(ctx context.Context, req *GetNodeStatsReq, options ...core.RequestOptionFunc) (*GetNodeStatsResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathNodes_Node_Stats
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetNodeStats] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &GetNodeStatsResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetNodeStats] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}
