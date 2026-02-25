package authorization

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/yao560909/emqx-admin-go-sdk/core"
)

const (
	ApiPathAuthorizationSourcesBuiltInDatabaseRulesAll             = "/api/v5/authorization/sources/built_in_database/rules/all"
	ApiPathAuthorizationSourcesBuiltInDatabaseRules                = "/api/v5/authorization/sources/built_in_database/rules"
	ApiPathAuthorizationSourcesBuiltInDatabaseRulesUsers           = "/api/v5/authorization/sources/built_in_database/rules/users"
	ApiPathAuthorizationSourcesBuiltInDatabaseRulesUsersUsername   = "/api/v5/authorization/sources/built_in_database/rules/users/{username}"
	ApiPathAuthorizationSourcesBuiltInDatabaseRulesClients         = "/api/v5/authorization/sources/built_in_database/rules/clients"
	ApiPathAuthorizationSourcesBuiltInDatabaseRulesClientsClientId = "/api/v5/authorization/sources/built_in_database/rules/clients/{clientid}"
	ApiPathAuthorizationSourcesTypeStatus                          = "/api/v5/authorization/sources/{type}/status"
	ApiPathAuthorizationSourcesTypeMove                            = "/api/v5/authorization/sources/{type}/move"
	ApiPathAuthorizationSourcesType                                = "/api/v5/authorization/sources/{type}"
	ApiPathAuthorizationSources                                    = "/api/v5/authorization/sources"
	ApiPathAuthorizationCache                                      = "/api/v5/authorization/cache"
	ApiPathAuthorizationSettings                                   = "/api/v5/authorization/settings"
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

// Delete all rules for all 'users', 'clients' and 'all'
func (s *AuthorizationService) DeleteAllRules(ctx context.Context, req *DeleteAllRulesReq, options ...core.RequestOptionFunc) (*DeleteAllRulesResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAuthorizationSourcesBuiltInDatabaseRules
	apiReq.HttpMethod = "DELETE"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[DeleteAllRules] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &DeleteAllRulesResp{APIResp: apiResp}
	if len(apiResp.RawBody) > 0 {
		err = json.Unmarshal(apiResp.RawBody, resp)
		if err != nil {
			s.config.Logger.Error(ctx, fmt.Sprintf("[DeleteAllRules] fail to unmarshal response body, error: %v", err.Error()))
			return nil, err
		}
	}
	return resp, nil
}

// Delete rules for 'all'
func (s *AuthorizationService) DeleteRulesForAll(ctx context.Context, req *DeleteRulesForAllReq, options ...core.RequestOptionFunc) (*DeleteRulesForAllResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAuthorizationSourcesBuiltInDatabaseRulesAll
	apiReq.HttpMethod = "DELETE"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[DeleteRulesForAll] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &DeleteRulesForAllResp{APIResp: apiResp}
	if len(apiResp.RawBody) > 0 {
		err = json.Unmarshal(apiResp.RawBody, resp)
		if err != nil {
			s.config.Logger.Error(ctx, fmt.Sprintf("[DeleteRulesForAll] fail to unmarshal response body, error: %v", err.Error()))
			return nil, err
		}
	}
	return resp, nil
}

// Create/Update the list of rules for 'all'
func (s *AuthorizationService) SetRulesForAll(ctx context.Context, req *SetRulesForAllReq, options ...core.RequestOptionFunc) (*SetRulesForAllResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAuthorizationSourcesBuiltInDatabaseRulesAll
	apiReq.HttpMethod = "PUT"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[SetRulesForAll] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &SetRulesForAllResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[SetRulesForAll] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// Show the list of rules for 'all'
func (s *AuthorizationService) ListRulesForAll(ctx context.Context, req *ListRulesForAllReq, options ...core.RequestOptionFunc) (*ListRulesForAllResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAuthorizationSourcesBuiltInDatabaseRulesAll
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListRulesForAll] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &ListRulesForAllResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[ListRulesForAll] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// Change the exection order of sources
func (s *AuthorizationService) MoveAuthorizationSource(ctx context.Context, req *MoveAuthorizationSourceReq, options ...core.RequestOptionFunc) (*MoveAuthorizationSourceResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAuthorizationSourcesTypeMove
	apiReq.HttpMethod = "POST"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[MoveAuthorizationSource] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &MoveAuthorizationSourceResp{APIResp: apiResp}
	if len(apiResp.RawBody) > 0 {
		err = json.Unmarshal(apiResp.RawBody, resp)
		if err != nil {
			s.config.Logger.Error(ctx, fmt.Sprintf("[MoveAuthorizationSource] fail to unmarshal response body, error: %v", err.Error()))
			return nil, err
		}
	}
	return resp, nil
}

// Get a authorization source
func (s *AuthorizationService) GetAuthorizationSource(ctx context.Context, req *GetAuthorizationSourceReq, options ...core.RequestOptionFunc) (*GetAuthorizationSourceResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAuthorizationSourcesType
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetAuthorizationSource] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &GetAuthorizationSourceResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetAuthorizationSource] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// Delete source
func (s *AuthorizationService) DeleteAuthorizationSource(ctx context.Context, req *DeleteAuthorizationSourceReq, options ...core.RequestOptionFunc) (*DeleteAuthorizationSourceResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAuthorizationSourcesType
	apiReq.HttpMethod = "DELETE"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[DeleteAuthorizationSource] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &DeleteAuthorizationSourceResp{APIResp: apiResp}
	if len(apiResp.RawBody) > 0 {
		err = json.Unmarshal(apiResp.RawBody, resp)
		if err != nil {
			s.config.Logger.Error(ctx, fmt.Sprintf("[DeleteAuthorizationSource] fail to unmarshal response body, error: %v", err.Error()))
			return nil, err
		}
	}
	return resp, nil
}

// Clean all authorization cache in the cluster.
func (s *AuthorizationService) CleanAuthorizationCache(ctx context.Context, req *CleanAuthorizationCacheReq, options ...core.RequestOptionFunc) (*CleanAuthorizationCacheResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAuthorizationCache
	apiReq.HttpMethod = "DELETE"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[CleanAuthorizationCache] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &CleanAuthorizationCacheResp{APIResp: apiResp}
	if len(apiResp.RawBody) > 0 {
		err = json.Unmarshal(apiResp.RawBody, resp)
		if err != nil {
			s.config.Logger.Error(ctx, fmt.Sprintf("[CleanAuthorizationCache] fail to unmarshal response body, error: %v", err.Error()))
			return nil, err
		}
	}
	return resp, nil
}

// Get authorization settings
func (s *AuthorizationService) GetAuthorizationSettings(ctx context.Context, req *GetAuthorizationSettingsReq, options ...core.RequestOptionFunc) (*GetAuthorizationSettingsResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAuthorizationSettings
	apiReq.HttpMethod = "GET"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetAuthorizationSettings] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &GetAuthorizationSettingsResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[GetAuthorizationSettings] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}

// Update authorization settings
func (s *AuthorizationService) UpdateAuthorizationSettings(ctx context.Context, req *UpdateAuthorizationSettingsReq, options ...core.RequestOptionFunc) (*UpdateAuthorizationSettingsResp, error) {
	apiReq := req.apiReq
	apiReq.ApiPath = ApiPathAuthorizationSettings
	apiReq.HttpMethod = "PUT"
	requester := core.NewRequester(s.config)
	apiResp, err := requester.DoRequest(apiReq, options...)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[UpdateAuthorizationSettings] fail to invoke api, error: %v", err.Error()))
		return nil, err
	}
	resp := &UpdateAuthorizationSettingsResp{APIResp: apiResp}
	err = json.Unmarshal(apiResp.RawBody, resp)
	if err != nil {
		s.config.Logger.Error(ctx, fmt.Sprintf("[UpdateAuthorizationSettings] fail to unmarshal response body, error: %v", err.Error()))
		return nil, err
	}
	return resp, nil
}
