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
