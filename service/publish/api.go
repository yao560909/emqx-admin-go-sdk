package publish

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/yao560909/emqx-admin-go-sdk/core"
)

const (
	ApiPathPublish     = "/api/v5/publish"
	ApiPathPublishBulk = "/api/v5/publish/bulk"
)

type PublishService struct {
	config *core.Config
}

func NewService(config *core.Config) *PublishService {
	return &PublishService{config: config}
}

// Publish a message
func (s *PublishService) Publish(ctx context.Context, req *PublishReq, options ...core.RequestOptionFunc) (*PublishResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathPublish
	apiReq.HttpMethod = "POST"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[Publish] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &PublishResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[Publish] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// Publish a batch of messages
func (s *PublishService) PublishABatchOfMessages(ctx context.Context, req *PublishReq, options ...core.RequestOptionFunc) (*PublishResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathPublishBulk
	apiReq.HttpMethod = "POST"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[PublishABatchOfMessages] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &PublishResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[PublishABatchOfMessages] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}
