package nodes

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/yao560909/emqx-admin-go-sdk/core"
)

const (
	ApiPathNodes = "/api/v5/nodes"
)

type NodesService struct {
	config *core.Config
}

func NewService(config *core.Config) *NodesService {
	return &NodesService{config: config}
}

func (n *NodesService) ListNodes(ctx context.Context, req *ListNodesReq, options ...core.RequestOptionFunc) (*ListNodesResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathNodes
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(n.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		n.config.Logger.Error(ctx, fmt.Sprintf("[ListNodes] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &ListNodesResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		n.config.Logger.Error(ctx, fmt.Sprintf("[ListNodes] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}
