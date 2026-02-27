package publish

import (
	"github.com/yao560909/emqx-admin-go-sdk/core"
)

type PublishReq struct {
	apiReq *core.APIReq
}

type PublishReqBuilder struct {
	apiReq *core.APIReq
}

type PublishReqBody struct {
	PayloadEncoding string     `json:"payload_encoding"`
	Topic           string     `json:"topic"`
	Qos             string     `json:"qos"`
	Clientid        string     `json:"clientid"`
	Payload         string     `json:"payload"`
	Retain          bool       `json:"retain"`
	Properties      Properties `json:"properties"`
}
type Properties struct {
	PayloadFormatIndicator string            `json:"payload_format_indicator"`
	MessageExpiryInterval  string            `json:"message_expiry_interval"`
	ResponseTopic          string            `json:"response_topic"`
	CorrelationData        string            `json:"correlation_data"`
	UserProperties         map[string]string `json:"user_properties"`
	ContentType            string            `json:"content_type"`
}

// NewPublishReqBuilder creates a new PublishReqBuilder
func NewPublishReqBuilder() *PublishReqBuilder {
	builder := &PublishReqBuilder{}
	builder.apiReq = &core.APIReq{
		PathParams:  core.PathParams{},
		QueryParams: core.QueryParams{},
		Body:        &PublishReqBody{},
	}
	return builder
}

// Topic sets the topic for the publish request
func (b *PublishReqBuilder) Topic(topic string) *PublishReqBuilder {
	b.apiReq.Body.(*PublishReqBody).Topic = topic
	return b
}

// Payload sets the payload for the publish request
func (b *PublishReqBuilder) Payload(payload string) *PublishReqBuilder {
	b.apiReq.Body.(*PublishReqBody).Payload = payload
	return b
}

// QoS sets the QoS level for the publish request
func (b *PublishReqBuilder) QoS(qos string) *PublishReqBuilder {
	b.apiReq.Body.(*PublishReqBody).Qos = qos
	return b
}

// Retain sets the retain flag for the publish request
func (b *PublishReqBuilder) Retain(retain bool) *PublishReqBuilder {
	b.apiReq.Body.(*PublishReqBody).Retain = retain
	return b
}

// ClientID sets the client ID for the publish request
func (b *PublishReqBuilder) ClientID(clientId string) *PublishReqBuilder {
	b.apiReq.Body.(*PublishReqBody).Clientid = clientId
	return b
}

func (b *PublishReqBuilder) PayloadEncoding(payloadEncoding string) *PublishReqBuilder {
	b.apiReq.Body.(*PublishReqBody).PayloadEncoding = payloadEncoding
	return b
}
func (b *PublishReqBuilder) Properties(properties Properties) *PublishReqBuilder {
	b.apiReq.Body.(*PublishReqBody).Properties = properties
	return b
}

// Build builds the PublishReq
func (b *PublishReqBuilder) Build() *PublishReq {
	req := &PublishReq{}
	req.apiReq = b.apiReq
	return req
}

type PublishResp struct {
	*core.APIResp `json:"-"`
	core.CodeError
	PublishResponse
}
