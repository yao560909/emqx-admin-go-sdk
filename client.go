package sdk

import (
	"github.com/yao560909/emqx-admin-go-sdk/core"
)

type Edition int

const (
	openSource Edition = iota
	enterprise
)

type ClientOptionFunc func(config *core.Config)

func WithLogger(logger core.Logger) ClientOptionFunc {
	return func(config *core.Config) {
		config.Logger = logger
	}
}

func WithLogLevel(logLevel core.LogLevel) ClientOptionFunc {
	return func(config *core.Config) {
		config.LogLevel = logLevel
	}
}

type Client struct {
	Edition Edition
	Config  *core.Config
}

func NewClient(edition Edition, appId, appSecret, target, scheme string, options ...ClientOptionFunc) *Client {
	config := &core.Config{
		APIKey:    appId,
		APISecret: appSecret,
		Target:    target,
		Scheme:    scheme,
	}
	for _, option := range options {
		option(config)
	}

	core.NewLogger(config)

	client := &Client{
		Edition: edition,
		Config:  config,
	}

	return client
}

func initService(client *Client, config *core.Config) {

}
