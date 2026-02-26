package cluster

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/yao560909/emqx-admin-go-sdk/core"
)

const (
	ApiPathCluster         = "/api/v5/cluster"
	ApiPathClusterTopology = "/api/v5/cluster/topology"
)

type ClusterService struct {
	config *core.Config
}

func NewService(config *core.Config) *ClusterService {
	return &ClusterService{config: config}
}

// Get cluster info
func (s *ClusterService) ListCluster(ctx context.Context, req *ListClusterReq, options ...core.RequestOptionFunc) (*ListClusterResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathCluster
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListCluster] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &ListClusterResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListCluster] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// Get cluster topology
func (s *ClusterService) GetClusterTopology(ctx context.Context, req *GetClusterTopologyReq, options ...core.RequestOptionFunc) (*GetClusterTopologyResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathClusterTopology
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetClusterTopology] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &GetClusterTopologyResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetClusterTopology] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}
