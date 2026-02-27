package clients

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/yao560909/emqx-admin-go-sdk/core"
)

const (
	ApiPathClients                             = "/api/v5/clients"
	ApiPathClients_Clientid                    = "/api/v5/clients/{clientid}"
	ApiPathClientsKickoutBulk                  = "/api/v5/clients/kickout/bulk"
	ApiPathClients_Clientid_Subscriptions      = "/api/v5/clients/{clientid}/subscriptions"
	ApiPathClients_Clientid_AuthorizationCache = "/api/v5/clients/{clientid}/authorization/cache"
)

type ClientsService struct {
	config *core.Config
}

func NewService(config *core.Config) *ClientsService {
	return &ClientsService{config: config}
}

// List clients
func (s *ClientsService) ListClients(ctx context.Context, req *ListClientsReq, options ...core.RequestOptionFunc) (*ListClientsResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathClients
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListClients] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &ListClientsResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListClients] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// Kick out client by client ID
func (s *ClientsService) KickOutClient(ctx context.Context, req *KickOutClientReq, options ...core.RequestOptionFunc) (*KickOutClientResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathClients_Clientid
	apiReq.HttpMethod = "DELETE"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[KickOutClient] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &KickOutClientResp{APIResp: apiResp}
	if len(apiResp.RawBody) > 0 {
		err = json.Unmarshal(apiResp.RawBody, resp)
		if err != nil {
			s.config.Logger.Error(ctx, fmt.Sprintf("[KickOutClient] fail to unmarshal response body, error: %v", err.Error()))
			return nil, err
		}
	}
	return resp, nil
}

// Kick out a batch of client by client IDs
func (s *ClientsService) BatchKickOutClient(ctx context.Context, req *BatchKickOutClientReq, options ...core.RequestOptionFunc) (*BatchKickOutClientResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathClientsKickoutBulk
	apiReq.HttpMethod = "POST"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[BatchKickOutClient] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &BatchKickOutClientResp{APIResp: apiResp}
	if len(apiResp.RawBody) > 0 {
		err = json.Unmarshal(apiResp.RawBody, resp)
		if err != nil {
			s.config.Logger.Error(ctx, fmt.Sprintf("[BatchKickOutClient] fail to unmarshal response body, error: %v", err.Error()))
			return nil, err
		}
	}
	return resp, nil
}

// Get client subscriptions
func (s *ClientsService) GetClientSubscriptions(ctx context.Context, req *GetClientSubscriptionsReq, options ...core.RequestOptionFunc) (*GetClientSubscriptionsResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathClients_Clientid_Subscriptions
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetClientSubscriptions] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &GetClientSubscriptionsResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetClientSubscriptions] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}

	return resp, nil
}

// Get clients info by client ID
func (s *ClientsService) GetClient(ctx context.Context, req *GetClientReq, options ...core.RequestOptionFunc) (*GetClientResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathClients_Clientid
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetClient] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &GetClientResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetClient] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// Get client authz cache in the cluster.
func (s *ClientsService) GetClientAuthzCache(ctx context.Context, req *GetClientAuthzCacheReq, options ...core.RequestOptionFunc) (*GetClientAuthzCacheResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathClients_Clientid_AuthorizationCache
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetClientAuthzCache] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &GetClientAuthzCacheResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetClientAuthzCache] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// Clean client authz cache in the cluster.
func (s *ClientsService) CleanAuthzCache(ctx context.Context, req *CleanAuthzCacheReq, options ...core.RequestOptionFunc) (*CleanAuthzCacheResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathClients_Clientid_AuthorizationCache
	apiReq.HttpMethod = "DELETE"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[CleanAuthzCache] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &CleanAuthzCacheResp{APIResp: apiResp}
	if len(apiResp.RawBody) > 0 {
		err = json.Unmarshal(apiResp.RawBody, resp)
		if err != nil {
			s.config.Logger.Error(ctx, fmt.Sprintf("[CleanAuthzCache] fail to unmarshal response body, error: %v", err.Error()))
			return nil, err
		}
	}
	return resp, nil
}
