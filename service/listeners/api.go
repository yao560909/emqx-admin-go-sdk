package listeners

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/yao560909/emqx-admin-go-sdk/core"
)

const (
	ApiPathListeners       = "/api/v5/listeners"
	ApiPathListenersStatus = "/api/v5/listeners_status"
)

type ListenersService struct {
	config *core.Config
}

func NewService(config *core.Config) *ListenersService {
	return &ListenersService{config: config}
}

// List all running node's listeners for the specified type.
func (s *ListenersService) ListListeners(ctx context.Context, req *ListListenersReq, options ...core.RequestOptionFunc) (*ListListenersResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathListeners
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListListeners] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &ListListenersResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListListeners] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// List all running node's listeners live status. group by listener type
func (s *ListenersService) ListListenersStatus(ctx context.Context, req *ListenerStatusReq, options ...core.RequestOptionFunc) (*ListenerStatusResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathListenersStatus
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListListenersStatus] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &ListenerStatusResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListListenersStatus] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}
