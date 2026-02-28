package monitor

import "github.com/yao560909/emqx-admin-go-sdk/core"

// Prometheus Config Structures
type GetPrometheusConfigReq struct {
	apiReq *core.APIReq
}

type GetPrometheusConfigResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	PrometheusConfig
}

type GetPrometheusConfigReqBuilder struct {
	apiReq *core.APIReq
}

func NewGetPrometheusConfigReqBuilder() *GetPrometheusConfigReqBuilder {
	builder := &GetPrometheusConfigReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *GetPrometheusConfigReqBuilder) Build() *GetPrometheusConfigReq {
	req := &GetPrometheusConfigReq{}
	req.apiReq = b.apiReq
	return req
}

type UpdatePrometheusConfigReq struct {
	apiReq *core.APIReq
}

type UpdatePrometheusConfigReqBody struct {
	PushGatewayServer     *string           `json:"push_gateway_server,omitempty"`
	Interval              *string           `json:"interval,omitempty"`
	Headers               map[string]string `json:"headers,omitempty"`
	JobName               *string           `json:"job_name,omitempty"`
	Enable                *bool             `json:"enable,omitempty"`
	VmDistCollector       *string           `json:"vm_dist_collector,omitempty"`
	MnesiaCollector       *string           `json:"mnesia_collector,omitempty"`
	VmStatisticsCollector *string           `json:"vm_statistics_collector,omitempty"`
	VmSystemInfoCollector *string           `json:"vm_system_info_collector,omitempty"`
	VmMemoryCollector     *string           `json:"vm_memory_collector,omitempty"`
	VmMsaccCollector      *string           `json:"vm_msacc_collector,omitempty"`
}

type UpdatePrometheusConfigResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	PrometheusConfig
}

type UpdatePrometheusConfigReqBuilder struct {
	apiReq *core.APIReq
}

func NewUpdatePrometheusConfigReqBuilder() *UpdatePrometheusConfigReqBuilder {
	builder := &UpdatePrometheusConfigReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &UpdatePrometheusConfigReqBody{},
	}
	return builder
}

/*
*
Default: false
Deprecated since 5.4.0, use prometheus.push_gateway.url instead
*/
func (b *UpdatePrometheusConfigReqBuilder) Enable(enable bool) *UpdatePrometheusConfigReqBuilder {
	b.apiReq.Body.(*UpdatePrometheusConfigReqBody).Enable = &enable
	return b
}

/*
*
Default: "15s"
Deprecated since 5.4.0, use prometheus.push_gateway.interval instead
*/
func (b *UpdatePrometheusConfigReqBuilder) Interval(interval string) *UpdatePrometheusConfigReqBuilder {
	b.apiReq.Body.(*UpdatePrometheusConfigReqBody).Interval = &interval
	return b
}

/*
*
Default: "http://127.0.0.1:9091"
Deprecated since 5.4.0, use prometheus.push_gateway.url instead
*/
func (b *UpdatePrometheusConfigReqBuilder) PushGatewayServer(pushGatewayServer string) *UpdatePrometheusConfigReqBuilder {
	b.apiReq.Body.(*UpdatePrometheusConfigReqBody).PushGatewayServer = &pushGatewayServer
	return b
}

/*
*
Default: {}
Deprecated since 5.4.0, use prometheus.push_gateway.headers instead
*/
func (b *UpdatePrometheusConfigReqBuilder) Headers(headers map[string]string) *UpdatePrometheusConfigReqBuilder {
	b.apiReq.Body.(*UpdatePrometheusConfigReqBody).Headers = headers
	return b
}

/*
*
Default: "${name}/instance/${name}~${host}"
Deprecated since 5.4.0, use prometheus.push_gateway.job_name instead
*/
func (b *UpdatePrometheusConfigReqBuilder) JobName(jobName string) *UpdatePrometheusConfigReqBuilder {
	b.apiReq.Body.(*UpdatePrometheusConfigReqBody).JobName = &jobName
	return b
}

/*
*
Default: "disabled"
Enum: "disabled" "enabled"
Deprecated since 5.4.0, use prometheus.collectors.vm_dist instead
*/
func (b *UpdatePrometheusConfigReqBuilder) VmDistCollector(vmDistCollector string) *UpdatePrometheusConfigReqBuilder {
	b.apiReq.Body.(*UpdatePrometheusConfigReqBody).VmDistCollector = &vmDistCollector
	return b
}

/*
*
Default: "disabled"
Enum: "enabled" "disabled"
Deprecated since 5.4.0, use prometheus.collectors.mnesia instead
*/
func (b *UpdatePrometheusConfigReqBuilder) MnesiaCollector(mnesiaCollector string) *UpdatePrometheusConfigReqBuilder {
	b.apiReq.Body.(*UpdatePrometheusConfigReqBody).MnesiaCollector = &mnesiaCollector
	return b
}

/*
*
Default: "disabled"
Enum: "enabled" "disabled"
Deprecated since 5.4.0, use prometheus.collectors.vm_statistics instead
*/
func (b *UpdatePrometheusConfigReqBuilder) VmStatisticsCollector(vmStatisticsCollector string) *UpdatePrometheusConfigReqBuilder {
	b.apiReq.Body.(*UpdatePrometheusConfigReqBody).VmStatisticsCollector = &vmStatisticsCollector
	return b
}

/*
*
Default: "disabled"
Enum: "enabled" "disabled"
Deprecated, use prometheus.collectors.vm_system_info instead
*/
func (b *UpdatePrometheusConfigReqBuilder) VmSystemInfoCollector(vmSystemInfoCollector string) *UpdatePrometheusConfigReqBuilder {
	b.apiReq.Body.(*UpdatePrometheusConfigReqBody).VmSystemInfoCollector = &vmSystemInfoCollector
	return b
}

/*
*
Default: "disabled"
Enum: "enabled" "disabled"
Deprecated since 5.4.0, use prometheus.collectors.vm_memory instead
*/
func (b *UpdatePrometheusConfigReqBuilder) VmMemoryCollector(vmMemoryCollector string) *UpdatePrometheusConfigReqBuilder {
	b.apiReq.Body.(*UpdatePrometheusConfigReqBody).VmMemoryCollector = &vmMemoryCollector
	return b
}

/*
*
Default: "disabled"
Enum: "enabled" "disabled"
Deprecated since 5.4.0, use prometheus.collectors.vm_msacc instead
*/
func (b *UpdatePrometheusConfigReqBuilder) VmMsaccCollector(vmMsaccCollector string) *UpdatePrometheusConfigReqBuilder {
	b.apiReq.Body.(*UpdatePrometheusConfigReqBody).VmMsaccCollector = &vmMsaccCollector
	return b
}

func (b *UpdatePrometheusConfigReqBuilder) Build() *UpdatePrometheusConfigReq {
	req := &UpdatePrometheusConfigReq{}
	req.apiReq = b.apiReq
	return req
}

// OpenTelemetry Config Structures
type GetOpenTelemetryConfigReq struct {
	apiReq *core.APIReq
}

type GetOpenTelemetryConfigResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	OpenTelemetryConfig
}

type GetOpenTelemetryConfigReqBuilder struct {
	apiReq *core.APIReq
}

func NewGetOpenTelemetryConfigReqBuilder() *GetOpenTelemetryConfigReqBuilder {
	builder := &GetOpenTelemetryConfigReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *GetOpenTelemetryConfigReqBuilder) Build() *GetOpenTelemetryConfigReq {
	req := &GetOpenTelemetryConfigReq{}
	req.apiReq = b.apiReq
	return req
}

type UpdateOpenTelemetryConfigReq struct {
	apiReq *core.APIReq
}

type UpdateOpenTelemetryConfigReqBody struct {
	Metrics  *Metrics  `json:"metrics,omitempty"`
	Logs     *Logs     `json:"logs,omitempty"`
	Traces   *Traces   `json:"traces,omitempty"`
	Exporter *Exporter `json:"exporter,omitempty"`
}

type Metrics struct {
	Enable   *bool   `json:"enable,omitempty"`
	Interval *string `json:"interval,omitempty"`
}
type Logs struct {
	Level          *string `json:"level,omitempty"`
	Enable         *bool   `json:"enable,omitempty"`
	ScheduledDelay *string `json:"scheduled_delay,omitempty"`
}

type Traces struct {
	Enable         *bool   `json:"enable,omitempty"`
	ScheduledDelay *string `json:"scheduled_delay,omitempty"`
	Filter         *Filter `json:"filter,omitempty"`
}

type Exporter struct {
	Endpoint   *string     `json:"endpoint,omitempty"`
	SslOptions *SslOptions `json:"ssl_options,omitempty"`
}

type UpdateOpenTelemetryConfigResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	OpenTelemetryConfig
}

type UpdateOpenTelemetryConfigReqBuilder struct {
	apiReq *core.APIReq
}

func NewUpdateOpenTelemetryConfigReqBuilder() *UpdateOpenTelemetryConfigReqBuilder {
	builder := &UpdateOpenTelemetryConfigReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &UpdateOpenTelemetryConfigReqBody{},
	}
	return builder
}
func (b *UpdateOpenTelemetryConfigReqBuilder) Metrics(metrics Metrics) *UpdateOpenTelemetryConfigReqBuilder {
	b.apiReq.Body.(*UpdateOpenTelemetryConfigReqBody).Metrics = &metrics
	return b
}
func (b *UpdateOpenTelemetryConfigReqBuilder) Logs(logs Logs) *UpdateOpenTelemetryConfigReqBuilder {
	b.apiReq.Body.(*UpdateOpenTelemetryConfigReqBody).Logs = &logs
	return b
}
func (b *UpdateOpenTelemetryConfigReqBuilder) Traces(traces Traces) *UpdateOpenTelemetryConfigReqBuilder {
	b.apiReq.Body.(*UpdateOpenTelemetryConfigReqBody).Traces = &traces
	return b
}
func (b *UpdateOpenTelemetryConfigReqBuilder) Exporter(exporter Exporter) *UpdateOpenTelemetryConfigReqBuilder {
	b.apiReq.Body.(*UpdateOpenTelemetryConfigReqBody).Exporter = &exporter
	return b
}

func (b *UpdateOpenTelemetryConfigReqBuilder) Build() *UpdateOpenTelemetryConfigReq {
	req := &UpdateOpenTelemetryConfigReq{}
	req.apiReq = b.apiReq
	return req
}

type GetPrometheusStatsReq struct {
	apiReq *core.APIReq
}

type GetPrometheusStatsReqBuilder struct {
	apiReq *core.APIReq
}

func NewGetPrometheusStatsReqBuilder() *GetPrometheusStatsReqBuilder {
	builder := &GetPrometheusStatsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *GetPrometheusStatsReqBuilder) Build() *GetPrometheusStatsReq {
	req := &GetPrometheusStatsReq{}
	req.apiReq = b.apiReq
	return req
}

type GetPrometheusStatsResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
}
