package banned

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/yao560909/emqx-admin-go-sdk/core"
)

const (
	ApiPathBanned         = "/api/v5/banned"
	ApiPathBanned_As__Who = "/api/v5/banned/{as}/{who}"
)

type BannedService struct {
	config *core.Config
}

func NewService(config *core.Config) *BannedService {
	return &BannedService{config: config}
}

// Add a client ID, username or IP address to the blacklist.
func (s *BannedService) Banned(ctx context.Context, req *BannedReq, options ...core.RequestOptionFunc) (*BannedResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathBanned
	apiReq.HttpMethod = "POST"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[Banned] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &BannedResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[Banned] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// Remove a client ID, username or IP address from the blacklist.
func (s *BannedService) RemoveBanned(ctx context.Context, req *RemoveBannedReq, options ...core.RequestOptionFunc) (*RemoveBannedResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathBanned_As__Who
	apiReq.HttpMethod = "DELETE"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[RemoveBanned] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &RemoveBannedResp{APIResp: apiResp}
	if len(apiResp.RawBody) > 0 {
		err = json.Unmarshal(apiResp.RawBody, resp)
		if err != nil {
			s.config.Logger.Error(ctx, fmt.Sprintf("[RemoveBanned] fail to unmarshal response body, error: %v", err.Error()))
			return nil, err
		}
	}
	return resp, nil
}

// Clear all banned data.
func (s *BannedService) ClearAllBanned(ctx context.Context, req *ClearAllBannedReq, options ...core.RequestOptionFunc) (*ClearAllBannedResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathBanned
	apiReq.HttpMethod = "DELETE"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ClearAllBanned] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &ClearAllBannedResp{APIResp: apiResp}
	if len(apiResp.RawBody) > 0 {
		err = json.Unmarshal(apiResp.RawBody, resp)
		if err != nil {
			s.config.Logger.Error(ctx, fmt.Sprintf("[ClearAllBanned] fail to unmarshal response body, error: %v", err.Error()))
			return nil, err
		}
	}
	return resp, nil
}

// List all currently banned client IDs, usernames and IP addresses.
func (s *BannedService) ListCurrentlyBanned(ctx context.Context, req *ListCurrentlyBannedReq, options ...core.RequestOptionFunc) (*ListCurrentlyBannedResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathBanned
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListCurrentlyBanned] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &ListCurrentlyBannedResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[BanListCurrentlyBannedned] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}
