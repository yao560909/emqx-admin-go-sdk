package configs

type SysTopicsConfig struct {
	SysEventMessages     *SysEventMessages `json:"sys_event_messages"`
	SysHeartbeatInterval string            `json:"sys_heartbeat_interval"`
	SysMsgInterval       string            `json:"sys_msg_interval"`
}

type SysEventMessages struct {
	ClientConnected    bool `json:"client_connected"`
	ClientDisconnected bool `json:"client_disconnected"`
	ClientSubscribed   bool `json:"client_subscribed"`
	ClientUnsubscribed bool `json:"client_unsubscribed"`
}

type SysmonConfig struct {
	SysmonOs SysmonOs `json:"os"`
	SysmonVm SysmonVm `json:"vm"`
}

type SysmonOs struct {
	CpuCheckInterval     *string `json:"cpu_check_interval,omitempty"`
	CpuHighWatermark     *string `json:"cpu_high_watermark,omitempty"`
	CpuLowWatermark      *string `json:"cpu_low_watermark,omitempty"`
	MemCheckInterval     *string `json:"mem_check_interval,omitempty"`
	ProcmemHighWatermark *string `json:"procmem_high_watermark,omitempty"`
	SysmemHighWatermark  *string `json:"sysmem_high_watermark,omitempty"`
}

type SysmonVm struct {
	BusyDistPort         *bool   `json:"busy_dist_port,omitempty"`
	BusyPort             *bool   `json:"busy_port,omitempty"`
	LargeHeap            *string `json:"large_heap,omitempty"`
	LongGc               *string `json:"long_gc,omitempty"`
	LongSchedule         *string `json:"long_schedule,omitempty"`
	ProcessCheckInterval *string `json:"process_check_interval,omitempty"`
	ProcessHighWatermark *string `json:"process_high_watermark,omitempty"`
	ProcessLowWatermark  *string `json:"process_low_watermark,omitempty"`
}

type GlobalZoneConfig struct {
	FlappingDetect *FlappingDetect `json:"flapping_detect"`
	ForceGc        *ForceGc        `json:"force_gc"`
	ForceShutdown  *ForceShutdown  `json:"force_shutdown"`
	Mqtt           *Mqtt           `json:"mqtt"`
}

type FlappingDetect struct {
	BanTime    *string `json:"ban_time,omitempty"`
	Enable     *bool   `json:"enable,omitempty"`
	MaxCount   *int    `json:"max_count,omitempty"`
	WindowTime *string `json:"window_time,omitempty"`
}

type ForceGc struct {
	Bytes  *string `json:"bytes,omitempty"`
	Count  *int    `json:"count,omitempty"`
	Enable *bool   `json:"enable,omitempty"`
}

type ForceShutdown struct {
	Enable         *bool   `json:"enable,omitempty"`
	MaxHeapSize    *string `json:"max_heap_size,omitempty"`
	MaxMailboxSize *int    `json:"max_mailbox_size,omitempty"`
}

type Mqtt struct {
	AwaitRelTimeout            *string  `json:"await_rel_timeout,omitempty"`
	ExclusiveSubscription      *bool    `json:"exclusive_subscription,omitempty"`
	IdleTimeout                *string  `json:"idle_timeout,omitempty"`
	IgnoreLoopDeliver          *bool    `json:"ignore_loop_deliver,omitempty"`
	KeepaliveMultiplier        *float64 `json:"keepalive_multiplier,omitempty"`
	MaxAwaitingRel             *int     `json:"max_awaiting_rel,omitempty"`
	MaxClientidLen             *int     `json:"max_clientid_len,omitempty"`
	MaxInflight                *int     `json:"max_inflight,omitempty"`
	MaxMqueueLen               *int     `json:"max_mqueue_len,omitempty"`
	MaxPacketSize              *string  `json:"max_packet_size,omitempty"`
	MaxQosAllowed              *int     `json:"max_qos_allowed,omitempty"`
	MaxSubscriptions           *string  `json:"max_subscriptions,omitempty"`
	MaxTopicAlias              *int     `json:"max_topic_alias,omitempty"`
	MaxTopicLevels             *int     `json:"max_topic_levels,omitempty"`
	MqueueDefaultPriority      *string  `json:"mqueue_default_priority,omitempty"`
	MqueuePriorities           *string  `json:"mqueue_priorities,omitempty"`
	MqueueStoreQos0            *bool    `json:"mqueue_store_qos0,omitempty"`
	PeerCertAsClientid         *string  `json:"peer_cert_as_clientid,omitempty"`
	PeerCertAsUsername         *string  `json:"peer_cert_as_username,omitempty"`
	ResponseInformation        *string  `json:"response_information,omitempty"`
	RetainAvailable            *bool    `json:"retain_available,omitempty"`
	RetryInterval              *string  `json:"retry_interval,omitempty"`
	ServerKeepalive            *string  `json:"server_keepalive,omitempty"`
	SessionExpiryInterval      *string  `json:"session_expiry_interval,omitempty"`
	SharedSubscription         *bool    `json:"shared_subscription,omitempty"`
	SharedSubscriptionStrategy *string  `json:"shared_subscription_strategy,omitempty"`
	StrictMode                 *bool    `json:"strict_mode,omitempty"`
	UpgradeQos                 *bool    `json:"upgrade_qos,omitempty"`
	UseUsernameAsClientid      *bool    `json:"use_username_as_clientid,omitempty"`
	WildcardSubscription       *bool    `json:"wildcard_subscription,omitempty"`
}

type AlarmConfig struct {
	Actions        []*string `json:"actions"`
	SizeLimit      int       `json:"size_limit"`
	ValidityPeriod string    `json:"validity_period"`
}

type DashboardConfig struct {
	Cors             bool                `json:"cors"`
	DefaultPassword  string              `json:"default_password"`
	DefaultUsername  string              `json:"default_username"`
	Listeners        *DashboardListeners `json:"listeners"`
	TokenExpiredTime string              `json:"token_expired_time"`
}

type DashboardListeners struct {
	Http *DashboardHttpListener `json:"http"`
}

type DashboardHttpListener struct {
	Backlog        int    `json:"backlog"`
	Bind           int    `json:"bind"`
	Inet6          bool   `json:"inet6"`
	Ipv6V6only     bool   `json:"ipv6_v6only"`
	MaxConnections int    `json:"max_connections"`
	NumAcceptors   int    `json:"num_acceptors"`
	ProxyHeader    bool   `json:"proxy_header"`
	SendTimeout    string `json:"send_timeout"`
}

type LogConfig struct {
	Console *Console `json:"console"`
	File    *File    `json:"file"`
}

type Console struct {
	Enable     bool   `json:"enable"`
	Formatter  string `json:"formatter"`
	Level      string `json:"level"`
	TimeOffset string `json:"time_offset"`
}

type File struct {
	Default *FileDefault `json:"default"`
}

type FileDefault struct {
	Enable        bool   `json:"enable"`
	Formatter     string `json:"formatter"`
	Level         string `json:"level"`
	Path          string `json:"path"`
	RotationCount int    `json:"rotation_count"`
	RotationSize  string `json:"rotation_size"`
	TimeOffset    string `json:"time_offset"`
}
