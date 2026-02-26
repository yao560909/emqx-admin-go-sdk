package cluster

import (
	"encoding/json"

	"github.com/yao560909/emqx-admin-go-sdk/core"
)

type ListClusterReq struct {
	apiReq *core.APIReq
}

type ListClusterReqBuilder struct {
	apiReq *core.APIReq
}

// NewListClusterReqBuilder creates a new ListClusterReqBuilder
func NewListClusterReqBuilder() *ListClusterReqBuilder {
	builder := &ListClusterReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

// Build builds the ListClusterReq
func (b *ListClusterReqBuilder) Build() *ListClusterReq {
	req := &ListClusterReq{}
	req.apiReq = b.apiReq
	return req
}

type ListClusterResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Cluster
}

type GetClusterTopologyReq struct {
	apiReq *core.APIReq
}

// GetClusterTopologyReqBuilder represents the builder for GetClusterTopologyReq

type GetClusterTopologyReqBuilder struct {
	apiReq *core.APIReq
}

// NewGetClusterTopologyReqBuilder creates a new GetClusterTopologyReqBuilder
func NewGetClusterTopologyReqBuilder() *GetClusterTopologyReqBuilder {
	builder := &GetClusterTopologyReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

// Build builds the GetClusterTopologyReq
func (b *GetClusterTopologyReqBuilder) Build() *GetClusterTopologyReq {
	req := &GetClusterTopologyReq{}
	req.apiReq = b.apiReq
	return req
}

// GetClusterTopologyResp represents the response for getting cluster topology

type GetClusterTopologyResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*ClusterTopology `json:"-"`
}

func (resp *GetClusterTopologyResp) UnmarshalJSON(b []byte) error {
	var clusterTopology []*ClusterTopology
	if err := json.Unmarshal(b, &clusterTopology); err == nil {
		resp.Data = clusterTopology
		return nil
	}
	type alias GetClusterTopologyResp
	return json.Unmarshal(b, (*alias)(resp))
}

// ClusterTopology represents the cluster topology

type ClusterTopology struct {
	CoreNode       string           `json:"core_node"`
	ReplicantNodes []*ReplicantNode `json:"replicant_nodes"`
}

// ReplicantNode represents the replicant node

type ReplicantNode struct {
	Node    string `json:"node"`
	Streams int    `json:"streams"`
}
