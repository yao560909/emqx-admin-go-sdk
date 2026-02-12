package sdk

import (
	"github.com/yao560909/emqx-admin-go-sdk/core"
	"github.com/yao560909/emqx-admin-go-sdk/service/nodes"
)

type Edition int

const (
	OpenSource Edition = iota
	Enterprise
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

	Nodes *nodes.NodesService
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
	initService(client, config)
	return client
}

func initService(client *Client, config *core.Config) {
	edition := client.Edition
	switch edition {
	case OpenSource:
		client.Nodes=nodes.NewService(config)
	case Enterprise:
		client.Nodes=nodes.NewService(config)
	default:
	}
}
