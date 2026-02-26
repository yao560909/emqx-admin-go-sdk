package alarms

import (
	"github.com/yao560909/emqx-admin-go-sdk/core"
)

// --------------------------------
type ListAlarmsReq struct {
	apiReq *core.APIReq
}

type ListAlarmsReqBuilder struct {
	apiReq *core.APIReq
}

func NewListAlarmsReqBuilder() *ListAlarmsReqBuilder {
	builder := &ListAlarmsReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
	}
	return builder
}

func (b *ListAlarmsReqBuilder) Build() *ListAlarmsReq {
	req := &ListAlarmsReq{}
	req.apiReq = b.apiReq
	return req
}

/*
*
It is used to specify the alarm type of the query.
When true, it returns the currently activated alarm,
and when it is false, it returns the historical alarm.
The default is false.
*/
func (b *ListAlarmsReqBuilder) Activated(activated string) *ListAlarmsReqBuilder {
	b.apiReq.QueryParams.Set("activated", activated)
	return b
}

/*
*
Default: 1
Example: page=1
Page number of the results to fetch.
*/
func (b *ListAlarmsReqBuilder) Page(page string) *ListAlarmsReqBuilder {
	b.apiReq.QueryParams.Set("page", page)
	return b
}

/*
*
integer [ 1 .. 1000 ]
Default: 100
Example: limit=50
Results per page(max 1000)
*/
func (b *ListAlarmsReqBuilder) Limit(limit string) *ListAlarmsReqBuilder {
	b.apiReq.QueryParams.Set("limit", limit)
	return b
}

type AlarmListResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	Data []*Alarm `json:"data"`
	Meta Meta     `json:"meta"`
}
