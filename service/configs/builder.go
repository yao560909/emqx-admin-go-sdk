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
