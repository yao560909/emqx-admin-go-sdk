package authorization

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/yao560909/emqx-admin-go-sdk/core"
)

const (
	ApiPathAuthorizationSourcesBuiltInDatabaseRulesUsers = "/api/v5/authorization/sources/built_in_database/rules/users"
)

type AuthorizationService struct {
	config *core.Config
}

func NewService(config *core.Config) *AuthorizationService {
	return &AuthorizationService{config: config}
}

//Show the list of rules for users
func (s *AuthorizationService) ListRulesForUsers(ctx context.Context, req *ListRulesForUsersReq, options ...core.RequestOptionFunc) (*ListRulesForUsersResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAuthorizationSourcesBuiltInDatabaseRulesUsers
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListRulesForUsers] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &ListRulesForUsersResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListRulesForUsers] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}
