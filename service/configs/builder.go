package configs

import "github.com/yao560909/emqx-admin-go-sdk/core"

type GetSysTopicsConfigReq struct {
	apiReq *core.APIReq
}

type GetSysTopicsConfigResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	SysTopicsConfig
}

type GetSysTopicsConfigReqBuilder struct {
	apiReq *core.APIReq
}

func NewGetSysTopicsConfigReqBuilder() *GetSysTopicsConfigReqBuilder {
	builder := &GetSysTopicsConfigReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *GetSysTopicsConfigReqBuilder) Build() *GetSysTopicsConfigReq {
	req := &GetSysTopicsConfigReq{}
	req.apiReq = b.apiReq
	return req
}

type UpdateSysTopicsConfigReq struct {
	apiReq *core.APIReq
}

type UpdateSysTopicsConfigReqBody struct {
	SysEventMessages     *SysEventMessages `json:"sys_event_messages"`
	SysHeartbeatInterval string            `json:"sys_heartbeat_interval"`
	SysMsgInterval       string            `json:"sys_msg_interval"`
}

type UpdateSysTopicsConfigResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	SysTopicsConfig
}

type UpdateSysTopicsConfigReqBuilder struct {
	apiReq *core.APIReq
}

func NewUpdateSysTopicsConfigReqBuilder() *UpdateSysTopicsConfigReqBuilder {
	builder := &UpdateSysTopicsConfigReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &UpdateSysTopicsConfigReqBody{},
	}
	return builder
}

func (b *UpdateSysTopicsConfigReqBuilder) SysEventMessages(sysEventMessages *SysEventMessages) *UpdateSysTopicsConfigReqBuilder {
	b.apiReq.Body.(*UpdateSysTopicsConfigReqBody).SysEventMessages = sysEventMessages
	return b
}

func (b *UpdateSysTopicsConfigReqBuilder) SysHeartbeatInterval(sysHeartbeatInterval string) *UpdateSysTopicsConfigReqBuilder {
	b.apiReq.Body.(*UpdateSysTopicsConfigReqBody).SysHeartbeatInterval = sysHeartbeatInterval
	return b
}

func (b *UpdateSysTopicsConfigReqBuilder) SysMsgInterval(sysMsgInterval string) *UpdateSysTopicsConfigReqBuilder {
	b.apiReq.Body.(*UpdateSysTopicsConfigReqBody).SysMsgInterval = sysMsgInterval
	return b
}

func (b *UpdateSysTopicsConfigReqBuilder) Build() *UpdateSysTopicsConfigReq {
	req := &UpdateSysTopicsConfigReq{}
	req.apiReq = b.apiReq
	return req
}

type GetSysmonConfigReq struct {
	apiReq *core.APIReq
}

type GetSysmonConfigResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	SysmonConfig
}

type GetSysmonConfigReqBuilder struct {
	apiReq *core.APIReq
}

func NewGetSysmonConfigReqBuilder() *GetSysmonConfigReqBuilder {
	builder := &GetSysmonConfigReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *GetSysmonConfigReqBuilder) Build() *GetSysmonConfigReq {
	req := &GetSysmonConfigReq{}
	req.apiReq = b.apiReq
	return req
}

type UpdateSysmonConfigReq struct {
	apiReq *core.APIReq
}

type UpdateSysmonConfigReqBody struct {
	SysmonVm SysmonVm `json:"vm"`
	SysmonOs SysmonOs `json:"os"`
}

type UpdateSysmonConfigResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	SysmonConfig
}

type UpdateSysmonConfigReqBuilder struct {
	apiReq *core.APIReq
}

func NewUpdateSysmonConfigReqBuilder() *UpdateSysmonConfigReqBuilder {
	builder := &UpdateSysmonConfigReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &UpdateSysmonConfigReqBody{},
	}
	return builder
}

func (b *UpdateSysmonConfigReqBuilder) Os(os SysmonOs) *UpdateSysmonConfigReqBuilder {
	b.apiReq.Body.(*UpdateSysmonConfigReqBody).SysmonOs = os
	return b
}

func (b *UpdateSysmonConfigReqBuilder) Vm(vm SysmonVm) *UpdateSysmonConfigReqBuilder {
	b.apiReq.Body.(*UpdateSysmonConfigReqBody).SysmonVm = vm
	return b
}

func (b *UpdateSysmonConfigReqBuilder) Build() *UpdateSysmonConfigReq {
	req := &UpdateSysmonConfigReq{}
	req.apiReq = b.apiReq
	return req
}

type GetGlobalZoneConfigReq struct {
	apiReq *core.APIReq
}

type GetGlobalZoneConfigResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	GlobalZoneConfig
}

type GetGlobalZoneConfigReqBuilder struct {
	apiReq *core.APIReq
}

func NewGetGlobalZoneConfigReqBuilder() *GetGlobalZoneConfigReqBuilder {
	builder := &GetGlobalZoneConfigReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *GetGlobalZoneConfigReqBuilder) Build() *GetGlobalZoneConfigReq {
	req := &GetGlobalZoneConfigReq{}
	req.apiReq = b.apiReq
	return req
}

type UpdateGlobalZoneConfigReq struct {
	apiReq *core.APIReq
}

type UpdateGlobalZoneConfigReqBody struct {
	FlappingDetect *FlappingDetect `json:"flapping_detect"`
	ForceGc        *ForceGc        `json:"force_gc"`
	ForceShutdown  *ForceShutdown  `json:"force_shutdown"`
	Mqtt           *Mqtt           `json:"mqtt"`
}

type UpdateGlobalZoneConfigResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	GlobalZoneConfig
}

type UpdateGlobalZoneConfigReqBuilder struct {
	apiReq *core.APIReq
}

func NewUpdateGlobalZoneConfigReqBuilder() *UpdateGlobalZoneConfigReqBuilder {
	builder := &UpdateGlobalZoneConfigReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &UpdateGlobalZoneConfigReqBody{},
	}
	return builder
}

func (b *UpdateGlobalZoneConfigReqBuilder) FlappingDetect(flappingDetect *FlappingDetect) *UpdateGlobalZoneConfigReqBuilder {
	b.apiReq.Body.(*UpdateGlobalZoneConfigReqBody).FlappingDetect = flappingDetect
	return b
}

func (b *UpdateGlobalZoneConfigReqBuilder) ForceGc(forceGc *ForceGc) *UpdateGlobalZoneConfigReqBuilder {
	b.apiReq.Body.(*UpdateGlobalZoneConfigReqBody).ForceGc = forceGc
	return b
}

func (b *UpdateGlobalZoneConfigReqBuilder) ForceShutdown(forceShutdown *ForceShutdown) *UpdateGlobalZoneConfigReqBuilder {
	b.apiReq.Body.(*UpdateGlobalZoneConfigReqBody).ForceShutdown = forceShutdown
	return b
}

func (b *UpdateGlobalZoneConfigReqBuilder) Mqtt(mqtt *Mqtt) *UpdateGlobalZoneConfigReqBuilder {
	b.apiReq.Body.(*UpdateGlobalZoneConfigReqBody).Mqtt = mqtt
	return b
}

func (b *UpdateGlobalZoneConfigReqBuilder) Build() *UpdateGlobalZoneConfigReq {
	req := &UpdateGlobalZoneConfigReq{}
	req.apiReq = b.apiReq
	return req
}

type GetAlarmConfigReq struct {
	apiReq *core.APIReq
}

type GetAlarmConfigResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	AlarmConfig
}

type GetAlarmConfigReqBuilder struct {
	apiReq *core.APIReq
}

func NewGetAlarmConfigReqBuilder() *GetAlarmConfigReqBuilder {
	builder := &GetAlarmConfigReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *GetAlarmConfigReqBuilder) Build() *GetAlarmConfigReq {
	req := &GetAlarmConfigReq{}
	req.apiReq = b.apiReq
	return req
}

type UpdateAlarmConfigReq struct {
	apiReq *core.APIReq
}

type UpdateAlarmConfigReqBody struct {
	Actions        []*string `json:"actions"`
	SizeLimit      int       `json:"size_limit"`
	ValidityPeriod string    `json:"validity_period"`
}

type UpdateAlarmConfigResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	AlarmConfig
}

type UpdateAlarmConfigReqBuilder struct {
	apiReq *core.APIReq
}

func NewUpdateAlarmConfigReqBuilder() *UpdateAlarmConfigReqBuilder {
	builder := &UpdateAlarmConfigReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &UpdateAlarmConfigReqBody{},
	}
	return builder
}

/*
*
Default: ["log","publish"]
The actions triggered when the alarm is activated.
Currently, the following actions are supported: log and publish.
log is to write the alarm to log (console or file).
publish is to publish the alarm as an MQTT message to the system topics:
$SYS/brokers/emqx@xx.xx.xx.x/alarms/activate and
$SYS/brokers/emqx@xx.xx.xx.x/alarms/deactivate
*/
func (b *UpdateAlarmConfigReqBuilder) Actions(actions []*string) *UpdateAlarmConfigReqBuilder {
	b.apiReq.Body.(*UpdateAlarmConfigReqBody).Actions = actions
	return b
}

/*
*
integer [ 1 .. 3000 ]
Default: 1000
The maximum total number of deactivated alarms to keep as history.
When this limit is exceeded, the oldest deactivated alarms are deleted to cap the total number.
*/
func (b *UpdateAlarmConfigReqBuilder) SizeLimit(sizeLimit int) *UpdateAlarmConfigReqBuilder {
	b.apiReq.Body.(*UpdateAlarmConfigReqBody).SizeLimit = sizeLimit
	return b
}

/*
*
Default: "24h"
Retention time of deactivated alarms. Alarms are not deleted immediately
when deactivated, but after the retention time.
*/
func (b *UpdateAlarmConfigReqBuilder) ValidityPeriod(validityPeriod string) *UpdateAlarmConfigReqBuilder {
	b.apiReq.Body.(*UpdateAlarmConfigReqBody).ValidityPeriod = validityPeriod
	return b
}

func (b *UpdateAlarmConfigReqBuilder) Build() *UpdateAlarmConfigReq {
	req := &UpdateAlarmConfigReq{}
	req.apiReq = b.apiReq
	return req
}
