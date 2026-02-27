package clients

import (
	"encoding/json"

	"github.com/yao560909/emqx-admin-go-sdk/core"
)

// -------------------------
type ListClientsReq struct {
	apiReq *core.APIReq
}

type ListClientsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*Client `json:"data"`
	Meta *Meta     `json:"meta"`
}

type ListClientsReqBuilder struct {
	apiReq *core.APIReq
}

func NewListClientsReqBuilder() *ListClientsReqBuilder {
	builder := &ListClientsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *ListClientsReqBuilder) Build() *ListClientsReq {
	req := &ListClientsReq{}
	req.apiReq = b.apiReq
	return req
}

func (b *ListClientsReqBuilder) Limit(limit string) *ListClientsReqBuilder {
	b.apiReq.QueryParams.Set("limit", limit)
	return b
}

func (b *ListClientsReqBuilder) Page(page string) *ListClientsReqBuilder {
	b.apiReq.QueryParams.Set("page", page)
	return b
}

func (b *ListClientsReqBuilder) Node(node string) *ListClientsReqBuilder {
	b.apiReq.QueryParams.Set("node", node)
	return b
}

func (b *ListClientsReqBuilder) Username(username string) *ListClientsReqBuilder {
	b.apiReq.QueryParams.Set("username", username)
	return b
}

func (b *ListClientsReqBuilder) LikeUsername(likeUsername string) *ListClientsReqBuilder {
	b.apiReq.QueryParams.Set("like_username", likeUsername)
	return b
}

func (b *ListClientsReqBuilder) LikeClientid(likeClientid string) *ListClientsReqBuilder {
	b.apiReq.QueryParams.Set("like_clientid", likeClientid)
	return b
}

func (b *ListClientsReqBuilder) IpAddress(ipAddress string) *ListClientsReqBuilder {
	b.apiReq.QueryParams.Set("ip_address", ipAddress)
	return b
}

/*
*
conn_state
Enum: "connected" "idle" "disconnected"
*/
func (b *ListClientsReqBuilder) ConnState(connState string) *ListClientsReqBuilder {
	b.apiReq.QueryParams.Set("conn_state", connState)
	return b
}

/*
*
gte_created_at
Search client session creation time by greater than or equal method, rfc3339 or timestamp(millisecond)
*/
func (b *ListClientsReqBuilder) GteCreatedAt(gteCreatedAt string) *ListClientsReqBuilder {
	b.apiReq.QueryParams.Set("gte_created_at", gteCreatedAt)
	return b
}

/*
*
clean_start
boolean
Whether the client uses a new session
*/
func (b *ListClientsReqBuilder) CleanStart(cleanStart string) *ListClientsReqBuilder {
	b.apiReq.QueryParams.Set("clean_start", cleanStart)
	return b
}

func (b *ListClientsReqBuilder) ProtoVer(protoVer string) *ListClientsReqBuilder {
	b.apiReq.QueryParams.Set("proto_ver", protoVer)
	return b
}

/*
*
lte_created_at
Search client session creation time by less than or equal method, rfc3339 or timestamp(millisecond)
*/
func (b *ListClientsReqBuilder) LteCreatedAt(lteCreatedAt string) *ListClientsReqBuilder {
	b.apiReq.QueryParams.Set("lte_created_at", lteCreatedAt)
	return b
}

/*
*
gte_connected_at
Search client connection creation time by greater than or equal method, rfc3339 or timestamp(epoch millisecond)
*/
func (b *ListClientsReqBuilder) GteConnectedAt(gteConnectedAt string) *ListClientsReqBuilder {
	b.apiReq.QueryParams.Set("gte_connected_at", gteConnectedAt)
	return b
}

/*
*
lte_connected_at
Search client connection creation time by greater than or equal method, rfc3339 or timestamp(epoch millisecond)
*/
func (b *ListClientsReqBuilder) LteConnectedAt(lteConnectedAt string) *ListClientsReqBuilder {
	b.apiReq.QueryParams.Set("lte_connected_at", lteConnectedAt)
	return b
}

// ---------------------------
type KickOutClientReq struct {
	apiReq *core.APIReq
}

type KickOutClientReqBuilder struct {
	apiReq *core.APIReq
}

func NewKickOutClientReqBuilder() *KickOutClientReqBuilder {
	builder := &KickOutClientReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *KickOutClientReqBuilder) Clientid(clientid string) *KickOutClientReqBuilder {
	b.apiReq.PathParams.Set("clientid", clientid)
	return b
}

func (b *KickOutClientReqBuilder) Build() *KickOutClientReq {
	req := &KickOutClientReq{}
	req.apiReq = b.apiReq
	return req
}

type KickOutClientResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

// --------------------------------
type BatchKickOutClientReq struct {
	apiReq *core.APIReq
}

type BatchKickOutClientReqBuilder struct {
	apiReq *core.APIReq
}

func NewBatchKickOutClientReqBuilder() *BatchKickOutClientReqBuilder {
	builder := &BatchKickOutClientReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *BatchKickOutClientReqBuilder) Build() *BatchKickOutClientReq {
	req := &BatchKickOutClientReq{}
	req.apiReq = b.apiReq
	return req
}

/*
*
Array[
string
]
*/
func (b *BatchKickOutClientReqBuilder) ClientIdArray(clientIdArray []string) *BatchKickOutClientReqBuilder {
	b.apiReq.Body = clientIdArray
	return b
}

type BatchKickOutClientResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}

type GetClientSubscriptionsReq struct {
	apiReq *core.APIReq
}

type GetClientSubscriptionsReqBuilder struct {
	apiReq *core.APIReq
}

func NewGetClientSubscriptionsReqBuilder() *GetClientSubscriptionsReqBuilder {
	builder := &GetClientSubscriptionsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *GetClientSubscriptionsReqBuilder) Clientid(clientid string) *GetClientSubscriptionsReqBuilder {
	b.apiReq.PathParams.Set("clientid", clientid)
	return b
}

func (b *GetClientSubscriptionsReqBuilder) Build() *GetClientSubscriptionsReq {
	req := &GetClientSubscriptionsReq{}
	req.apiReq = b.apiReq
	return req
}

type GetClientSubscriptionsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*Subscription `json:"-"`
}

func (resp *GetClientSubscriptionsResp) UnmarshalJSON(b []byte) error {
	var subscriptions []*Subscription
	if err := json.Unmarshal(b, &subscriptions); err == nil {
		resp.Data = subscriptions
		return nil
	}
	type alias GetClientSubscriptionsResp
	return json.Unmarshal(b, (*alias)(resp))
}

type GetClientReq struct {
	apiReq *core.APIReq
}

type GetClientReqBuilder struct {
	apiReq *core.APIReq
}

func NewGetClientReqBuilder() *GetClientReqBuilder {
	builder := &GetClientReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *GetClientReqBuilder) Clientid(clientid string) *GetClientReqBuilder {
	b.apiReq.PathParams.Set("clientid", clientid)
	return b
}

func (b *GetClientReqBuilder) Build() *GetClientReq {
	req := &GetClientReq{}
	req.apiReq = b.apiReq
	return req
}

type GetClientResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Client
}
