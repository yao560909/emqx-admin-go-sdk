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
	CpuCheckInterval     string `json:"cpu_check_interval"`
	CpuHighWatermark     string `json:"cpu_high_watermark"`
	CpuLowWatermark      string `json:"cpu_low_watermark"`
	MemCheckInterval     string `json:"mem_check_interval"`
	ProcmemHighWatermark string `json:"procmem_high_watermark"`
	SysmemHighWatermark  string `json:"sysmem_high_watermark"`
}

type SysmonVm struct {
	BusyDistPort         bool   `json:"busy_dist_port"`
	BusyPort             bool   `json:"busy_port"`
	LargeHeap            string `json:"large_heap"`
	LongGc               string `json:"long_gc"`
	LongSchedule         string `json:"long_schedule"`
	ProcessCheckInterval string `json:"process_check_interval"`
	ProcessHighWatermark string `json:"process_high_watermark"`
	ProcessLowWatermark  string `json:"process_low_watermark"`
}

type GlobalZoneConfig struct {
	FlappingDetect *FlappingDetect `json:"flapping_detect"`
	ForceGc        *ForceGc        `json:"force_gc"`
	ForceShutdown  *ForceShutdown  `json:"force_shutdown"`
	Mqtt           *Mqtt           `json:"mqtt"`
}

type FlappingDetect struct {
	BanTime    string `json:"ban_time"`
	Enable     bool   `json:"enable"`
	MaxCount   int    `json:"max_count"`
	WindowTime string `json:"window_time"`
}

type ForceGc struct {
	Bytes  string `json:"bytes"`
	Count  int    `json:"count"`
	Enable bool   `json:"enable"`
}

type ForceShutdown struct {
	Enable         bool   `json:"enable"`
	MaxHeapSize    string `json:"max_heap_size"`
	MaxMailboxSize int    `json:"max_mailbox_size"`
}

type Mqtt struct {
	AwaitRelTimeout            string  `json:"await_rel_timeout"`
	ExclusiveSubscription      bool    `json:"exclusive_subscription"`
	IdleTimeout                string  `json:"idle_timeout"`
	IgnoreLoopDeliver          bool    `json:"ignore_loop_deliver"`
	KeepaliveMultiplier        float64 `json:"keepalive_multiplier"`
	MaxAwaitingRel             int     `json:"max_awaiting_rel"`
	MaxClientidLen             int     `json:"max_clientid_len"`
	MaxInflight                int     `json:"max_inflight"`
	MaxMqueueLen               int     `json:"max_mqueue_len"`
	MaxPacketSize              string  `json:"max_packet_size"`
	MaxQosAllowed              int     `json:"max_qos_allowed"`
	MaxSubscriptions           string  `json:"max_subscriptions"`
	MaxTopicAlias              int     `json:"max_topic_alias"`
	MaxTopicLevels             int     `json:"max_topic_levels"`
	MqueueDefaultPriority      string  `json:"mqueue_default_priority"`
	MqueuePriorities           string  `json:"mqueue_priorities"`
	MqueueStoreQos0            bool    `json:"mqueue_store_qos0"`
	PeerCertAsClientid         string  `json:"peer_cert_as_clientid"`
	PeerCertAsUsername         string  `json:"peer_cert_as_username"`
	ResponseInformation        string  `json:"response_information"`
	RetainAvailable            bool    `json:"retain_available"`
	RetryInterval              string  `json:"retry_interval"`
	ServerKeepalive            string  `json:"server_keepalive"`
	SessionExpiryInterval      string  `json:"session_expiry_interval"`
	SharedSubscription         bool    `json:"shared_subscription"`
	SharedSubscriptionStrategy string  `json:"shared_subscription_strategy"`
	StrictMode                 bool    `json:"strict_mode"`
	UpgradeQos                 bool    `json:"upgrade_qos"`
	UseUsernameAsClientid      bool    `json:"use_username_as_clientid"`
	WildcardSubscription       bool    `json:"wildcard_subscription"`
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
