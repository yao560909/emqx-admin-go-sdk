package alarms

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/yao560909/emqx-admin-go-sdk/core"
)

const (
	ApiPathAlarms = "/api/v5/alarms"
)

type AlarmsService struct {
	config *core.Config
}

func NewService(config *core.Config) *AlarmsService {
	return &AlarmsService{config: config}
}

// List currently activated alarms or historical alarms, determined by query parameters.
func (s *AlarmsService) ListAlarms(ctx context.Context, req *ListAlarmsReq, options ...core.RequestOptionFunc) (*AlarmListResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAlarms
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListAlarms] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &AlarmListResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListAlarms] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// Remove all historical alarms.
func (s *AlarmsService) DeleteAlarms(ctx context.Context, req *DeleteAlarmsReq, options ...core.RequestOptionFunc) (*DeleteAlarmsResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAlarms
	apiReq.HttpMethod = "DELETE"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[DeleteAlarms] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &DeleteAlarmsResp{APIResp: apiResp}
	if len(apiResp.RawBody) > 0 {
		err = json.Unmarshal(apiResp.RawBody, resp)
		if err != nil {
			s.config.Logger.Error(ctx, fmt.Sprintf("[DeleteAlarms] fail to unmarshal response body, error: %v", err.Error()))
			return nil, err
		}
	}
	return resp, nil
}
