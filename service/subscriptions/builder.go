package subscriptions

import (
	"github.com/yao560909/emqx-admin-go-sdk/core"
)

// --------------------------------
type ListSubscriptionsReq struct {
	apiReq *core.APIReq
}

type ListSubscriptionsReqBuilder struct {
	apiReq *core.APIReq
}

func NewListSubscriptionsReqBuilder() *ListSubscriptionsReqBuilder {
	builder := &ListSubscriptionsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *ListSubscriptionsReqBuilder) Build() *ListSubscriptionsReq {
	req := &ListSubscriptionsReq{}
	req.apiReq = b.apiReq
	return req
}

func (b *ListSubscriptionsReqBuilder) Clientid(clientID string) *ListSubscriptionsReqBuilder {
	b.apiReq.QueryParams.Set("clientid", clientID)
	return b
}

func (b *ListSubscriptionsReqBuilder) Topic(topic string) *ListSubscriptionsReqBuilder {
	b.apiReq.QueryParams.Set("topic", topic)
	return b
}

func (b *ListSubscriptionsReqBuilder) MatchTopic(matchTopic string) *ListSubscriptionsReqBuilder {
	b.apiReq.QueryParams.Set("match_topic", matchTopic)
	return b
}

func (b *ListSubscriptionsReqBuilder) ShareGroup(shareGroup string) *ListSubscriptionsReqBuilder {
	b.apiReq.QueryParams.Set("share_group", shareGroup)
	return b
}

func (b *ListSubscriptionsReqBuilder) Node(node string) *ListSubscriptionsReqBuilder {
	b.apiReq.QueryParams.Set("node", node)
	return b
}

func (b *ListSubscriptionsReqBuilder) Page(page string) *ListSubscriptionsReqBuilder {
	b.apiReq.QueryParams.Set("page", page)
	return b
}

func (b *ListSubscriptionsReqBuilder) Limit(limit string) *ListSubscriptionsReqBuilder {
	b.apiReq.QueryParams.Set("limit", limit)
	return b
}

type SubscriptionListResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*Subscription `json:"data"`
	Meta Meta            `json:"meta"`
}
