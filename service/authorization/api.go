package authorization

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/yao560909/emqx-admin-go-sdk/core"
)

const (
	ApiPathAuthorizationSourcesBuiltInDatabaseRulesUsers           = "/api/v5/authorization/sources/built_in_database/rules/users"
	ApiPathAuthorizationSourcesBuiltInDatabaseRulesUsersUsername   = "/api/v5/authorization/sources/built_in_database/rules/users/{username}"
	ApiPathAuthorizationSourcesBuiltInDatabaseRulesClients         = "/api/v5/authorization/sources/built_in_database/rules/clients"
	ApiPathAuthorizationSourcesBuiltInDatabaseRulesClientsClientId = "/api/v5/authorization/sources/built_in_database/rules/clients/{clientid}"
	ApiPathAuthorizationSourcesTypeStatus                          = "/api/v5/authorization/sources/{type}/status"
	ApiPathAuthorizationSources                                    = "/api/v5/authorization/sources"
)

type AuthorizationService struct {
	config *core.Config
}

func NewService(config *core.Config) *AuthorizationService {
	return &AuthorizationService{config: config}
}

// Show the list of rules for users
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

// Add new rule for 'username'
func (s *AuthorizationService) AddRule(ctx context.Context, req *AddRuleReq, options ...core.RequestOptionFunc) (*AddRuleResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAuthorizationSourcesBuiltInDatabaseRulesUsers
	apiReq.HttpMethod = "POST"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[AddRule] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &AddRuleResp{APIResp: apiResp}
	if len(apiResp.RawBody) > 0 {
		err = json.Unmarshal(apiResp.RawBody, resp)
		if err != nil {
			s.config.Logger.Error(ctx, fmt.Sprintf("[AddRule] fail to unmarshal response body, error: %v", err.Error()))
			return nil, err
		}
	}
	return resp, nil
}

// Show the list of rules for clients
func (s *AuthorizationService) ListRulesForClients(ctx context.Context, req *ListRulesForClientsReq, options ...core.RequestOptionFunc) (*ListRulesForClientsResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAuthorizationSourcesBuiltInDatabaseRulesClients
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListRulesForClients] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &ListRulesForClientsResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListRulesForClients] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// Add new rule for 'clientid'
func (s *AuthorizationService) AddRuleForClients(ctx context.Context, req *AddRuleForClientsReq, options ...core.RequestOptionFunc) (*AddRuleForClientsResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAuthorizationSourcesBuiltInDatabaseRulesClients
	apiReq.HttpMethod = "POST"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[AddRuleForClients] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &AddRuleForClientsResp{APIResp: apiResp}
	if len(apiResp.RawBody) > 0 {
		err = json.Unmarshal(apiResp.RawBody, resp)
		if err != nil {
			s.config.Logger.Error(ctx, fmt.Sprintf("[AddRuleForClients] fail to unmarshal response body, error: %v", err.Error()))
			return nil, err
		}
	}
	return resp, nil
}

// Get rule for 'clientid'
func (s *AuthorizationService) GetRuleForClient(ctx context.Context, req *GetRuleForClientReq, options ...core.RequestOptionFunc) (*GetRuleForClientResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAuthorizationSourcesBuiltInDatabaseRulesClientsClientId
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetRuleForClient] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &GetRuleForClientResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetRuleForClient] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// Set rule for 'clientid'
func (s *AuthorizationService) SetRuleForClient(ctx context.Context, req *SetRuleForClientReq, options ...core.RequestOptionFunc) (*SetRuleForClientResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAuthorizationSourcesBuiltInDatabaseRulesClientsClientId
	apiReq.HttpMethod = "PUT"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[SetRuleForClient] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &SetRuleForClientResp{APIResp: apiResp}
	if len(apiResp.RawBody) > 0 {
		err = json.Unmarshal(apiResp.RawBody, resp)
		if err != nil {
			s.config.Logger.Error(ctx, fmt.Sprintf("[SetRuleForClient] fail to unmarshal response body, error: %v", err.Error()))
			return nil, err
		}
	}
	return resp, nil
}

// Delete rule for 'clientid'
func (s *AuthorizationService) DeleteRuleForClient(ctx context.Context, req *DeleteRuleForClientReq, options ...core.RequestOptionFunc) (*DeleteRuleForClientResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAuthorizationSourcesBuiltInDatabaseRulesClientsClientId
	apiReq.HttpMethod = "DELETE"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[DeleteRuleForClient] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &DeleteRuleForClientResp{APIResp: apiResp}
	if len(apiResp.RawBody) > 0 {
		err = json.Unmarshal(apiResp.RawBody, resp)
		if err != nil {
			s.config.Logger.Error(ctx, fmt.Sprintf("[DeleteRuleForClient] fail to unmarshal response body, error: %v", err.Error()))
			return nil, err
		}
	}
	return resp, nil
}

// Get a authorization source status
func (s *AuthorizationService) GetAuthorizationSourceStatus(ctx context.Context, req *GetAuthorizationSourceStatusReq, options ...core.RequestOptionFunc) (*GetAuthorizationSourceStatusResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAuthorizationSourcesTypeStatus
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetAuthorizationSourceStatus] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &GetAuthorizationSourceStatusResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetAuthorizationSourceStatus] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// List all authorization sources
func (s *AuthorizationService) ListAuthorizationSources(ctx context.Context, req *ListAuthorizationSourcesReq, options ...core.RequestOptionFunc) (*ListAuthorizationSourcesResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAuthorizationSources
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListAuthorizationSources] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &ListAuthorizationSourcesResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListAuthorizationSources] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// Get rule for 'username'
func (s *AuthorizationService) GetRuleForUser(ctx context.Context, req *GetRuleForUserReq, options ...core.RequestOptionFunc) (*GetRuleForUserResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAuthorizationSourcesBuiltInDatabaseRulesUsersUsername
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetRuleForUser] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &GetRuleForUserResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetRuleForUser] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// Set rule for 'username'
func (s *AuthorizationService) SetRuleForUser(ctx context.Context, req *SetRuleForUserReq, options ...core.RequestOptionFunc) (*SetRuleForUserResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAuthorizationSourcesBuiltInDatabaseRulesUsersUsername
	apiReq.HttpMethod = "PUT"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[SetRuleForUser] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &SetRuleForUserResp{APIResp: apiResp}
	if len(apiResp.RawBody) > 0 {
		err = json.Unmarshal(apiResp.RawBody, resp)
		if err != nil {
			s.config.Logger.Error(ctx, fmt.Sprintf("[SetRuleForUser] fail to unmarshal response body, error: %v", err.Error()))
			return nil, err
		}
	}
	return resp, nil
}

// Delete rule for 'username'
func (s *AuthorizationService) DeleteRuleForUser(ctx context.Context, req *DeleteRuleForUserReq, options ...core.RequestOptionFunc) (*DeleteRuleForUserResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAuthorizationSourcesBuiltInDatabaseRulesUsersUsername
	apiReq.HttpMethod = "DELETE"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[DeleteRuleForUser] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &DeleteRuleForUserResp{APIResp: apiResp}
	if len(apiResp.RawBody) > 0 {
		err = json.Unmarshal(apiResp.RawBody, resp)
		if err != nil {
			s.config.Logger.Error(ctx, fmt.Sprintf("[DeleteRuleForUser] fail to unmarshal response body, error: %v", err.Error()))
			return nil, err
		}
	}
	return resp, nil
}
