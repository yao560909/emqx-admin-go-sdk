package metrics

import (
	"encoding/json"

	"github.com/yao560909/emqx-admin-go-sdk/core"
)

// -------------------
type StatsReq struct {
	apiReq *core.APIReq
}

type StatsReqBuilder struct {
	apiReq *core.APIReq
}

func NewStatsReqBuilder() *StatsReqBuilder {
	builder := &StatsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

/*
*
boolean
Calculation aggregate for all nodes
*/
func (b *StatsReqBuilder) Aggregate(aggregate string) *StatsReqBuilder {
	b.apiReq.QueryParams.Set("aggregate", aggregate)
	return b
}

func (b *StatsReqBuilder) Build() *StatsReq {
	req := &StatsReq{}
	req.apiReq = b.apiReq
	return req
}

type StatsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Stats
	Data []*Stats `json:"-"`
}

func (resp *StatsResp) UnmarshalJSON(b []byte) error {
	var stats []*Stats
	if err := json.Unmarshal(b, &stats); err == nil {
		resp.Data = stats
		return nil
	}
	type alias StatsResp
	return json.Unmarshal(b, (*alias)(resp))
}

// ---------------------
type MetricsReq struct {
	apiReq *core.APIReq
}

type MetricsReqBuilder struct {
	apiReq *core.APIReq
}

func NewMetricsReqBuilder() *MetricsReqBuilder {
	builder := &MetricsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *MetricsReqBuilder) Build() *MetricsReq {
	req := &MetricsReq{}
	req.apiReq = b.apiReq
	return req
}

type MetricsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*Metrics `json:"-"`
}

func (resp *MetricsResp) UnmarshalJSON(b []byte) error {
	var metrics []*Metrics
	if err := json.Unmarshal(b, &metrics); err == nil {
		resp.Data = metrics
		return nil
	}
	type alias MetricsResp
	return json.Unmarshal(b, (*alias)(resp))
}

// --------------------------------------
type ListMonitorForTheClusterReq struct {
	apiReq *core.APIReq
}

type ListMonitorForTheClusterReqBuilder struct {
	apiReq *core.APIReq
}

func NewListMonitorForTheClusterReqBuilder() *ListMonitorForTheClusterReqBuilder {
	builder := &ListMonitorForTheClusterReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

/*
*
integer >= 1
Example: latest=300
The latest N seconds data. Like 300 for 5 min.
*/
func (b *ListMonitorForTheClusterReqBuilder) Latest(latest string) *ListMonitorForTheClusterReqBuilder {
	b.apiReq.QueryParams.Set("latest", latest)
	return b
}

func (b *ListMonitorForTheClusterReqBuilder) Build() *ListMonitorForTheClusterReq {
	req := &ListMonitorForTheClusterReq{}
	req.apiReq = b.apiReq
	return req
}

type ListMonitorForTheClusterResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*Monitor `json:"-"`
}

func (resp *ListMonitorForTheClusterResp) UnmarshalJSON(b []byte) error {
	var monitors []*Monitor
	if err := json.Unmarshal(b, &monitors); err == nil {
		resp.Data = monitors
		return nil
	}
	type alias ListMonitorForTheClusterResp
	return json.Unmarshal(b, (*alias)(resp))
}

// --------------------------------
type ListMonitorForNodeReq struct {
	apiReq *core.APIReq
}

type ListMonitorForNodeReqBuilder struct {
	apiReq *core.APIReq
}

func NewListMonitorForNodeReqBuilder() *ListMonitorForNodeReqBuilder {
	builder := &ListMonitorForNodeReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

/*
*
integer >= 1
Example: latest=300
The latest N seconds data. Like 300 for 5 min.
*/
func (b *ListMonitorForNodeReqBuilder) Latest(latest string) *ListMonitorForNodeReqBuilder {
	b.apiReq.QueryParams.Set("latest", latest)
	return b
}

/*
*
string
Example: emqx@127.0.0.1
EMQX node name.
*/
func (b *ListMonitorForNodeReqBuilder) Node(node string) *ListMonitorForNodeReqBuilder {
	b.apiReq.PathParams.Set("node", node)
	return b
}

func (b *ListMonitorForNodeReqBuilder) Build() *ListMonitorForNodeReq {
	req := &ListMonitorForNodeReq{}
	req.apiReq = b.apiReq
	return req
}

type ListMonitorForNodeResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*Monitor `json:"-"`
}

func (resp *ListMonitorForNodeResp) UnmarshalJSON(b []byte) error {
	var monitors []*Monitor
	if err := json.Unmarshal(b, &monitors); err == nil {
		resp.Data = monitors
		return nil
	}
	type alias ListMonitorForNodeResp
	return json.Unmarshal(b, (*alias)(resp))
}

// -----------------------------------------
type CurrentMonitorForTheClusterReq struct {
	apiReq *core.APIReq
}

type CurrentMonitorForTheClusterReqBuilder struct {
	apiReq *core.APIReq
}

func NewCurrentMonitorForTheClusterReqBuilder() *CurrentMonitorForTheClusterReqBuilder {
	builder := &CurrentMonitorForTheClusterReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *CurrentMonitorForTheClusterReqBuilder) Build() *CurrentMonitorForTheClusterReq {
	req := &CurrentMonitorForTheClusterReq{}
	req.apiReq = b.apiReq
	return req
}

type CurrentMonitorForTheClusterResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	MonitorCurrent
}

// -----------------------------------
type CurrentMonitorForNodeRep struct {
	apiReq *core.APIReq
}

type CurrentMonitorForNodeRepBuilder struct {
	apiReq *core.APIReq
}

func NewCurrentMonitorForNodeRepBuilder() *CurrentMonitorForNodeRepBuilder {
	builder := &CurrentMonitorForNodeRepBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

/*
*
string
Example: emqx@127.0.0.1
EMQX node name.
*/
func (b *CurrentMonitorForNodeRepBuilder) Node(node string) *CurrentMonitorForNodeRepBuilder {
	b.apiReq.PathParams.Set("node", node)
	return b
}

func (b *CurrentMonitorForNodeRepBuilder) Build() *CurrentMonitorForNodeRep {
	req := &CurrentMonitorForNodeRep{}
	req.apiReq = b.apiReq
	return req
}

type CurrentMonitorForNodeResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	MonitorCurrent
}
