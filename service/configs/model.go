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
