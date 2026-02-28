package sdk

import (
	"github.com/yao560909/emqx-admin-go-sdk/core"
	"github.com/yao560909/emqx-admin-go-sdk/service/alarms"
	"github.com/yao560909/emqx-admin-go-sdk/service/apikeys"
	"github.com/yao560909/emqx-admin-go-sdk/service/authentication"
	"github.com/yao560909/emqx-admin-go-sdk/service/authorization"
	"github.com/yao560909/emqx-admin-go-sdk/service/banned"
	"github.com/yao560909/emqx-admin-go-sdk/service/clients"
	"github.com/yao560909/emqx-admin-go-sdk/service/cluster"
	"github.com/yao560909/emqx-admin-go-sdk/service/configs"
	"github.com/yao560909/emqx-admin-go-sdk/service/dashboard"
	"github.com/yao560909/emqx-admin-go-sdk/service/listeners"
	"github.com/yao560909/emqx-admin-go-sdk/service/metrics"
	"github.com/yao560909/emqx-admin-go-sdk/service/monitor"
	"github.com/yao560909/emqx-admin-go-sdk/service/nodes"
	"github.com/yao560909/emqx-admin-go-sdk/service/publish"
	"github.com/yao560909/emqx-admin-go-sdk/service/subscriptions"
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
	Authorization  *authorization.AuthorizationService
	Configs        *configs.ConfigsService
	Alarms         *alarms.AlarmsService
	Subscriptions  *subscriptions.SubscriptionsService
	APIKeys        *apikeys.APIKeysService
	Cluster        *cluster.ClusterService
	Dashboard      *dashboard.DashboardService
	Listeners      *listeners.ListenersService
	Monitor        *monitor.MonitorService
	Publish        *publish.PublishService
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
		client.Authorization = authorization.NewService(config)
		client.Configs = configs.NewService(config)
		client.Alarms = alarms.NewService(config)
		client.Subscriptions = subscriptions.NewService(config)
		client.APIKeys = apikeys.NewService(config)
		client.Cluster = cluster.NewService(config)
		client.Dashboard = dashboard.NewService(config)
		client.Listeners = listeners.NewService(config)
		client.Monitor = monitor.NewService(config)
		client.Publish = publish.NewService(config)
	case Enterprise:
		client.Nodes = nodes.NewService(config)
		client.Clients = clients.NewService(config)
		client.Metrics = metrics.NewService(config)
		client.Banned = banned.NewService(config)
		client.Authentication = authentication.NewService(config)
		client.Topics = topics.NewService(config)
		client.Authorization = authorization.NewService(config)
		client.Configs = configs.NewService(config)
		client.Alarms = alarms.NewService(config)
		client.Subscriptions = subscriptions.NewService(config)
		client.APIKeys = apikeys.NewService(config)
		client.Cluster = cluster.NewService(config)
		client.Dashboard = dashboard.NewService(config)
		client.Listeners = listeners.NewService(config)
		client.Monitor = monitor.NewService(config)
		client.Publish = publish.NewService(config)
	default:
	}
}
