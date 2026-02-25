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
	PushGatewayServer     string            `json:"push_gateway_server"`
	Interval              string            `json:"interval"`
	Headers               map[string]string `json:"headers"`
	JobName               string            `json:"job_name"`
	Enable                bool              `json:"enable"`
	VmDistCollector       string            `json:"vm_dist_collector"`
	MnesiaCollector       string            `json:"mnesia_collector"`
	VmStatisticsCollector string            `json:"vm_statistics_collector"`
	VmSystemInfoCollector string            `json:"vm_system_info_collector"`
	VmMemoryCollector     string            `json:"vm_memory_collector"`
	VmMsaccCollector      string            `json:"vm_msacc_collector"`
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
	}
	return builder
}

func (b *UpdatePrometheusConfigReqBuilder) Enable(enable bool) *UpdatePrometheusConfigReqBuilder {
	b.apiReq.Body.(*UpdatePrometheusConfigReqBody).Enable = enable
	return b
}

func (b *UpdatePrometheusConfigReqBuilder) Interval(interval string) *UpdatePrometheusConfigReqBuilder {
	b.apiReq.Body.(*UpdatePrometheusConfigReqBody).Interval = interval
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
}

type Metrics struct {
	Enable   bool   `json:"enable"`
	Interval string `json:"interval"`
}
type Logs struct {
	Level          string `json:"level"`
	Enable         bool   `json:"enable"`
	ScheduledDelay string `json:"scheduled_delay"`
}

type Traces struct {
	Enable         bool   `json:"enable"`
	ScheduledDelay string `json:"scheduled_delay"`
	Filter         *Filter `json:"filter"`
}

type Exporter struct {
	Endpoint   string     `json:"endpoint"`
	SslOptions *SslOptions `json:"ssl_options"`
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
	}
	return builder
}

func (b *UpdateOpenTelemetryConfigReqBuilder) Build() *UpdateOpenTelemetryConfigReq {
	req := &UpdateOpenTelemetryConfigReq{}
	req.apiReq = b.apiReq
	return req
}
