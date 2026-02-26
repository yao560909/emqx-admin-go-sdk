package subscriptions

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/yao560909/emqx-admin-go-sdk/core"
)

const (
	ApiPathSubscriptions = "/api/v5/subscriptions"
)

type SubscriptionsService struct {
	config *core.Config
}

func NewService(config *core.Config) *SubscriptionsService {
	return &SubscriptionsService{config: config}
}

// List subscriptions
func (s *SubscriptionsService) ListSubscriptions(ctx context.Context, req *ListSubscriptionsReq, options ...core.RequestOptionFunc) (*SubscriptionListResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathSubscriptions
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListSubscriptions] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &SubscriptionListResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListSubscriptions] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}
