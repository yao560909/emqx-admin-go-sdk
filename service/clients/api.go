package clients

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/yao560909/emqx-admin-go-sdk/core"
)

const (
	ApiPathClients            = "/api/v5/clients"
	ApiPathClients_Clientid   = "/api/v5/clients/{clientid}"
	ApiPathClientsKickoutBulk = "/api/v5/clients/kickout/bulk"
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
