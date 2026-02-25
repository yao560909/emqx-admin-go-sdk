package sdk

import (
	"github.com/yao560909/emqx-admin-go-sdk/core"
	"github.com/yao560909/emqx-admin-go-sdk/service/authentication"
	"github.com/yao560909/emqx-admin-go-sdk/service/banned"
	"github.com/yao560909/emqx-admin-go-sdk/service/clients"
	"github.com/yao560909/emqx-admin-go-sdk/service/metrics"
	"github.com/yao560909/emqx-admin-go-sdk/service/nodes"
	"github.com/yao560909/emqx-admin-go-sdk/service/topics"
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

	Nodes          *nodes.NodesService
	Clients        *clients.ClientsService
	Metrics        *metrics.MetricsService
	Banned         *banned.BannedService
	Authentication *authentication.AuthenticationService
	Topics         *topics.TopicsService
	Authorization  *authentication.AuthenticationService
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
		client.Nodes = nodes.NewService(config)
		client.Clients = clients.NewService(config)
		client.Metrics = metrics.NewService(config)
		client.Banned = banned.NewService(config)
		client.Authentication = authentication.NewService(config)
		client.Topics = topics.NewService(config)
		client.Authorization = authentication.NewService(config)
	case Enterprise:
		client.Nodes = nodes.NewService(config)
		client.Clients = clients.NewService(config)
		client.Metrics = metrics.NewService(config)
		client.Banned = banned.NewService(config)
		client.Authentication = authentication.NewService(config)
		client.Topics = topics.NewService(config)
		client.Authorization = authentication.NewService(config)
	default:
	}
}
