package core

import (
	"bytes"
	"fmt"
	"time"

	"github.com/valyala/fasthttp"
)

var reqTranslator ReqTranslator

type Requester struct {
	client *fasthttp.Client
	uri    *fasthttp.URI
}

func NewRequester(config *Config) *Requester {
	uri := &fasthttp.URI{}
	uri.SetUsername(config.APIKey)
	uri.SetPassword(config.APISecret)
	uri.SetScheme(config.Scheme)
	uri.SetHost(config.Target)
	return &Requester{
		uri: uri,
		client: &fasthttp.Client{
			Name:                "emqx-admin-go-sdk",
			MaxConnsPerHost:     1000,
			MaxIdleConnDuration: 30 * time.Second,
			ReadTimeout:         5 * time.Second,
			WriteTimeout:        5 * time.Second,
			MaxConnWaitTimeout:  5 * time.Second,
			TLSConfig:           config.TLSClientConfig.ToTLSConfig(),
			DialDualStack:       true,
		},
	}
}

func (r *Requester) DoRequest(apiReq *APIReq, options ...RequestOptionFunc) (*APIResp, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	reqURI := fasthttp.AcquireURI()
	defer fasthttp.ReleaseURI(reqURI)
	r.uri.CopyTo(reqURI)
	if apiReq.SkipAuth {
		reqURI.SetUsername("")
		reqURI.SetPassword("")
	}
	req.SetURI(reqURI)

	path, err := reqTranslator.Path(apiReq.ApiPath, apiReq.PathParams)
	if err != nil {
		return nil, err
	}
	req.URI().SetPath(path)

	req.URI().QueryArgs().Reset()
	for k, vs := range apiReq.QueryParams {
		for _, v := range vs {
			req.URI().QueryArgs().Add(k, v)
		}
	}

	req.Header.SetMethod(apiReq.HttpMethod)

	opt := &RequestOption{
		Header: make(map[string]string, 8),
	}
	for _, f := range options {
		f(opt)
	}
	for k, v := range opt.Header {
		req.Header.Set(k, v)
	}

	contentType, body, err := reqTranslator.Payload(apiReq.Body)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", contentType)
		req.SetBody(body)
	} else {
		req.ResetBody()
	}

	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	err = r.client.Do(req, resp)
	if err != nil {
		return nil, fmt.Errorf("request %s failed: %w", req.URI().String(), err)
	}
	statusCode := resp.StatusCode()
	apiResp := &APIResp{
		StatusCode: statusCode,
		Header:     make(map[string]string, 8),
	}
	resp.Header.VisitAll(func(k, v []byte) {
		apiResp.Header[string(bytes.Clone(k))] = string(bytes.Clone(v))
	})
	apiResp.RawBody = append([]byte(nil), resp.Body()...)
	return apiResp, nil
}
