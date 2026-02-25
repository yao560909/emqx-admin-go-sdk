package topics

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/yao560909/emqx-admin-go-sdk/core"
)

const (
	ApiPathTopics       = "/api/v5/topics"
	ApiPathTopics_Topic = "/api/v5/topics/{topic}"
)

type TopicsService struct {
	config *core.Config
}

func NewService(config *core.Config) *TopicsService {
	return &TopicsService{config: config}
}

// List topics
func (s *TopicsService) ListTopics(ctx context.Context, req *ListTopicsReq, options ...core.RequestOptionFunc) (*ListTopicsResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathTopics
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListTopics] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &ListTopicsResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListTopics] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// Get topic
func (s *TopicsService) GetTopic(ctx context.Context, req *GetTopicReq, options ...core.RequestOptionFunc) (*GetTopicResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathTopics_Topic
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetTopic] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &GetTopicResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetTopic] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}
